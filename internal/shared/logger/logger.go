package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// defaultLogLevel default log level
	defaultLogLevel = zerolog.TraceLevel
)

// InitLogger initializes the logger
func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(defaultLogLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = formatLevel
	output.FormatMessage = formatMessage
	output.FormatFieldName = formatFieldName
	output.FormatFieldValue = formatFieldValue

	log.Logger = log.Output(output)
	log.Info().Msg("Logger initialized.")
}

// formatLevel sets format log level
func formatLevel(
	i interface{},
) string {
	return strings.ToUpper(fmt.Sprintf("| %-5s |", i))
}

// formatMessage sets format log message
func formatMessage(
	i interface{},
) string {
	return fmt.Sprintf("%s", i)
}

// formatFieldName sets format log field name
func formatFieldName(
	i interface{},
) string {
	return fmt.Sprintf("%s=", i)
}

// formatFieldValue sets format log field value
func formatFieldValue(
	i interface{},
) string {
	return fmt.Sprintf("%s", i)
}

// SetLogLevel sets global log level
func SetLogLevel(
	logLevel string,
) {
	trimmedLevel := strings.TrimSpace(logLevel)
	level, err := zerolog.ParseLevel(trimmedLevel)
	if trimmedLevel == "" || err != nil {
		level = defaultLogLevel
		log.Debug().
			Str("logLevel", level.String()).
			Msg("Application has no log level configuration.")
	} else {
		log.Debug().
			Str("logLevel", level.String()).
			Msg("Log level has been set up.")
	}
	zerolog.SetGlobalLevel(level)
}
