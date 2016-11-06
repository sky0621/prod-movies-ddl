package main

import (
	"flag"
	"log"

	md "github.com/sky0621/prod-movies-ddl"

	"github.com/BurntSushi/toml"
)

func main() {
	configpath := flag.String("f", "./config.toml", "設定ファイル（config.toml）の格納パス")
	var config md.Config
	_, err := toml.DecodeFile(*configpath, &config)
	if err != nil {
		log.Fatal(err)
	}

	logger, logfile := md.SetupLogger(&config)
	defer logfile.Close()
	logger.Info("App Start.")

}
