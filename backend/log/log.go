package log

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
	"syscall"
)

var (
	textLogger         *slog.Logger
	slogHandlerOptions *slog.HandlerOptions
	textHandler        *slog.TextHandler
	logLevel           *slog.LevelVar
)

const (
	loggerFileName     string = "backend-logs.log"
	jsonLoggerFileName string = "backend-logs.json"
	LOGDIRDOCKER       string = "/tmp/logs"
)

func printFolderDetails(folderPath string) {
	fileInfo, err := os.Stat(folderPath)
	if err != nil {
		log.Fatalf("Error retrieving folder info: %v", err)
	}

	// Type assertion to get the underlying data structure.
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatalf("Error getting file stats")
	}

	// Get the owner and group IDs
	uid := stat.Uid
	gid := stat.Gid

	// Get user details
	owner, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		log.Fatalf("Error looking up user: %v", err)
	}

	// Get group details
	group, err := user.LookupGroupId(fmt.Sprint(gid))
	if err != nil {
		log.Fatalf("Error looking up group: %v", err)
	}

	fmt.Printf("Folder: %s\n", folderPath)
	fmt.Printf("Owner: %s (%s)\n", owner.Username, owner.Name)
	fmt.Printf("Group: %s (%s)\n", group.Name, group.Gid)
	fmt.Printf("Permissions: %s\n", fileInfo.Mode())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Last modified: %s\n", fileInfo.ModTime())
}

func printWhoAmI() {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error fetching current user: %v", err)
	}
	fmt.Printf("Current user: %s\n", currentUser.Username)
}

func init() {
	printWhoAmI()
	printFolderDetails(LOGDIRDOCKER)
	slogHandlerOptions = &slog.HandlerOptions{}
	logLevel = &slog.LevelVar{} // INFO
	logFilePath := filepath.Join(LOGDIRDOCKER, loggerFileName)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	slogHandlerOptions.Level = logLevel
	loggerWriter := io.Writer(logFile)
	textHandler = slog.NewTextHandler(loggerWriter, slogHandlerOptions)
	textLogger = slog.New(textHandler)
}

func SetSlogLevel(level slog.Level) {
	logLevel.Set(level)
}
func GetSlogLevel() slog.Level {
	return logLevel.Level()
}

func Info(errorMessage error) {
	textLogger.Info(errorMessage.Error())
}

func Debug(errorMessage error) {
	textLogger.Debug(errorMessage.Error())
}

func Error(errorMessage error) {
	textLogger.Error(errorMessage.Error())
}

func Infof(format string, a ...any) {
	errorMessage := fmt.Sprintf(format, a...)
	textLogger.Info(errorMessage)
}

func Debugf(format string, a ...any) {
	errorMessage := fmt.Sprintf(format, a...)
	textLogger.Debug(errorMessage)
}

func Errorf(format string, a ...any) {
	errorMessage := fmt.Sprintf(format, a...)
	textLogger.Error(errorMessage)
}
