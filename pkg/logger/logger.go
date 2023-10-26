package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
	logFile *os.File
}

func (l *Logger) CLose() {
	l.logFile.Close()
}

// NewLogger create a new logger
func NewLogger(logFile bool, logPath string) *Logger {
	var out io.Writer
	if logFile {
		// log to file
		logDir := filepath.Dir(logPath)

		CreateDirIfNotExist := func(dir string) error {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				return os.MkdirAll(dir, 0o777)
			}
			return nil
		}

		if err := CreateDirIfNotExist(logDir); err != nil {
			panic(err)
		}
		var (
			logfile *os.File
			err     error
		)
		logfile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		out = logfile
	} else {
		out = os.Stdout
	}

	logger := &Logger{}
	if logFile {
		logger.logFile = out.(*os.File)
	}
	logger.Logger = zerolog.New(zerolog.ConsoleWriter{
		NoColor:    true,
		Out:        out,
		TimeFormat: "2006-01-02 15:04:05",
	}).With().Timestamp().Logger()
	return logger
}
