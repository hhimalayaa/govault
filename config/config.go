package config

type Config struct {
	APP APP `envConfig:"APP"`
}

type APP struct {
	NAME string `envConfig:"APP_NAME"`
	PORT string `envConfig:"PORT_NO"`
}
