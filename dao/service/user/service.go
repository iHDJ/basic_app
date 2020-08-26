package user

import (
	"github.com/jinzhu/gorm"
)

type UserService struct {
	orm *gorm.DB
}

func NewService(orm *gorm.DB) *UserService {
	return &UserService{orm: orm}
}
