package enumeration

import "basic_app/dao"

type Service struct {
	dao *dao.Dao
}

func NewService() *Service {
	return &Service{dao: dao.Current}
}
