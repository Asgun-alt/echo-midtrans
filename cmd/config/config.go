package config

import "github.com/spf13/viper"

func InitConfig() error {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
