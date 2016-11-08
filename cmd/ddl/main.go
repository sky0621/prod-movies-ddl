package main

import (
	"flag"
	"log"
	"os"

	"github.com/mitchellh/cli"
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
	ctx, err := md.NewCtx(*configpath)
	if err != nil {
		return md.ExitCodeSetupError
	}
	defer ctx.Close()
	ctx.Logger.Entry.Info("App Start.")

	c := cli.NewCLI("movies-ddl", "0.1.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"create": func() (cli.Command, error) {
			return &md.CreateCommand{Ctx: ctx}, nil
		},
		"drop": func() (cli.Command, error) {
			return &md.DropCommand{Ctx: ctx}, nil
		},
	}
	code, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	return code
}
