package util

import "github.com/spf13/viper"

type Config struct {
	DB          Database `mapstructure:"database"`
	TelegramBot TGBot    `mapstructure:"telegram_bot"`
}

type Database struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name"`
}

type TGBot struct {
	Token string `mapstructure:"token"`
}

// NOTE need to make load once function

// load config.json
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")

	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
