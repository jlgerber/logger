package logger

import (
	"fmt"
	"os"
)

type Logger struct {
	Level LogLevel
}

func (l *Logger) Debug(i ...interface{}) {
	if l.Level <= DEBUG {

		fmt.Printf("[DEBUG] - %s\n", fmt.Sprint(i...))
	}
}

func (l *Logger) Info(i ...interface{}) {
	if l.Level <= INFO {

		fmt.Printf("[INFO]  - %s\n", fmt.Sprint(i...))
	}
}

func (l *Logger) Warning(i ...interface{}) {
	if l.Level <= WARNING {

		fmt.Printf("[WARN]  - %s\n", fmt.Sprint(i...))
	}
}

func (l *Logger) Error(i ...interface{}) {
	if l.Level <= ERROR {

		fmt.Printf("[ERROR] - %s\n", fmt.Sprint(i...))
	}
}

func (l *Logger) Fatal(i ...interface{}) {
	if l.Level <= FATAL {

		fmt.Printf("[FATAL] - %s\n", fmt.Sprint(i...))
		os.Exit(1)
	}
}

func (l *Logger) SetLevel(level string) {
	l.Level = levelFromString(level)
}

var log *Logger

// GetLogger - return a pointer to Logger
func GetLogger() *Logger {
	if log == nil {

		log = &Logger{WARNING}
	}
	return log
}
