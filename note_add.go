package main

import (
	"fmt"
	"os"
)

func MakeNote(name string) {
	num, exnote := FindNote(name)
	if num == -1 {
		var newfile string = DataBase() + "/" + name + ".md"
		file, err := os.Create(newfile)
		ErrorCheck(err)
		defer file.Close()
		fmt.Printf("Note %q has been created\n", newfile)
	} else {
		fmt.Println("The note name name already exists:")
		fmt.Println(num, exnote)
	}
}

func MakeNoteTest(name string) {
	num, exnote := FindNote(name)
	if num == -1 {
		var newfile string = DataBase() + "/" + name + ".md"
		file, err := os.Create(newfile)
		ErrorCheck(err)
		defer file.Close()
		fmt.Printf("Note %q has been created\n", newfile)
	} else {
		fmt.Println("The note name name already exists:")
		fmt.Println(num, exnote)
	}
}
