package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	// configFileName config file name
	configFileName = ".env"
	// configFilePath path to config file
	configFilePath = "./"
)

type Config struct {
	DB struct {
		MySQL struct {
			Read struct {
				Host     string `mapstructure:"HOST"`
				Name     string `mapstructure:"NAME"`
				Password string `mapstructure:"PASSWORD"`
				Port     int    `mapstructure:"PORT"`
				User     string `mapstructure:"USER"`
			} `mapstructure:"READ"`

			Write struct {
				Host     string `mapstructure:"HOST"`
				Name     string `mapstructure:"NAME"`
				Password string `mapstructure:"PASSWORD"`
				Port     int    `mapstructure:"PORT"`
				User     string `mapstructure:"USER"`
			} `mapstructure:"WRITE"`
		} `mapstructure:"MYSQL"`
	} `mapstructure:"DB"`

	Server struct {
		LogLevel string `mapstructure:"LOG_LEVEL"`
	} `mapstructure:"SERVER"`
}

var (
	// instance singleton object of configuration
	instance *Config
)

// GetInstance get singleton object of configuration
func GetInstance() *Config {
	if instance == nil {
		filePath := configFilePath + configFileName
		GetConfig(filePath, &instance)
	}

	return instance
}

// GetConfig unmarshal configuration file into `config` structure
func GetConfig(
	filePath string,
	config interface{},
) {
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
