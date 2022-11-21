package main

import (
	"github.com/spf13/viper"
)

func readConfig() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return vp, nil
}
