package main

import (
	"fmt"
)

func TestSnip() {
	width, _ := Tty_Size()
	for n := 1; n <= width; n++ {
		fmt.Print("-")
	}
	fmt.Println("")
}
