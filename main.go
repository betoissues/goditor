package main

import (
	"golang.org/x/term"
	"os"
)

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		exitTerm(err)
	}

	globalState.oldState = oldState
	globalState.restoreTerm = restoreState

	defer exitTerm(nil)

	for {
		editorRefreshScreen()
		processKeyPress()
	}

}
