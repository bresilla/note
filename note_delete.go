package main

import (
	"fmt"
	"path"
)

func DeleteNote(name string) {
	file, folder := path.Split(name)
	fmt.Println(file, folder)
}
