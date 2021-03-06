package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"basic_app/router/middleware"

	"github.com/gin-gonic/gin"
)

var (
	csrfParam   = `<meta name="csrf-param" content="authenticity_token">`
	csrfToken   = `<meta name="csrf-token" content="%s">`
	csrfPattern = regexp.MustCompile(`<head>`)
)

func FrontendHandle(c *gin.Context) {
	if c.Request.Method != "GET" {
		return
	}

	var token = middleware.GetCSRFToken(c)
	body, err := getFrontendResponse()

	if err != nil {
		fmt.Println(err)
		return
	}

	if c.Request.URL.Path == "/404" {
		c.Redirect(301, "/403")
		return
	}

	c.Data(
		200,
		"text/html;charset=utf-8",
		csrfPattern.ReplaceAll(body, getCsrfElement(token)),
	)
}

func getCsrfElement(csrfToken string) []byte {
	csrfStr := csrfParam + `<meta name="csrf-token" content="` + csrfToken + `">`

	return []byte(csrfStr)
}

func getFrontendResponse() (body []byte, err error) {
	resp, err := http.Get("http://ant-design-pro.basic_app.docker:8000/404")

	if err == nil {
		defer resp.Body.Close()
	}

	return ioutil.ReadAll(resp.Body)
}
