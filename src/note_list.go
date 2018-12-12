package note

import (
	"fmt"
	"io"
	"os"
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

func ListDirOnly() []string {
	files, err := ListFiles(DataBase())
	ErrorCheck(err)
	var notePath []string
	for _, f := range files {
		if f.IsDir == false {
			continue
		}
		notePath = append(notePath, f.Path)
	}
	return notePath
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
