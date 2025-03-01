package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/getbread/goose/lib/goose"
)

var createCmd = &Command{
	Name:    "create",
	Usage:   "",
	Summary: "Create the scaffolding for a new migration",
	Help:    `create extended help here...`,
	Run:     createRun,
}

func createRun(cmd *Command, args ...string) {

	if len(args) < 1 {
		log.Fatal("goose create: migration name required")
	}

	migrationType := "go" // default to Go migrations
	if len(args) >= 2 {
		migrationType = args[1]
	}

	conf, err := dbConfFromFlags()
	if err != nil {
		log.Fatal(err)
	}

	if err = os.MkdirAll(conf.MigrationsDir, 0777); err != nil {
		log.Fatal(err)
	}

	n, err := goose.CreateMigration(args[0], migrationType, conf.MigrationsDir, time.Now())
	if err != nil {
		println("#1")
		log.Fatal(err)
	}

	a, e := filepath.Abs(n)
	if e != nil {
		println("#2")
		log.Fatal(e)
	}

	fmt.Println("goose: created", a)
}
