package main

import (
	"errors"
	"flag"
	"log/slog"
	"wedding/api"
	"wedding/database"
	"wedding/log"
)

var (
	debugFlag = flag.Bool("debug", true, "Debug messages enabled")
	newUser   = flag.Bool("NewUser", false, "Create a new user")
	username  = flag.String("Username", "", "Username for the new user")
	password  = flag.String("Password", "", "Password for the new user")
	role      = flag.String("Role", "", "Role for the new user")
)

func init() {
	flag.Parse()
	if *debugFlag {
		log.SetSlogLevel(slog.LevelDebug)
	}
}

func validateNewUserFlags() error {
	if *newUser && (*username == "" || *password == "" || *role == "") {
		return errors.New("all --NewUser, --Username, --Password, and --Role flags are required together")
	}
	return nil
}

func main() {
	log.Infof("v1.0")
	log.Debugf("debug enabled")
	log.Infof("starting server")
	if err := validateNewUserFlags(); err != nil {
		log.Error(err)
		return
	}
	if *newUser {
		err := database.NewUser(*username, *password, *role)
		if err != nil {
			log.Error(err)
			return
		}
		log.Infof("New user created successfully")
		return
	}
	log.Error(api.App.Listen(":8080"))
}
