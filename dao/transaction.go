package dao

import "github.com/jinzhu/gorm"

func (dao *Dao) Begin() *Dao {
	dao.orm.Commit()
	return &Dao{
		orm: dao.orm.Begin(),
	}
}

func (dao *Dao) Rollback() error {
	return dao.orm.Rollback().Error
}

func (dao *Dao) Commit() error {
	return dao.orm.Commit().Error
}

func (dao *Dao) Transaction(t func(dao *Dao) error) {
	dao.orm.Transaction(func(tx *gorm.DB) error {
		txDao := &Dao{orm: tx}

		return t(txDao)
	})
}
