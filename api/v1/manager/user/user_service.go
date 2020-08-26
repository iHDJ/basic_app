package user

import (
	"basic_app/dao"

	"github.com/gin-gonic/gin"
)

var _ service = &Service{} //校验接口是否满足

type Service struct {
	dao *dao.Dao
}

func NewService() *Service {
	return &Service{dao: dao.Current}
}

type service interface {
	//获取当前用户信息和后台权限列表
	CurrentUser(c *gin.Context)

	//用户登陆，登入成功后返回user_id并设置session
	SignIn(c *gin.Context)

	//用户注册
	SignUp(c *gin.Context)
}
