package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initialize the logger
func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = formatLevel
	output.FormatMessage = formatMessage
	output.FormatFieldName = formatFieldName
	output.FormatFieldValue = formatFieldValue

	log.Logger = log.Output(output)
	log.Trace().Msg("Logger initialized.")
}

// formatLevel set format log level
func formatLevel(i interface{}) string {
	return strings.ToUpper(fmt.Sprintf("| %-5s |", i))
}

// formatMessage set format log message
func formatMessage(i interface{}) string {
	return fmt.Sprintf("%s", i)
}

// formatFieldName set format log field name
func formatFieldName(i interface{}) string {
	return fmt.Sprintf("%s=", i)
}

// formatFieldValue set format log field value
func formatFieldValue(i interface{}) string {
	return fmt.Sprintf("%s", i)
}
