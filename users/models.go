package users

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/Degoke/dekube/common"
)

type UserModel struct {
	ID uint `gorm:"primary_key"`
	Email string `gorm:"column:email;unique_index"`
	PasswordHash string `gorm:"column:password;not null"`
}

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}

	bytePassword := []byte(password)

	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Model Error: (user:setPassword)", err)
		return errors.New("an error occured")
	}
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindOneUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}