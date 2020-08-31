package manager

import (
	"basic_app/dao"
	"basic_app/dao/model"
	"basic_app/locales"
)

type PasswordSignInService struct {
	dao *dao.Dao

	Username string
	Password string
}

func (service *PasswordSignInService) verify() (err error) {
	if service.Username == "" || service.Password == "" {

		return
	}

	return
}

func (service *PasswordSignInService) Execute() (err error) {
	var (
		user model.User
		//member model.Member
	)

	if err = service.verify(); err != nil {
		err = locales.Error("10001_no_query_param")
		return
	}

	if user, err = service.dao.FindUserByName(service.Username); err != nil {
		err = locales.Error("20001_user_or_password_wrong")
		return
	}

	if err = user.ConfirmPassword(service.Password); err != nil {
		err = locales.Error("20001_user_or_password_wrong")
		return
	}

	// member, err := service.dao.FindMemberByUserID(model.ManagerGroup.ID, user.ID)
	// if err != nil {
	// 	err = locales.Error("20001_user_or_password_wrong")
	// 	return
	// }

	return
}
