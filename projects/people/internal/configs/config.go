package configs

import "os"

type Configs struct {
	AppName        string
	HTTPServerPort string
	PostgresDSN    string
}

func New() *Configs {
	return &Configs{
		AppName:        os.Getenv("APP_NAME"),
		HTTPServerPort: os.Getenv("HTTP_SERVER_PORT"),
		PostgresDSN:    os.Getenv("POSTGRES_DSN"),
	}
}
