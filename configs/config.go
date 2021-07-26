package configs

import (
	"sync"

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
	config *Config
	once   sync.Once
)

func GetInstance() *Config {
	if config == nil {
		once.Do(func() {
			configFile := configFilePath + configFileName
			viper.SetConfigFile(configFile)
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatal().
					Err(err).
					Str("file", configFile).
					Msg("Failed to read configuration file.")
			}

			err = viper.Unmarshal(&config)
			if err != nil {
				log.Fatal().
					Err(err).
					Str("file", configFile).
					Msg("Failed to unmarshal configuration file.")
			}

			log.Info().Msg("Configuration initialized.")
		})
	}

	return config
}
