package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type LogWriter struct {
	Logger    *logrus.Logger
	Level     logrus.Level
	Component string
}

func (d LogWriter) Write(p []byte) (n int, err error) {
	var entry *logrus.Entry

	if d.Logger == nil {
		entry = logrus.WithField("component", d.Component)
	} else {
		entry = d.Logger.WithField("component", d.Component)
	}

	entry.Log(d.Level, strings.TrimRight(string(p), "\n"))

	return len(p), nil
}
