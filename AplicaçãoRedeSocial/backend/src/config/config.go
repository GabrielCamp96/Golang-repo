package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DbConnection string = ""
	ApiPort      int    = 0
)

func LoadEnv() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(fmt.Sprintf("Erro no carregamento do ambiente: %s", erro))
	}

	ApiPort, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		ApiPort = 9000
	}

	DbConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"))
}
