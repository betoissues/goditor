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

	sBuilder.WriteString(fmt.Sprintf("\x1b[%d;%dH", E.cy+1, E.cx+1))
	sBuilder.WriteString("\x1b[?25h")

	fmt.Fprint(os.Stdout, sBuilder.String())
	sBuilder.Reset()
}

func editorDrawRows() {
	for y := 0; y < E.termRows; y++ {
		if y == E.termRows/3 {
			welcome := "goditor -- version " + GODITOR_VERSION
			if len(welcome) > E.termCols {
				welcome = welcome[:E.termCols]
			}

			padding := (E.termCols - len(welcome)) / 2
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
		if y < E.termRows-1 {
			sBuilder.WriteString("\r\n")
		}
	}
}
