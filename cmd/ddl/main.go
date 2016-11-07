package main

import (
	"flag"
	"log"
	"os"

	md "github.com/sky0621/prod-movies-ddl"

	"github.com/BurntSushi/toml"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	// treat panic

	return wrappedMain()
}

func wrappedMain() int {

	configpath := flag.String("f", "./config.toml", "設定ファイル（config.toml）の格納パス")
	var config md.Config
	_, err := toml.DecodeFile(*configpath, &config)
	if err != nil {
		log.Fatal(err)
	}

	logger, logfile := md.SetupLogger(&config)
	defer logfile.Close()
	logger.Info("App Start.")

	return md.ExitCodeOK
}
