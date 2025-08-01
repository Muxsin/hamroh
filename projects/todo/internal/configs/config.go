package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type configs struct {
	AppName        string
	HTTPServerPort string
	PostgresDSN    string
}

func New() *configs {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(err.Error())
	}

	return &configs{
		AppName:        "Hamroh Todo",
		HTTPServerPort: "8080",
		PostgresDSN:    os.Getenv("DSN_POSTGRES"),
	}
}
