package logging

import (
	"log"
	"github.com/sirupsen/logrus"
)

func Configure(level string) {
	lvl, err := logrus.ParseLevel(level)

	if err != nil {
		log.Fatal("Error raised when setting log level.", err)
	}

	logrus.SetLevel(lvl)
}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}
