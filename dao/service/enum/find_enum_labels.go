package enum

import (
	"basic_app/dao/model"
	"basic_app/dao/service"
)

func (dao *EnumerationService) FindEnumLabels(form QueryForm) (labels []model.EnumerationLabel, total int, err error) {
	if err = dao.orm.Model(model.EnumerationLabel{}).Count(&total).Error; err != nil {
		return
	}

	s := service.NewPaginationService(dao.orm)
	s.Page = form.Page
	s.PageSize = form.PageSize

	if form.EnumerationID > 0 {
		s.Where("enumeration_id = ?", form.EnumerationID)
	}

	if form.Query != "" {
		s.Where("concat(label, value) = ?", form.EnumerationID)
	}

	err = s.Execute(&labels)
	return
}
