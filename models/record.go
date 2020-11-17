package models

import "fmt"

type Record struct {
	Id        int     `xorm:"id pk autoincr" json:"id"`
	Name      string  `xorm:"name" json:"name"`
	Sn        string  `xorm:"sn" json:"sn"`
	Company   string  `xorm:"company" json:"company"`
	Num       float32 `xorm:"num" json:"num"`
	Tel       string  `xorm:"tel" json:"mobile"`
	Type      int     `xorm:"type" json:"type"`
	UserNo    string  `xorm:"user_no" json:"user_no"`
	IdNum     string  `xorm:"id_num" json:"id_num"`
	FaceImage string  `xorm:"face_image" json:"face_image"`
	CreatedAt int     `xorm:"created_at" json:"created_at"`
}
type RecordDao struct {
}

func (m *RecordDao) TableName() string {
	return "t_record"
}

//FindProducts 查找符合条件的所有商品
func (this *RecordDao) GetList(storeId int, q map[string]interface{}) (p []Record, total int64, err error) {
	u := UserDao{}
	query := engine.Table(this.TableName()).Alias("d").
		Join("left", []string{u.TableName(), "u"}, "d.user_id=u.id")
	if storeId > 0 {
		query = query.And("d.store_id= ?", storeId)
	}
	if q["start_time"].(int) > 0 {
		query = query.And("d.created_at >= ?", q["start_time"])
	}
	if q["end_time"].(int) > 0 {
		query = query.And("d.created_at < ?", q["end_time"])
	}
	if q["num"].(float64) > 0 {
		query = query.And("d.num >= ?", q["num"])
	}

	if q["name"].(string) != "" {
		query = query.And(fmt.Sprintf("u.name like '%%%s%%'", q["name"]))
	}
	if q["id_num"].(string) != "" {
		query = query.And(fmt.Sprintf("u.id_num like '%%%s%%'", q["id_num"]))
	}
	if q["user_no"].(string) != "" {
		query = query.And(fmt.Sprintf("u.user_no like '%%%s%%'", q["user_no"]))
	}

	total, err = query.Clone().Count()
	if err != nil {
		return
	}
	if total == 0 {
		return
	}

	s := StoreDao{}
	err = query.
		Limit(q["limit"].(int), q["offset"].(int)).
		Join("left", []string{s.TableName(), "s"}, "s.id=d.store_id").
		Select("d.id,d.num,d.type,d.sn,d.created_at,u.name,u.tel,s.name as company,u.user_no,u.id_num").
		Desc("d.created_at").
		Find(&p)
	if err != nil {
		return
	}
	return
}
