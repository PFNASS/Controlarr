package logging

import (
	// "os"
	// "context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func InitLogger(level string, output string) (*Logger, error) {
	
	Log := &Logger{}

	switch level {
		case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}

	// to optimise disable console writer on prod JSON could be mucheasier to parse
	if output == "stdout" && os.Getenv("ENV") != "production" {
		Log.logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false}).With().Timestamp().Logger()
	} else {
		switch output {
			case "stdout":
				Log.logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
			case "stderr":
				Log.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
			default:
				f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				Log.logger = zerolog.New(f).With().Timestamp().Logger()
		}
	}

	return Log, nil
}

func (Log *Logger) Debug(msg string) {
	Log.logger.Debug().Msg(msg)
}


func (Log *Logger) Info(msg string) {
	Log.logger.Info().Msg(msg)
}


func (Log *Logger) Infof(format string, v ...interface{}) {
	Log.logger.Info().Msgf(format, v...)
}

func (Log *Logger) Warn(msg string) {
	Log.logger.Warn().Msg(msg)
}

func (Log *Logger) Error(msg string, err error) {
	Log.logger.Err(err).Msg(msg)
}

func (Log *Logger) Fatal(msg string, err error) {
	Log.logger.Fatal().Err(err).Msg(msg)
}

func (Log *Logger) Panic(msg string, err error) {
	Log.logger.Panic().Err(err).Msg(msg)
}