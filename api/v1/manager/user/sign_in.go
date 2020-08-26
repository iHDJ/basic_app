package user

import (
	"basic_app/dao/model"
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	RequestParamErr  = errors.New("请求参数错误")
	PasswordWrongErr = errors.New("密码不正确")
)

type UserSignInRequest struct {
	Type string `json:"method" binding:"required"`

	//when method = 'account'
	Account  string `json:"account"`
	Password string `json:"password"`

	//when method = 'phone'
	Phone    string `json:"phone"`
	AuthCode string `json:"auth_code"`
}

//后台用户登入
func (service *Service) SignIn(c *gin.Context) {
	var (
		request UserSignInRequest
		err     error
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		return
	}

	switch request.Type {
	case "account":
		err = service._PhoneSignIn(c, request)
	case "phone":
		err = service._PasswordSignIn(c, request)
	default:
		err = RequestParamErr
	}

	if err != nil {
		return
	}

}

//手机登入
func (service *Service) _PhoneSignIn(c *gin.Context, request UserSignInRequest) (err error) {
	err = errors.New("未实现")
	return
}

//账户密码登入
func (service *Service) _PasswordSignIn(c *gin.Context, request UserSignInRequest) (err error) {
	user, err := service.dao.FindUserByName(request.Account)
	if err != nil {
		return err
	}

	if err = user.ConfirmPassword(request.Password); err != nil {
		return PasswordWrongErr
	}

	service._GenerateSessionID(user)
	return
}

func (service *Service) _GenerateSessionID(user model.User) (sessionID string) {

	return
}
