package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uint64 `gorm:"primary_key"`
	Name           string
	Email          string
	Phone          string
	AvatarURL      string
	PasswordDigest []byte
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
}

func (user *User) ConfirmPassword(password string) (err error) {
	err = bcrypt.CompareHashAndPassword(user.PasswordDigest, []byte(password))
	return
}

func (user *User) SetPassword(password string) (err error) {
	user.PasswordDigest, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return
}
