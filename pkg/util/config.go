package util

import "github.com/spf13/viper"

type Config struct {
	// DB map[string]interface{} `mapstructure:"database"`
	DB Database `mapstructure:"database"`
}

type Database struct{
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
