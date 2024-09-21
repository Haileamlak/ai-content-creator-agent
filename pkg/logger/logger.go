package logger

import (
    "log"
    "os"
)

type Logger struct {
    InfoLogger  *log.Logger
    ErrorLogger *log.Logger
    WarnLogger  *log.Logger
}

// NewLogger initializes the logger with standard output and error streams.
func NewLogger() *Logger {
    return &Logger{
        InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
        ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
        WarnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

// Info logs informational messages.
func (l *Logger) Info(msg string) {
    l.InfoLogger.Println(msg)
}

// Error logs error messages.
func (l *Logger) Error(msg string) {
    l.ErrorLogger.Println(msg)
}

// Warn logs warning messages.
func (l *Logger) Warn(msg string) {
    l.WarnLogger.Println(msg)
}
