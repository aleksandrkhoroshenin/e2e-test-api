package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

const UsageExitCode = 2

func main() {
	flag.Usage = usage

	flag.Parse()

	dbConnectionUrl, exists := os.LookupEnv("DB_CONNECTION_URL")
	if !exists {
		log.Fatal("DB_CONNECTION_URL env variable does not exist")
	}
	options, err := pg.ParseURL(dbConnectionUrl)
	failOnError(err, "Could not load DB configuration")

	dbConnection := pg.Connect(options)
	defer dbConnection.Close()

	oldVersion, newVersion, err := migrations.Run(dbConnection, flag.Args()...)

	if err != nil {
		exitf(err.Error())
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(UsageExitCode)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
