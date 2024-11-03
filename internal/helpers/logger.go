package helpers

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
)

type Logger struct {
	logger *log.Logger
}

func getFileName(filePath string) string {
	parts := strings.Split(filePath, "/")
	return parts[len(parts)-1]
}

func NewLogger(writter io.Writer) *Logger {
	logger := log.New(writter, "", log.LstdFlags|log.Ldate|log.Ltime)
	return &Logger{
		logger,
	}
}

func (l *Logger) Info(format string, v ...any) {
	_, filePath, line, _ := runtime.Caller(1)
	file := getFileName(filePath)

	fileLine := fmt.Sprintf("%s:%d:", file, line)

	prefix := "%s INFO: "
	l.logger.Printf(prefix+format, append([]any{fileLine}, v...)...)
}

func (l *Logger) Success(format string, v ...any) {
	_, filePath, line, _ := runtime.Caller(1)
	file := getFileName(filePath)

	fileLine := fmt.Sprintf("%s:%d:", file, line)

	prefix := "%s SUCCESS: "
	l.logger.Printf(prefix+format, append([]any{fileLine}, v...)...)
}

func (l *Logger) Warning(format string, v ...any) {
	_, filePath, line, _ := runtime.Caller(1)
	file := getFileName(filePath)

	fileLine := fmt.Sprintf("%s:%d:", file, line)

	prefix := "%s WARNING: "
	l.logger.Printf(prefix+format, append([]any{fileLine}, v...)...)
}

func (l *Logger) Error(format string, v ...any) {
	_, filePath, line, _ := runtime.Caller(1)
	file := getFileName(filePath)

	fileLine := fmt.Sprintf("%s:%d:", file, line)

	prefix := "%s ERROR: "
	l.logger.Printf(prefix+format, append([]any{fileLine}, v...)...)
}

func (l *Logger) Fatal(format string, v ...any) {
	_, filePath, line, _ := runtime.Caller(1)
	file := getFileName(filePath)

	fileLine := fmt.Sprintf("%s:%d:", file, line)

	prefix := "%s ERROR: "
	l.logger.Fatalf(prefix+format, append([]any{fileLine}, v...)...)
}
