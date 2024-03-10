package main

import (
	"fmt"
	"os"
	"strings"
)

var sBuilder strings.Builder

func editorRefreshScreen() {
	sBuilder.WriteString("\x1b[?25l")
	sBuilder.WriteString("\x1b[H")

	editorDrawRows()

	sBuilder.WriteString("\x1b[H")
	sBuilder.WriteString("\x1b[?25h")
	fmt.Fprint(os.Stdout, sBuilder.String())
	sBuilder.Reset()
}

func editorDrawRows() {
	for y := 0; y < globalState.screenrows; y++ {
		if y == globalState.screenrows/3 {
			welcome := "goditor -- version 0.0.1"
			if len(welcome) > globalState.screencols {
				welcome = welcome[:globalState.screencols]
			}
			padding := (globalState.screencols - len(welcome)) / 2
			if padding > 0 {
				sBuilder.WriteString("~")
				padding--
			}

			for padding > 0 {
				sBuilder.WriteString(" ")
				padding--
			}
			sBuilder.WriteString(welcome)
		} else {
			sBuilder.WriteString("~")
		}

		sBuilder.WriteString("\x1b[K")
		if y < globalState.screenrows-1 {
			sBuilder.WriteString("\r\n")
		}
	}
}
