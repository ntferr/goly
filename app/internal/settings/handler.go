package settings

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Settings struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`

	Database struct {
		Name string `env:"DB_NAME"`
		Host string `env:"DB_HOST"`
		Port string `env:"DB_PORT"`
		User string `env:"DB_USER"`
		Pwd  string `env:"DB_PWD"`
	}
}

var settings Settings

func init() {
	if _, err := env.UnmarshalFromEnviron(&settings); err != nil {
		log.Fatal("Error on initiate envs", err)
	}
}

func GetSettings() Settings {
	return settings
}
