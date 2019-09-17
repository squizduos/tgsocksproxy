package main

type AppConfig struct {
	Debug    bool   `envconfig:"debug"`
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"1080"`
	User     string
	Password string
	Restrict bool `default:"true"`
}
