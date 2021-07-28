package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	configFileName = ".env"
	configFilePath = "./"
)

type Config struct {
	Server struct {
		LogLevel string `mapstructure:"LOG_LEVEL"`
	} `mapstructure:"SERVER"`
}

var (
	instance *Config
)

func GetInstance() *Config {
	if instance == nil {
		filePath := configFilePath + configFileName
		GetConfig(filePath, &instance)
	}

	return instance
}

func GetConfig(filePath string, config interface{}) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("file", filePath).
			Msg("Failed to read configuration file.")
	}

	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("file", filePath).
			Msg("Failed to unmarshal configuration file.")
	}

	log.Info().Msg("Configuration initialized.")
}
