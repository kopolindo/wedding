package main

import (
	"flag"
	"log/slog"
	"wedding/api"
	"wedding/log"
)

var (
	debugFlag = flag.Bool("debug", false, "Debug messages enabled")
)

func init() {
	flag.Parse()
	if *debugFlag {
		log.SetSlogLevel(slog.LevelDebug)
	}
}

func main() {
	log.Debugf("debug enabled")
	log.Error(api.App.Listen(":8080"))
}
