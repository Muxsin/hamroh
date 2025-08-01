package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	AppName        string
	HTTPServerPort string
	PostgresDSN    string
}

func New() *Configs {
	godotenv.Load(".env")
	
	return &Configs{
		AppName:        "Hamroh Todo",
		HTTPServerPort: "8080",
		PostgresDSN:    os.Getenv("DSN_POSTGRES"),
	}
}
