package config

import (
	env "github.com/caarlos0/env/v11"
)

type Config struct {
	TgToken string `env:"TG_TOKEN,required,notEmpty"`
}

var Cfg Config

func init() {
	parse()

	err := env.Parse(&Cfg)
	if err != nil {
		panic(err)
	}
}
