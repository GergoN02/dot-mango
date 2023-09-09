package logger

import (
	"log"
	"os"
	"sync"
)

type LogLevel string

const (
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
	Error LogLevel = "ERROR"
	Debug LogLevel = "DEBUG"
)

var (
	defaultLogger *FileLogger
	initOnce      sync.Once
)

type FileLogger struct {
	logger *log.Logger
	mu     sync.Mutex
}

func Init(filename string) error {
	var err error
	initOnce.Do(func() {
		defaultLogger, err = newFileLogger(filename)
	})
	return err
}

func newFileLogger(filename string) (*FileLogger, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	l := log.New(f, "", log.LstdFlags)
	return &FileLogger{logger: l}, nil
}

func Log(message string, level LogLevel) {
	defaultLogger.Log(message, level)
}

func (l *FileLogger) Log(message string, level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Printf("[%s] %s", level, message)
}
