package common

import (
	"log"

	"github.com/joho/godotenv"
)

var myEnv map[string]string

func LoadENV() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	myEnv = env
}

func GetENV(key string) string {
	return myEnv[key]
}