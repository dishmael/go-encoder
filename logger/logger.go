package logger

import (
	"io"
	"os"
	"path/filepath"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

// Initialize the logger by reading the LOG_LEVEL variable from the .env file.
// If the .env file does not exist or the LOG_LEVEL is not properly defined,
// we will default to 'debug'. We are also hardcoded to log to both STDOUT and
// a log file named encoder.log.
func init() {
	const DEFAULT_LEVEL = "debug"

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		Logger.Warnf("No .env file found, using default log level: %s\n", DEFAULT_LEVEL)
	}

	// Get the log level from the environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = DEFAULT_LEVEL // Default log level
	}

	// Parse the log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		Logger.Warnf("Invalid log level in .env file, using default log level: %s\n", DEFAULT_LEVEL)
		level = logrus.InfoLevel
	}

	// Determine path for log file
	ex, err := os.Executable()
	if err != nil {
		Logger.Fatalf("Failed to determine executable: %v", err)
	}

	// Create a log file
	filePath := filepath.Dir(ex) + "/encoder.log"
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatalf("Failed to open log file: %v", err)
	}

	// Set the output to both stdout and the log file
	mw := io.MultiWriter(os.Stdout, logFile)
	Logger.SetOutput(mw)

	// Set the log level (can be changed to logrus.DebugLevel, logrus.InfoLevel, etc.)
	Logger.SetLevel(level)

	// Set the format for our logging
	Logger.SetFormatter(&nested.Formatter{
		HideKeys:     true,
		NoColors:     true,
		TrimMessages: true,
		FieldsOrder:  []string{"component", "category"},
	})
}
