package seed

import (
	"basic_app/dao/model"

	"github.com/jinzhu/gorm"
)

func SeedManagerData(db *gorm.DB) (err error) {
	tx := db.Begin()

	adminUser := model.User{
		Name:  "ChinaHDJ",
		Email: "chinahdj1@gmail.com",
	}
	adminUser.SetPassword("asdfasdF")

	if err = tx.Create(&adminUser).Error; err != nil {
		tx.Rollback()
		return
	}

	adminMember := model.Member{
		UserID:  adminUser.ID,
		GroupID: ManagerGroup.ID,
	}

	if err = tx.Create(&adminMember).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	return
}
