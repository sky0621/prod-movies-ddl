package moviesddl

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// Logger ...
type Logger struct {
	Entry   *logrus.Entry
	logfile *os.File
}

// NewLogger ...
func NewLogger(config *Config) (*Logger, error) {
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
		return nil, err
	}
	logrusEntry.Logger.Out = logfile

	level, err := logrus.ParseLevel(config.Log.LogLevel)
	if err != nil {
		return nil, err
	}
	logrusEntry.Logger.Level = level

	return &Logger{Entry: logrusEntry, logfile: logfile}, nil
}

// Close ...
func (l *Logger) Close() {
	l.logfile.Close()
}

// SetupLogger ... ログ設定の初期化
// FIXME ... 使わなくなったので消す！
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
