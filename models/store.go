package models

type Store struct {
	Id        int    `xorm:"id pk autoincr"`
	UserName  string `xorm:"username" json:"username"`
	Pwd       string `xorm:"password"  json:"-"`
	Address   string `xorm:"address"  json:"address"`
	Tel       string `xorm:"tel"  json:"tel"`
	Role      int    `xorm:"role"  json:"role"`
	Name      string `xorm:"name"  json:"name"`
	Status    int    `xorm:"status" json:"status"`
	UpdatedAt int64  `xorm:"updated_at" json:"-"`
	CreatedAt int64  `xorm:"created_at" json:"-"`
}
type StoreDao struct {
}

func (m *StoreDao) TableName() string {
	return "t_store"
}

//FindProducts 查找符合条件的所有商品
func (this *StoreDao) GetList(size, offset int) (p []Store, total int64, err error) {
	query := engine.Table(this.TableName()).
		Alias("d").
		Where("status=1 and role=1")
	total, err = query.Clone().Count()
	if err != nil {
		return
	}
	if total == 0 {
		return
	}
	if size > 0 {
		query = query.Limit(size, offset)
	}
	err = query.Find(&p)
	if err != nil {
		return
	}
	return
}

func (this *StoreDao) FindById(Id int) (res Store, err error) {
	_, err = engine.Table(this.TableName()).
		Where("id = ? and status = 1", Id).
		Get(&res)
	return
}

func (this *StoreDao) FindByUserName(username string) (res Store, err error) {
	_, err = engine.Table(this.TableName()).
		Where("username = ? and status = 1", username).
		Get(&res)
	return
}

func (this *StoreDao) Add(m Store) (err error) {
	_, err = engine.Table(this.TableName()).Insert(&m)
	return
}

func (this *StoreDao) Update(m *Store) (err error) {
	_, err = engine.Table(this.TableName()).Where("id = ?", m.Id).Update(Store{
		Pwd:       m.Pwd,
		Name:      m.Name,
		Tel:       m.Tel,
		Address:   m.Address,
		UpdatedAt: m.UpdatedAt,
	})
	return
}
