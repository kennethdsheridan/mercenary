package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	// Production level logger configuration
	config := zap.NewProductionConfig()

	// Modifying time format from Epoch to ISO8601
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Setting the output format to JSON
	config.Encoding = "json"

	encoderConfig := zap.NewDevelopmentEncoderConfig()

	// Changing the key name for time from "ts" to "timestamp"
	encoderConfig.TimeKey = "timestamp"

	// Modifying time format from Epoch to ISO8601
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.EncoderConfig = encoderConfig

	// Set output path to a file
	config.OutputPaths = []string{"stdout", "appLog.log"}

	// Building the logger
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// Info() is a wrapper for the zap.Info() method for logging informational messages
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug() is a wrapper for the zap.Debug() method for logging debug messages
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error() is a wrapper for the zap.Error() method for logging error messages
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
