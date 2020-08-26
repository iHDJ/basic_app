package enum

import (
	"basic_app/dao/model"
	"errors"
)

type QueryForm struct {
	Page            int
	PageSize        int
	EnumerationID   uint64
	EnumerationName uint64
	Opened          bool
	Query           string
}

func (service *EnumerationService) FindQueryEnumerations(form QueryForm) {
	//todo
}

func (dao *EnumerationService) FindEnumLabelsByEnumID(id uint64) (labels []model.EnumerationLabel, err error) {
	if _, err = dao.FindEnumerationByID(id); err != nil {
		return
	}

	err = dao.orm.Where("enumeration_id = ?", id).Find(labels).Error
	return
}

func (dao *EnumerationService) FindOpenedEnumLabelsByEnumID(id uint64) (labels []model.EnumerationLabel, err error) {
	var enum model.Enumeration

	if enum, err = dao.FindEnumerationByID(id); err != nil {
		return
	}

	if enum.Opened == false {
		return nil, errors.New("该枚举尚未开放")
	}

	err = dao.orm.Where("enumeration_id = ?", id).Find(labels).Error
	return
}
