package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

type EditorConfig struct {
	cx          int
	cy          int
	termCols    int
	termRows    int
	restoreTerm func()
	oldState    *term.State
}

var E = EditorConfig{}

var escSeq = '\x1b'

func restoreState() {
	term.Restore(int(os.Stdin.Fd()), E.oldState)
}

func makeRaw() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		exitTerm(err)
	}

	E.oldState = oldState
	E.restoreTerm = restoreState
}

func initTerm() {
	cols, rows, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		exitTerm(err)
	}

	E.termCols = cols
	E.termRows = rows
}

func exitTerm(err error) {
	E.cx = 0
	E.cy = 0
	editorRefreshScreen()

	// allows using `exitTerm` if `makeRaw` fails
	if E.restoreTerm != nil {
		E.restoreTerm()
	}

	if err != nil {
		fmt.Println("error: " + err.Error())
		E.restoreTerm()
		os.Exit(1)
	}

	os.Exit(0)
}
