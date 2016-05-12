package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/getbread/goose/lib/goose"
)

var downCmd = &Command{
	Name:    "down",
	Usage:   "",
	Summary: "Roll back the version by 1",
	Help:    `down extended help here...`,
	Run:     downRun,
}

func downRun(cmd *Command, args ...string) {

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	current, err := goose.GetDBVersion(conf)
	if err != nil {
		log.Fatal(err)
	}

	var previous int64

	if len(args) > 1 {
		previous, err = goose.GetEarliestSharedDBVersion(args[0], args[1])
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to determine the lowest migration number between %s and %s due to error: %s", args[0], args[1], err.Error()))
		}
		conf.MigrationsDir = args[0]

	} else if len(args) > 0 {
		previous, err = strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to parse %s as a migration version number due to error: %s", args[0], err.Error()))
		}

	} else {
		previous, err = goose.GetPreviousDBVersion(conf.MigrationsDir, current)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Migrating down to %d\n", previous)

	if err = goose.RunMigrations(conf, conf.MigrationsDir, previous); err != nil {
		log.Fatal(err)
	}
}
