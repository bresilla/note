package main

import (
	"fmt"
	"os"
)

func DeleteEmptyDir() {
	dirList := ListDirOnly()
	for _, d := range dirList {
		isEmpty, _ := IsDirEmpty(d)
		if isEmpty && d != DataBase() {
			os.RemoveAll(d)
		}
	}
}

func DeleteNote(name string) {
	_, note := FindNote(name)
	if note != "" {
		fmt.Printf("Note %q deleted\n", note)
		os.Remove(note)
	}
	DeleteEmptyDir()
}
