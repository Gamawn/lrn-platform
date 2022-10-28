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
	}
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

	config.Database.Password = viper.GetString("database.password")
	config.Database.Username = viper.GetString("database.user")
	config.Database.Host = viper.GetString("database.host")
	config.Database.DbName = viper.GetString("database.dbname")

	return config, nil
}
