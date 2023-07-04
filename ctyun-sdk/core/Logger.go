package core

import (
	"log"
)

const (
	LogFatal = iota
	LogError
	LogWarn
	LogInfo
)

type Logger interface {
	Log(level int, message ...interface{})
}

type DefaultLogger struct {
	Level int
}

func NewDefaultLogger(level int) *DefaultLogger {
	return &DefaultLogger{level}
}

func (logger DefaultLogger) Log(level int, message ...interface{}) {
	if level <= logger.Level {
		log.Println(message...)
	}
}
