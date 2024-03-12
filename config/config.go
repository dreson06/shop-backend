package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

var Cfg Config

type Mode string

const (
	ModeRelease Mode = "release"
	ModeDev     Mode = "dev"
)

type Config struct {
	Port        string `env:"port" env-default:"8610"`
	PostgresUrl string `env:"postgres_url"`
	AccessToken string `env:"access_token"`
	Mode        Mode   `env:"mode"`
}

func init() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		panic(err)
	}
	if Cfg.AccessToken == "" {
		panic("access_token_secret missing")
	}

	if IsRelease() {
		//setting UTC time globally.
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			panic(err)
		}
		time.Local = loc
	}
}

func IsRelease() bool {
	return Cfg.Mode == ModeRelease
}
