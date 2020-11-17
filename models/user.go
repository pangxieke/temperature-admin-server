package models

import "fmt"

type User struct {
	Id         int    `xorm:"id pk autoincr" json:"id"`
	Name       string `xorm:"name" json:"name"`
	UserNo     string `xorm:"user_no" json:"user_no"`
	Tel        string `xorm:"tel" json:"tel"`
	IdNum      string `xorm:"id_num" json:"id_num"`
	Age        int    `xorm:"age" json:"age"`
	CompanyId  int    `xorm:"company_id" json:"company_id"`
	Company    string `xorm:"company" json:"company"`
	Department string `xorm:"department" json:"department"`
	FaceImage  string `xorm:"face_image" json:"face_image"`
	CreatedAt  int    `xorm:"created_at" json:"created_at"`
}
type UserDao struct {
}

func (m *UserDao) TableName() string {
	return "t_user"
}

//FindProducts 查找符合条件的所有商品
func (this *UserDao) GetList(storeId int, q map[string]interface{}) (p []User, total int64, err error) {
	query := engine.Table(this.TableName()).Alias("d")
	if storeId > 0 {
		query = query.And("d.company_id= ?", storeId)
	}
	if q["name"].(string) != "" {
		query = query.And(fmt.Sprintf("d.name like '%%%s%%'", q["name"]))
	}
	total, err = query.Clone().Count()
	if err != nil {
		return
	}
	if total == 0 {
		return
	}
	err = query.
		Limit(q["limit"].(int), q["offset"].(int)).
		Find(&p)
	if err != nil {
		return
	}
	return
}

func (this *UserDao) FindById(Id int) (res User, err error) {
	_, err = engine.Table(this.TableName()).
		Where("id = ?", Id).
		Get(&res)
	return
}

func (this *UserDao) Update(m User) (err error) {
	_, err = engine.Table(this.TableName()).Where("id = ?", m.Id).Update(User{
		Name:       m.Name,
		Age:        m.Age,
		Tel:        m.Tel,
		IdNum:      m.IdNum,
		Department: m.Department,
	})
	return
}

func (this *UserDao) Del(Id int) (err error) {
	_, err = engine.Table(this.TableName()).
		Where("id = ? ", Id).
		Cols("company_id", "company").
		Update(User{Company: "", CompanyId: 0})
	return
}
