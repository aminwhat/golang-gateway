package log_package

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	LOG_LEVEL_INFO    = "info"
	LOG_LEVEL_WARNING = "warning"
	LOG_LEVEL_ERROR   = "error"
	LOG_LEVEL_DEBUG   = "debug"
)

type logConfig struct {
	logLevel string
}

func (config *logConfig) exec(message string) {
	var completeMessage string

	switch config.logLevel {
	case LOG_LEVEL_INFO:
		completeMessage = color.BlueString("LOG") + ": " + message
	case LOG_LEVEL_WARNING:
		completeMessage = color.YellowString("LOG") + ": " + message
	case LOG_LEVEL_ERROR:
		completeMessage = color.RedString("LOG") + ": " + message
	case LOG_LEVEL_DEBUG:
		completeMessage = color.GreenString("LOG") + ": " + message
	default:
		completeMessage = color.CyanString("Custom "+config.logLevel) + ": " + message
	}

	fmt.Println(completeMessage)
}
