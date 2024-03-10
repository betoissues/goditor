package main

import (
	"fmt"
	"os"
)

func editorRefreshScreen() {
	fmt.Fprint(os.Stdout, "\x1b[2J")
}
