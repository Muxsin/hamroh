package configs

type Configs struct {
	AppName        string
	HTTPServerPort string
}

func New() *Configs {
	return &Configs{
		AppName:        "Hamroh Todo",
		HTTPServerPort: "8080",
	}
}
