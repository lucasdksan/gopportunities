package tools

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate)

	return &Logger{
		debug:   log.New(writer, "Debug ", logger.Flags()),
		info:    log.New(writer, "Info ", logger.Flags()),
		warning: log.New(writer, "Warning ", logger.Flags()),
		err:     log.New(writer, "Err ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(v interface{}) {
	l.debug.Println(v)
}

func (l *Logger) Info(v interface{}) {
	l.info.Println(v)
}

func (l *Logger) Warning(v interface{}) {
	l.warning.Println(v)
}

func (l *Logger) Err(v interface{}) {
	l.err.Println(v)
}

func (l *Logger) Debugf(format string, v interface{}) {
	l.debug.Printf(format, v)
}

func (l *Logger) Infof(format string, v interface{}) {
	l.info.Printf(format, v)
}

func (l *Logger) Warningf(format string, v interface{}) {
	l.warning.Printf(format, v)
}

func (l *Logger) Errf(format string, v interface{}) {
	l.err.Printf(format, v)
}
