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
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBNAME", "")
	v.SetDefault("DBUSER", "")
	v.SetDefault("DBPASS", "")
	v.SetDefault("DBHOST", "")
	v.SetDefault("DBPORT", "")

	var config Config

	err := v.Unmarshal(&config)
	if err != nil {
		log.Panic("Error configuration")
	}

	return config
}

func DbUrlConnection(config Config) string {
	return fmt.Sprintf(
		"h2://%s%s@%s%s/%s", config.DbUser, config.DbPass,
		config.DbHost, config.DbPort, config.DbName)
}
