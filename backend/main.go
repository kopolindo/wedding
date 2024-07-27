package main

import (
	"flag"
	"log/slog"
	"wedding/api"
	"wedding/log"
)

var (
	debugFlag = flag.Bool("debug", true, "Debug messages enabled")
)

func init() {
	flag.Parse()
	if *debugFlag {
		log.SetSlogLevel(slog.LevelDebug)
	}
}

func main() {
	log.Infof("v1.0")
	log.Debugf("debug enabled")
	log.Infof("starting server")
	log.Error(api.App.Listen(":8080"))
}
