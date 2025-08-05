package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Configs struct {
	AppName        string
	HTTPServerPort string
	PostgresDSN    string
}

func New() *Configs {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(err.Error())
	}

	return &Configs{
		AppName:        os.Getenv("APP_NAME"),
		HTTPServerPort: os.Getenv("HTTP_SERVER_PORT"),
		PostgresDSN:    os.Getenv("DSN_POSTGRES"),
	}
}
