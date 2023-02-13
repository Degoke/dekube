package common

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {
	dbPassword := GetENV("DB_PASSWORD")
	dbHost := GetENV("DB_HOST")
	dbUser := GetENV("DB_USER")
	dbName := GetENV("DB_NAME")
	dbPort := GetENV("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init)", err)
	}

	postgresdb, err := db.DB()
	if err != nil {
		fmt.Println("db err: (Init)", err)
	}
	postgresdb.SetMaxIdleConns(10)
	
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}