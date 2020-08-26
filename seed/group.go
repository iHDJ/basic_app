package seed

import (
	"basic_app/dao/model"

	"github.com/jinzhu/gorm"
)

var (
	ManagerGroup = model.Group{ID: 1}
)

func SeedGroup(db *gorm.DB) (err error) {

	return
}
