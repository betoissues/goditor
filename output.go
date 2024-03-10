package main

import (
	"fmt"
	"os"
)

func editorRefreshScreen() {
	fmt.Fprint(os.Stdout, "\x1b[2J")
	fmt.Fprint(os.Stdout, "\x1b[H")

	editorDrawRows()

	fmt.Fprint(os.Stdout, "\x1b[H")
}

func editorDrawRows() {
	for y := 0; y < 24; y++ {
		fmt.Fprint(os.Stdout, "~\r\n")
	}
}
