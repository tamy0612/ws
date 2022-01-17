package main

import (
	"log"
    "os"

	"github.com/urfave/cli/v2"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/tamy0612/ws/command"
)


func main() {
	db, _ := sqlx.Connect("sqlite3", "./dict/ejdict.sqlite3")

	app := &cli.App{
		Name: "ws",
        Usage:  "search words with conditions",
        Flags:  command.SearchFlags(),
        Action: command.Search(db),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
