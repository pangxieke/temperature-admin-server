package models

type Device struct {
	Id        int    `xorm:"id pk autoincr" json:"id"`
	Sn        string `xorm:"sn" json:"sn"`
	Mac       string `xorm:"mac"  json:"mac"`
	StoreId   int    `xorm:"store_id"  json:"store_id"`
	Status    int    `xorm:"status" json:"status"`
	Province  int    `xorm:"province" json:"province"`
	City      int    `xorm:"city" json:"city"`
	Remake    string `xorm:"remake"  json:"remake"`
	CreatedAt int    `xorm:"created_at" json:"created_at"`
}

type DeviceStore struct {
	Device    `xorm:"extends"`
	StoreName string `xorm:"store_name" `
}

type DeviceDao struct {
}

func (m *DeviceDao) TableName() string {
	return "t_device"
}

//FindProducts 查找符合条件的所有商品
func (this *DeviceDao) GetList(size, offset, storeId int) (p []DeviceStore, total int64, err error) {
	query := engine.Table(this.TableName()).Alias("d")
	if storeId > 0 {
		query = query.And("d.store_id= ?", storeId)
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
		Join("left", []string{s.TableName(), "s"}, "s.id=d.store_id").
		Select("d.id,d.sn,d.status,d.store_id,d.province,d.city,d.remake,d.created_at,s.name as store_name").
		Desc("d.created_at").
		Limit(size, offset).
		Find(&p)
	if err != nil {
		return
	}
	return
}

func (this *DeviceDao) FindBySn(sn string) (res Device, err error) {
	_, err = engine.Table(this.TableName()).
		Where("sn = ? ", sn).
		Get(&res)
	return
}

func (this *DeviceDao) Bind(m Device) (err error) {
	_, err = engine.Table(this.TableName()).
		Where("id = ? ", m.Id).
		Cols("store_id", "remake", "province", "city").
		Update(Device{StoreId: m.StoreId, Remake: m.Remake, Province: m.Province, City: m.City})
	return
}
