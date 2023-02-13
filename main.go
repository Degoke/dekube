package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Degoke/dekube/common"
	"github.com/Degoke/dekube/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {
	common.LoadENV()
	db := common.Init()
	Migrate(db)

	postgresdb, err := db.DB();
	if err != nil {
		log.Fatal("db err: (Main)", err)
	}
	defer postgresdb.Close()

	r := gin.Default()
	v1 := r.Group("/api")

	users.UsersRegister(v1.Group("/users"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))

	port := common.GetENV("PORT")
	r.Run(":" + port)
}