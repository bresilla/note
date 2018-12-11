package main

import (
	"fmt"
	"os"
)

func MakeNote(name string) {
	var newfile string = DataBase() + "/" + name + ".md"

	file, err := os.Create(newfile)
	ErrorCheck(err)
	defer file.Close()

	fmt.Println(newfile)
}
