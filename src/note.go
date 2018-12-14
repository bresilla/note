package note

import (
	"log"
	"os"
	"path"

	"github.com/urfave/cli"
)

func ErrorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func DataBase() string {
	var configDir = path.Join(homeDir(), ".notes")
	return configDir
}

func Run() {
	DeleteEmptyDir()
	var (
		listAll string
	)
	app := cli.NewApp()
	app.Name = "Note"
	app.Usage = "A minimal notetaking app from your CLI"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "colors",
		},
		cli.StringFlag{
			Name: "path",
		},
		cli.StringFlag{
			Name: "clean",
		},
	}
	app.Action = func(c *cli.Context) error {
		PrintList(ListNotePaths())
		return nil
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
					cli.ShowCommandHelp(c, "add")
				}
				return nil
			},
		},
		{
			Name:  "edit",
			Usage: "Edit existing note from database",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				if name != "" {
					EditNote(name)
				} else {
					cli.ShowCommandHelp(c, "edit")
				}
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a note from the database",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				if name != "" {
					DeleteNote(name)
				} else {
					cli.ShowCommandHelp(c, "delete")
				}
				return nil
			},
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
					PrintList(ListNoteNames())
				} else if listAll == "all" {
					Print(HighLight, Red, None, DashBorder())
					Print(HighLight, Red, None, ShowBanner())
					PrintList(ListNoteNames())
					Print(HighLight, Red, None, DashBorder())
				} else if listAll == "dir" {
					PrintList(ListNotePaths())
				} else if listAll == "group" {
					PrintList(ListNoteGroup())
				} else {
					cli.ShowCommandHelp(c, "list")
				}
				return nil
			},
		},
		{
			Name:  "test",
			Usage: "Snippets to test during development",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	DeleteEmptyDir()
	err := app.Run(os.Args)
	ErrorCheck(err)
}
