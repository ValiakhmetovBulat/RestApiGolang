package logger

import (
	"github.com/sirupsen/logrus"
	prefixied "github.com/x-cray/logrus-prefixed-formatter"
	"io"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

func setupFormatter() *prefixied.TextFormatter {
	customFormatter := new(prefixied.TextFormatter)
	customFormatter.TimestampFormat = "01-02 15:04:05"
	customFormatter.FullTimestamp = true

	return customFormatter
}

// SetupLogger sets logger mode according to environment mode
func SetupLogger(env string) {
	switch env {
	case envLocal:
		Logger.SetLevel(logrus.DebugLevel)
		Logger.Formatter = setupFormatter()
	case envDev:
		Logger.SetLevel(logrus.DebugLevel)
		Logger.Formatter = setupFormatter()
	case envProd:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

// Debug logs a debug message
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Debugf logs a formatted debug messsage
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Info logs an informational message
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Infof logs a formatted informational message
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Error logs an error message
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Errorf logs a formatted error message
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Warn logs a warning message
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warnf logs a formatted warning message
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Fatal logs a fatal error message
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Fatalf logs a formatted fatal error message
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// WithFields returns a new log enty with the provided fields
func WithFields(fields logrus.Fields) *logrus.Entry {
	return Logger.WithFields(fields)
}

// Writer returns the current logging writer
func Writer() *io.PipeWriter {
	return Logger.Writer()
}
