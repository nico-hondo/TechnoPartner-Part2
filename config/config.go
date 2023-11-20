package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(filename string) *viper.Viper {
	config := viper.New()

	config.SetConfigName(filename)

	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error while parsing configuration file", err)
	}

	return config
}
