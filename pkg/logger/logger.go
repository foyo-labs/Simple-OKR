package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func ConfigureLogging() *logrus.Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	}

	log.SetLevel(logrus.DebugLevel)

	return log
}
