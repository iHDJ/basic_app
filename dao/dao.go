package dao

import (
	"basic_app/dao/model"
	"basic_app/dao/service"
	"basic_app/dao/service/enum"
	"basic_app/dao/service/member"
	"basic_app/dao/service/user"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	Current *Dao
)

type Dao struct {
	*enum.EnumerationService
	*user.UserService
	*member.MemberService

	orm *gorm.DB
}

func New(orm *gorm.DB, redis *redis.Client) (dao *Dao, err error) {
	dao = &Dao{orm: orm}

	if err = dao.init(); err != nil {
		dao = nil
		return
	}
	return
}

func (dao *Dao) init() (err error) {
	err = dao.orm.AutoMigrate(
		model.User{},
		model.Member{},

		model.Group{},
		//model.Organization{},

		model.Enumeration{},
		model.EnumerationLabel{},

		model.Authority{},
		model.AuthorityItem{},
		model.MemberAuthority{},
	).Error

	return
}

func (dao *Dao) NewPaginationService(page int, size int) (s *service.PaginationService) {
	s = service.NewPaginationService(dao.orm)
	s.Page = page
	s.PageSize = size
	return
}
