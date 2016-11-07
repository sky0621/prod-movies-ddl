package main

import (
	"flag"
	"log"
	"os"

	md "github.com/sky0621/prod-movies-ddl"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	// treat panic
	defer func() {
		err := recover()
		if err != nil {
			log.Println("Panic occured.")
			// FIXME 後始末

		}
	}()

	return wrappedMain()
}

func wrappedMain() int {

	configpath := flag.String("f", "./config.toml", "設定ファイル（config.toml）の格納パス")
	flag.Parse()
	config, err := md.NewConfig(*configpath)
	if err != nil {
		return md.ExitCodeConfigError
	}

	logger, logfile := md.SetupLogger(config)
	defer logfile.Close()
	logger.Info("App Start.")

	return md.ExitCodeOK
}
