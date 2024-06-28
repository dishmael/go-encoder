package logger

import (
	"fmt"
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

	// Try to find the .env file
	if pwd, err := os.Getwd(); err != nil {
		panic(err)

	} else {
		var envFile string
		paths := [2]string{".env", "../.env"}

		for _, p := range paths {
			if _, err := os.Stat(filepath.Join(pwd, p)); err == nil {
				envFile = filepath.Join(pwd, p)
				break
			}
		}

		// Load environment variables from .env file
		err := godotenv.Load(envFile)
		if err != nil {
			fmt.Printf("[WARN] No .env file found, using default log level: %s\n", DEFAULT_LEVEL)
		}
	}

	// Get the log level from the environment variable
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = DEFAULT_LEVEL // Default log level
	}

	// Parse the log level
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		fmt.Printf("[WARN] Invalid log level in .env file, using default log level: %s\n", DEFAULT_LEVEL)
		level = logrus.InfoLevel
	}

	// Determine path for executable file
	ex, err := os.Executable()
	if err != nil {
		fmt.Printf("[ERROR] Failed to determine executable")
	}

	// Create a log file
	file := filepath.Join(filepath.Dir(ex), "encoder.log")
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("[FATAL] Failed to open log file")
	}

	// Set the output to both stdout and the log file
	//mw := io.MultiWriter(os.Stdout, logFile)
	mw := io.MultiWriter(logFile)
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
