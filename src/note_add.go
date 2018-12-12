package note

import (
	"fmt"
	"os"
	"path"
)

func MakeNote(name string) {
	folder, file := path.Split(name)
	num, found_note := FindNote(file)
	if num == -1 {
		if folder != "" {
			joinDir := DataBase() + "/" + folder
			os.MkdirAll(joinDir, os.ModePerm)
			newfile := joinDir + file + ".md"
			note, err := os.Create(newfile)
			ErrorCheck(err)
			defer note.Close()
			fmt.Printf("Note %q has been created\n", newfile)
		} else {
			newfile := DataBase() + "/" + file + ".md"
			note, err := os.Create(newfile)
			ErrorCheck(err)
			defer note.Close()
			fmt.Printf("Note %q has been created\n", newfile)
		}
	} else {
		fmt.Println("The note name name already exists:")
		fmt.Println(num, found_note)
	}
}
