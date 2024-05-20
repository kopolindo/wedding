package log

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

var (
	textLogger         *slog.Logger
	jsonLogger         *slog.Logger
	slogHandlerOptions *slog.HandlerOptions
	textHandler        *slog.TextHandler
	jsonHandler        *slog.JSONHandler
	logLevel           *slog.LevelVar
)

const (
	loggerFileName     string = "logs.log"
	jsonLoggerFileName string = "logs.json"
)

func init() {
	slogHandlerOptions = &slog.HandlerOptions{}
	logLevel = &slog.LevelVar{} // INFO

	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("error during slog initialization: %s\n", err.Error())
	}
	logDir := filepath.Join(root, "backend-logs")
	err = ensureDir(logDir)
	if err != nil {
		log.Println("error during directory creation")
	}
	loggerFilePath := filepath.Join(logDir, loggerFileName)
	jsonLoggerFilePath := filepath.Join(logDir, jsonLoggerFileName)
	loggerFile, err := os.OpenFile(loggerFilePath, os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Println("error during log file creation/opening")
	}
	jsonLoggerFile, err := os.OpenFile(jsonLoggerFilePath, os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Println("error during log file creation/opening")
	}

	slogHandlerOptions.Level = logLevel

	loggerWriter := io.Writer(loggerFile)
	jsonLoggerWriter := io.Writer(jsonLoggerFile)

	textHandler = slog.NewTextHandler(loggerWriter, slogHandlerOptions)
	jsonHandler = slog.NewJSONHandler(jsonLoggerWriter, slogHandlerOptions)

	textLogger = slog.New(textHandler)
	jsonLogger = slog.New(jsonHandler)
}

func SetSlogLevel(level slog.Level) {
	logLevel.Set(level)
}
func GetSlogLevel() slog.Level {
	return logLevel.Level()
}

func Info(errorMessage error) {
	textLogger.Info(errorMessage.Error())
	jsonLogger.Info(errorMessage.Error())
}

func Debug(errorMessage error) {
	textLogger.Debug(errorMessage.Error())
	jsonLogger.Debug(errorMessage.Error())
}

func Error(errorMessage error) {
	textLogger.Error(errorMessage.Error())
	jsonLogger.Error(errorMessage.Error())
}

func Infof(format string, a ...any) {
	errorMessage := fmt.Sprintf(format, a...)
	textLogger.Info(errorMessage)
	jsonLogger.Info(errorMessage)
}

func Debugf(format string, a ...any) {
	errorMessage := fmt.Sprintf(format, a...)
	textLogger.Debug(errorMessage)
	jsonLogger.Debug(errorMessage)
}

func Errorf(format string, a ...any) {
	errorMessage := fmt.Sprintf(format, a...)
	textLogger.Error(errorMessage)
	jsonLogger.Error(errorMessage)
}
