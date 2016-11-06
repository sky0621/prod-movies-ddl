package moviesddl

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// SetupLogger ... ログ設定の初期化
func SetupLogger(config *Config) (*logrus.Entry, *os.File) {
	logrusEntry := logrus.WithFields(logrus.Fields{
		"host":   config.Server.Host,
		"system": "movies-ddl",
	})
	logrusEntry.Logger.Formatter = new(logrus.TextFormatter)
	logrusEntry.Logger.Formatter = new(logrus.JSONFormatter) // default

	_, err := os.Stat(config.Log.Filepath)
	var logfile *os.File
	if err == nil {
		logfile, err = os.OpenFile(config.Log.Filepath, os.O_APPEND, 0666)
	} else {
		logfile, err = os.Create(config.Log.Filepath)
	}
	if err != nil {
		panic(err)
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(config.Log.LogLevel)
	if err != nil {
		panic(err)
	}
	logrusEntry.Logger.Level = level
	return logrusEntry, logfile
}
