package member

import (
	"basic_app/dao/model"

	"github.com/jinzhu/gorm"
)

type MemberService struct {
	orm *gorm.DB
}

func (service *MemberService) FindMembersByUserID(userID uint64) (members []model.Member, err error) {
	err = service.orm.Where("user_id = ?", userID).Find(&members).Error
	return
}

func (service *MemberService) FindMemberByUserID(groupID, userID uint64) (member model.Member, err error) {
	err = service.orm.Where("user_id = ? and group_id = ?", userID, groupID).First(&member).Error
	return
}
