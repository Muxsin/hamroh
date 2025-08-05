package configs

import "os"

type configs struct {
	AppName        string
	HTTPServerPort string
	PostgresDSN    string
}

func New() *configs {
	return &configs{
		AppName:        os.Getenv("APP_NAME"),
		HTTPServerPort: os.Getenv("HTTP_SERVER_PORT"),
		PostgresDSN:    os.Getenv("POSTGRES_DSN"),
	}
}
