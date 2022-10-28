package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		DbName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
