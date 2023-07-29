package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
func Fatal(msg string) {
	log.Fatal().Msg(msg)
	os.Exit(1)
}
func FatalF(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
	os.Exit(1)
}
func Error(msg string) {
	log.Error().Msg(msg)
}
func ErrorF(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}
func Warn(msg string) {
	log.Warn().Msg(msg)
}
func WarnF(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}
func Info(msg string) {
	log.Info().Msg(msg)
}
func InfoF(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}
