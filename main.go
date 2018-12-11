package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	cli "gopkg.in/urfave/cli.v1"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func DataBase() string {
	var configDir = path.Join(os.Getenv("HOME"), ".notes")
	return configDir
}

func ListNotes() {
	files, err := ioutil.ReadDir(DataBase())
	check(err)

	for _, f := range files {
		fmt.Printf("%v/%v\n", DataBase(), f.Name())
	}
}

func MakeNote(name string) {
	var newfile string = DataBase() + "/" + name + ".md"

	file, err := os.Create(newfile)
	check(err)
	defer file.Close()

	fmt.Println(newfile)
}

func main() {
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
					MakeNote(name)
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
	check(err)
}
