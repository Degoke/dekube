package users

import (
	"github.com/Degoke/dekube/common"
	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (u *UserSerializer) Response() UserResponse {
	userModel := u.c.MustGet("user_model").(UserModel)
	user := UserResponse{
		ID: userModel.ID,
		Email: userModel.Email,
		Token: common.GenToken(userModel.ID),
	}
	return user
}