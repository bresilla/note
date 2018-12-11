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

func main() {
	var (
		listAll string
	)

	fmt.Println("---------------------------------------")
	app := cli.NewApp()
	app.Name = "Note"
	app.Usage = "A minimal notetaking app from your CLI"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		PrintList(ListNotes_Name())
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
					MakeNote(name)
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
					Name:        "format, f",
					Usage:       ": all, dir, group",
					Destination: &listAll,
				},
			},
			Action: func(c *cli.Context) error {
				if listAll == "" {
					PrintList(ListNotes_Name())
				} else if listAll == "all" {
					fmt.Println("TODO")
				} else if listAll == "dir" {
					PrintList(ListNotes_Dir())
				} else if listAll == "group" {
					PrintList(ListNotes_Group())
				} else {
					cli.ShowCommandHelp(c, "list")
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	ErrorCheck(err)
	fmt.Println("---------------------------------------")
}
