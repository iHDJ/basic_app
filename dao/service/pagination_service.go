package service

import "github.com/jinzhu/gorm"

const (
	DefaultPageSize = 25
)

type Service interface {
	Execute(data interface{}) (err error)
}

type PaginationService struct {
	orm      *gorm.DB
	Page     int
	PageSize int
}

func NewPaginationService(orm *gorm.DB) *PaginationService {
	return &PaginationService{orm: orm}
}

func (service *PaginationService) Where(query string, args ...interface{}) {
	service.orm = service.orm.Where(query, args)
	return
}

func (service *PaginationService) Execute(data interface{}) (err error) {
	offset := service.PageSize * (service.Page - 1)

	err = service.orm.Limit(25).Offset(offset).Find(data).Error
	return
}
