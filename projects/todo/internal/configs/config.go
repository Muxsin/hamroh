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
		AppName:        "Hamroh Todo",
		HTTPServerPort: "8080",
		PostgresDSN:    os.Getenv("DSN_POSTGRES"),
	}
}
