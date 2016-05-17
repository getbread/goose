package main

import (
	"fmt"
	"log"

	"github.com/getbread/goose/lib/goose"
)

var upCmd = &Command{
	Name:    "up",
	Usage:   "",
	Summary: "Migrate the DB to the most recent version available",
	Help:    `up extended help here...`,
	Run:     upRun,
}

func upRun(cmd *Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	target, err := goose.GetMostRecentDBVersion(conf.MigrationsDir)
	if err != nil {
		log.Fatal(fmt.Errorf("Failed to read the most recent database version from %s due to error: %s", conf.MigrationsDir, err.Error()))
	}

	if err := goose.RunMigrations(conf, conf.MigrationsDir, target); err != nil {
		log.Fatal(fmt.Errorf("Failed to run the goose migrations on %v with directory %s and target %d due to error: %s", conf, conf.MigrationsDir, target, err.Error()))
	}
}
