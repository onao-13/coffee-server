package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func UploadDevConfig() Config {
	v := viper.New()
	v.SetDefault("PORT", "8081")
	v.SetDefault("DBNAME", "test")
	v.SetDefault("DBUSER", "user")
	v.SetDefault("DBPASS", "pass")
	v.SetDefault("DBHOST", "localhost")
	v.SetDefault("DBPORT", "8085")

	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Panic("Error configuration")
	}

	return config
}

func (config *Config) DbUrlConnection() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", config.DbUser,
		config.DbPass, config.DbHost, config.DbPort, config.DbName,
	)
}
