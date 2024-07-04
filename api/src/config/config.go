package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConectBank = ""
	Door       = 0
	SecretKey  []byte
)

func Load() {
	var err error
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Door, err = strconv.Atoi(os.Getenv("API_DOOR"))
	if err != nil {
		Door = 7070
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	ConectBank = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_KEY"),
		os.Getenv("DB_NAME"),
	)
}
