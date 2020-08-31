package user

import (
	"basic_app/library/result"
	"basic_app/library/sessions"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	PasswordWrongErr = errors.New("2002_user_or_password_wrong")
)

type UserSignInRequest struct {
	Type     string `json:"method" binding:"required"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

//后台用户登入
func (service *Service) SignIn(c *gin.Context) {
	fmt.Println("SignIn")
	var (
		request UserSignInRequest
		err     error
	)

	session := sessions.FindSessionByContext(c)

	fmt.Println("session id", session.ID)

	if err = c.ShouldBindJSON(&request); err != nil {
		result.Error(c, "1002_param_incomplete")
		return
	}
}
