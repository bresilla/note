package main

import (
	"fmt"
	"path"
	"regexp"
	"strings"
)

func ListNotes_Dir() []string {
	files, err := ListFiles(DataBase())
	ErrorCheck(err)
	var notePath []string
	for _, f := range files {
		if f.IsDir == true {
			continue
		}
		notePath = append(notePath, f.Path)
	}
	return notePath
}

var Note_Path = ListNotes_Dir()

func ListNotes_Name() []string {
	var note_names []string
	for _, n := range Note_Path {
		_, file := path.Split(n)
		s := strings.Split(file, ".")
		note_names = append(note_names, s[0])
	}
	return note_names
}

func ListNotes_Group() []string {
	var note_group []string
	for _, n := range Note_Path {
		folder, file := path.Split(n)
		database_path := DataBase() + "/"
		re := regexp.MustCompile(database_path)
		sub := re.ReplaceAllString(folder, "")
		group := sub + file
		note_group = append(note_group, group)
	}
	return note_group
}

func PrintList(notes []string) {
	for _, n := range notes {
		fmt.Println(n)
	}
}
