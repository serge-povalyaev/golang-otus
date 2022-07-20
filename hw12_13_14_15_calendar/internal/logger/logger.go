package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type CalendarLogger struct {
	logrusLogger *logrus.Logger
	Level        logrus.Level
}

func New(logLevel, filePath string) *CalendarLogger {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Fatal(err)
	}
	log.Level = level

	log.Out = os.Stdout
	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Info("Ошибка в работе с файлом")
		}
	}

	return &CalendarLogger{
		Level:        level,
		logrusLogger: log,
	}
}

func (l *CalendarLogger) getEntry() *logrus.Entry {
	return l.logrusLogger.WithFields(logrus.Fields{})
}

func (l *CalendarLogger) Error(message string) {
	l.getEntry().Error(message)
}

func (l *CalendarLogger) Info(message string) {
	l.logrusLogger.Info(message)
}
