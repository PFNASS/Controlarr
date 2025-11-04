package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/Projects/GoProjects/go-arrs/go-arrs/pkg/config")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}