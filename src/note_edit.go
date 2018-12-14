package note

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const defaultEditor = "nvim"

func OpenInEditor(filepath string) bool {
	var cmd *exec.Cmd

	editor := os.Getenv("EDITOR")

	if len(editor) > 0 {
		cmd = exec.Command(editor, filepath)
	} else {
		if _, err := os.Stat("/usr/bin/sensible-editor"); err == nil {
			cmd = exec.Command("/usr/bin/sensible-editor", filepath)
		} else {
			cmd = exec.Command("/usr/bin/env", defaultEditor, filepath)
		}
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println("Error:", err)
		log.Println("File not saved:", filepath)
		return false
	}
	return true
}

func EditNote(name string) {
	_, note := FindNote(name)
	if note != "" {
		OpenInEditor(note)
	} else {
		fmt.Println("That note does not exist!")
	}
	DeleteEmptyDir()
}
