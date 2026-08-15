package logger

import "github.com/rs/zerolog"

func SetLogLevel(levelStr string) {
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil || level == zerolog.NoLevel {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
}
