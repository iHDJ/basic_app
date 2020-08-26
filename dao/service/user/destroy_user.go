package user

func (service *UserService) DestroyUserByID(id uint64) (err error) {
	user, err := service.FindUserByID(id)
	if err != nil {
		return err
	}

	err = service.orm.Delete(&user).Error
	return
}
