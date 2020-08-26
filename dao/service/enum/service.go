package enum

import (
	"basic_app/dao/model"
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	EnumNameExistedErr = errors.New("该枚举名字已被占用")
)

type EnumerationService struct {
	orm *gorm.DB
}

func NewService(orm *gorm.DB) *EnumerationService {
	return &EnumerationService{orm: orm}
}

func (dao *EnumerationService) UpdateEnumeration(enum model.Enumeration) (err error) {

	return
}

func (service *EnumerationService) CreateEnumeration(enum model.Enumeration) (err error) {
	if _, errFindEnum := service.FindEnumerationByName(enum.Name); errFindEnum == nil {
		return EnumNameExistedErr
	}

	err = service.orm.Create(&enum).Error
	return
}

func (service *EnumerationService) CreateEnumerationLabel(enumID uint64, label model.EnumerationLabel) (err error) {

	return
}
