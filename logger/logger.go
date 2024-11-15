package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string, fields map[string]interface{})
	Warn(msg string)
	Error(err error, msg string)
	Fatal(err error, msg string)
}

type zeroLogger struct {
	log zerolog.Logger
}

func New() *zeroLogger {
	return &zeroLogger{
		log: zerolog.New(os.Stdout).With().Timestamp().Caller().Logger(),
	}
}

func (l *zeroLogger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *zeroLogger) Info(msg string, fields map[string]interface{}) {
	event := l.log.Info()

	if fields != nil {
		for key, value := range fields {
			event = event.Interface(key, value)
		}
	}

	event.Msg(msg)
}

func (l *zeroLogger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *zeroLogger) Error(err error, msg string) {
	l.log.Error().Err(err).Msg(msg)
}

func (l *zeroLogger) Fatal(err error, msg string) {
	l.log.Fatal().Err(err).Msg(msg)
}
