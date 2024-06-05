package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DatabaseUrl string `mapstructure:"DATABASE_URL"`
}

var Config Env

func InitEnv() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %s", err)
	}
}
