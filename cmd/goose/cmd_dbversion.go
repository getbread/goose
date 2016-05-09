package main

import (
	"fmt"
	"log"

	"github.com/getbread/goose/lib/goose"
)

var dbVersionCmd = &Command{
	Name:    "dbversion",
	Usage:   "",
	Summary: "Print the current version of the database",
	Help:    `dbversion extended help here...`,
	Run:     dbVersionRun,
}

func dbVersionRun(cmd *Command, args ...string) {
	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 2 {
		version, err := goose.GetEarliestSharedDBVersion(args[0], args[1])

		if err != nil {
			fmt.Printf("Failed to determine the earliest shared version due to error: %s", err.Error())
		} else {
			fmt.Printf("Earliest shared version between %s and %s is: %d\n", args[0], args[1], version)
		}

	} else {
		current, err := goose.GetDBVersion(conf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("goose: dbversion %v\n", current)
	}

}
