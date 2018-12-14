package note

import (
	"fmt"
	"io"
	"os"
)

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func DeleteEmptyDir() {
	dirList := ListPathsOnly()
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
		fmt.Printf("Note %q deleted.\n", note)
		os.Remove(note)
	} else {
		fmt.Println("That note does not exist!")
	}
	DeleteEmptyDir()
}
