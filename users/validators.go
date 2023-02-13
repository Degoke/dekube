package users

import (

	"github.com/Degoke/dekube/common"
	"github.com/gin-gonic/gin"

)

type UserModelValidator struct {
	User struct {
		Email string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (u *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, u)
	if err != nil {
		return err
	}

	u.userModel.Email = u.User.Email
	u.userModel.setPassword(u.User.Password)

	return nil

}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}

type LoginValidator struct {
	User struct {
		Email string `form:"email" json:"email" binding:"required,email"`
		Password string `form:"password" json:"password" binding:"required"`
	} `json:"user"`
	userModel UserModel `json:"-"`
}

func (l *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, l)
	if err != nil {
		return err
	}

	l.userModel.Email = l.User.Email
	l.userModel.setPassword(l.User.Password)

	return nil

}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}