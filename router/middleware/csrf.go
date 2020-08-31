package middleware

import (
	"basic_app/library/sessions"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/dchest/uniuri"
	"github.com/gin-gonic/gin"
)

const (
	CsrfSecret = "csrf_secret"
	CsrfSalt   = "csrf_salt"
	CsrfToken  = "csrf_token"
)

var defaultIgnoreMethods = []string{"GET", "HEAD", "OPTIONS"}

var defaultErrorFunc = func(c *gin.Context) {
	panic(errors.New("CSRF token mismatch"))
}

var defaultTokenGetter = func(c *gin.Context) string {
	r := c.Request

	if authorization := c.Request.Header.Get("Authorization"); authorization != "" {
		return authorization
	}

	return r.Header.Get("X-CSRF-TOKEN")
}

// Options stores configurations for a CSRF middleware.
type CsrfOptions struct {
	Secret        string
	IgnoreMethods []string
	ErrorFunc     gin.HandlerFunc
	TokenGetter   func(c *gin.Context) string
}

func tokenize(secret, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt+"-"+secret)

	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func inArray(arr []string, value string) bool {
	inarr := false

	for _, v := range arr {
		if v == value {
			inarr = true
			break
		}
	}

	return inarr
}

// Middleware validates CSRF token.
func CSRF(options CsrfOptions) gin.HandlerFunc {
	ignoreMethods := options.IgnoreMethods
	errorFunc := options.ErrorFunc
	tokenGetter := options.TokenGetter

	if ignoreMethods == nil {
		ignoreMethods = defaultIgnoreMethods
	}

	if errorFunc == nil {
		errorFunc = defaultErrorFunc
	}

	if tokenGetter == nil {
		tokenGetter = defaultTokenGetter
	}

	return func(c *gin.Context) {
		session := sessions.FindSessionByContext(c)
		c.Set(CsrfSecret, options.Secret)

		if inArray(ignoreMethods, c.Request.Method) {
			c.Next()
			return
		}

		salt := session.Get(CsrfSalt)
		fmt.Println("salt", salt)

		if salt == "" {
			errorFunc(c)
			return
		}

		token := tokenGetter(c)

		if tokenize(options.Secret, salt) != token {
			errorFunc(c)
			return
		}

		c.Next()
	}
}

// GetToken returns a CSRF token.
func GetCSRFToken(c *gin.Context) string {
	session := sessions.FindSessionByContext(c)
	secret := c.MustGet(CsrfSecret).(string)
	fmt.Println("secret", secret)

	if t, ok := c.Get(CsrfToken); ok {
		return t.(string)
	}

	salt := session.Get(CsrfSalt)
	if salt == "" {
		fmt.Println("salt is empty")
		salt = uniuri.New()
		session.Set(CsrfSalt, salt)
		session.Save(15 * time.Minute) //因为这个会话是新会话且没登入过的，所以默认让他保持15分钟的有效时间
	}

	token := tokenize(secret, salt)
	c.Set(CsrfToken, token)

	return token
}
