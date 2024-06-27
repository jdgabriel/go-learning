package config

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

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "[DEBUG]: ", logger.Flags()),
		info:    log.New(writer, "[INFO]: ", logger.Flags()),
		warning: log.New(writer, "[WARNING]: ", logger.Flags()),
		err:     log.New(writer, "[ERROR]: ", logger.Flags()),
		writer:  writer,
	}
}

/* ===== Funções auxiliares para uso do Logger ===== */

// Logs não formatados
func (log *Logger) Debug(v ...interface{}) {
	log.debug.Println(v...)
}

func (log *Logger) Info(v ...interface{}) {
	log.info.Println(v...)
}

func (log *Logger) Warning(v ...interface{}) {
	log.warning.Println(v...)
}

func (log *Logger) Error(v ...interface{}) {
	log.err.Println(v...)
}

// Logs formatados
func (log *Logger) Debugf(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}

func (log *Logger) Infof(format string, v ...interface{}) {
	log.info.Printf(format, v...)
}

func (log *Logger) Warningf(format string, v ...interface{}) {
	log.warning.Printf(format, v...)
}

func (log *Logger) Errorf(format string, v ...interface{}) {
	log.err.Printf(format, v...)
}
