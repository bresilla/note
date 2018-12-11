package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func ErrorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func EmptyRun(argument string) {
	MakeNote(argument)
}

func main() {
	fmt.Println("---------------------------------------")
	app := cli.NewApp()
	app.Name = "Note"
	app.Usage = "A minimal notetaking app from your CLI"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		ListNotes()
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "colors",
		},
		cli.StringFlag{
			Name: "path",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add a new note to the database",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				if name != "" {
					EmptyRun(name)
				} else {
					fmt.Println("USAGE: note add [name]")
				}
				return nil
			},
		},
		{
			Name:  "edit",
			Usage: "Edit an existing notes",
		},
		{
			Name:  "delete",
			Usage: "Delete a note from the database",
		},
		{
			Name:  "show",
			Usage: "Concatenate and show a note",
		},
		{
			Name:  "list",
			Usage: "List all notes from the database",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "list-all",
				},
			},
		},
	}

	err := app.Run(os.Args)
	ErrorCheck(err)
}
