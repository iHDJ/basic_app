package enum

import (
	"basic_app/dao/model"
	daoService "basic_app/dao/service"
)

func (dao *EnumerationService) FindEnumerationByID(id uint64) (enum model.Enumeration, err error) {
	err = dao.orm.First(enum, id).Error
	return
}

func (dao *EnumerationService) FindEnumerationByName(enumName string) (enum model.Enumeration, err error) {
	err = dao.orm.Where("name = ?", enumName).First(enum).Error
	return
}

//返回所有的枚举
func (service *EnumerationService) FindEnumerations(form EnumerationsQueryForm) (enums []model.Enumeration, err error) {
	s := daoService.NewPaginationService(service.orm)
	s.Page = form.Page
	s.PageSize = form.PageSize

	if form.ID != 0 {
		s.Where("id = ?", form.ID)
		err = s.Execute(&enums)
		return
	}

	if form.Name != "" {
		s.Where("name = ?", form.Name)
	}

	if form.IsSystem {
		s.Where("is_system = ?", true)
	}

	if form.IsOpened {
		s.Where("is_opened = ?", true)
	}

	err = s.Execute(&enums)
	return
}
