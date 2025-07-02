package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Debug            bool   `mapstructure:"DEBUG"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           int    `mapstructure:"DB_PORT"`
	DBUsername       string `mapstructure:"DB_USERNAME"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
}

func LoadConfig() (Config, error) {
	var c Config

	viper.AutomaticEnv()

	err := viper.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return c, nil
}
