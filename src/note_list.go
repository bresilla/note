package note

import (
	"fmt"
)

func ListNoteNames() []string {
	var noteList []string
	_, notes := ListPathsNFiles()
	for _, d := range notes {
		noteList = append(noteList, d.Name)
	}
	return noteList
}

func ListNoteGroup() []string {
	var noteList []string
	_, notes := ListPathsNFiles()
	for _, d := range notes {
		noteList = append(noteList, d.Parent)
	}
	return noteList
}

func ListNotePaths() []string {
	var noteList []string
	_, notes := ListPathsNFiles()
	for _, d := range notes {
		noteList = append(noteList, d.Path)
	}
	return noteList
}

func ListPathsOnly() []string {
	var dirList []string
	paths, _ := ListPathsNFiles()
	for _, d := range paths {
		dirList = append(dirList, d.Path)
	}
	return dirList
}

func PrintList(notes []string) {
	for _, n := range notes {
		fmt.Println(n)
	}
}
