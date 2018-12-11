package main

import (
	"fmt"
	"os"
	"path"
)

func DataBase() string {
	var configDir = path.Join(os.Getenv("HOME"), ".notes")
	return configDir
}

func ListNotes() []string {
	files, err := ListFiles(DataBase())
	ErrorCheck(err)

	var notes []string

	for _, f := range files {
		if f.IsDir == true {
			continue
		}
		fmt.Println(f.Path)
		notes = append(notes, f.Path)
	}
	return notes
}
