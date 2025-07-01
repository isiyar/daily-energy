package main

import "github.com/spf13/viper"

type Config struct {
	DB_HOST            string `mapstructure:"DB_HOST"`
	DB_PORT            int    `mapstructure:"DB_PORT"`
	DB_USERNAME        string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD        string `mapstructure:"DB_PASSWORD"`
	DB_NAME            string `mapstructure:"DB_NAME"`
	TELEGRAM_BON_TOKEN string `mapstructure:"TELEGRAM_BOT_TOKEN"`
}

func LoadConfig() (c Config, err error) {
	viper.SetConfigFile("../.env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
