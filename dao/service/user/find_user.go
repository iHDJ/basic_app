package user

import (
	"basic_app/dao/model"
	daoService "basic_app/dao/service"
)

func (service *UserService) FindUserByName(account string) (user model.User, err error) {
	err = service.orm.Where("name = ?", account).First(&user).Error
	return
}

func (service *UserService) FindUserByID(id uint64) (user model.User, err error) {
	err = service.orm.First(&user, id).Error
	return
}

func (service *UserService) FindUsers(page int, pageSize int) (users []model.User, err error) {
	s := daoService.NewPaginationService(service.orm)
	s.Page = page
	s.PageSize = pageSize

	err = s.Execute(&users)
	return
}
