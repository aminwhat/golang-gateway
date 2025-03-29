package log_package

func Info(message string) {
	log := logConfig{logLevel: LOG_LEVEL_INFO}

	log.exec(message)
}

func Warning(message string) {
	log := logConfig{logLevel: LOG_LEVEL_WARNING}

	log.exec(message)
}

func Error(message string) {
	log := logConfig{logLevel: LOG_LEVEL_ERROR}

	log.exec(message)
}

func Debug(message string) {
	log := logConfig{logLevel: LOG_LEVEL_DEBUG}

	log.exec(message)
}
