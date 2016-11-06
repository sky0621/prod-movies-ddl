package moviesddl

import (
	"log"
	"os"

	"github.com/Sirupsen/logrus"
)

func SetupLogger(config *LogConfig) error {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logfile, err := os.Create(config.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	// logrus.SetOutput(logfile)

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(level)

	logrus.Println("Test")
	return nil
}
