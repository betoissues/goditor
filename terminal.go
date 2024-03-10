package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

var globalState = struct {
	restoreTerm func()
	oldState *term.State
}{
	nil,
	nil,
}

// @TODO: review best way to implement the reader without
// this mostly useless wrapper

// type KeyReader struct {
// 	*bufio.Reader
// }

// func NewKeyReader() *KeyReader {
// 	return &KeyReader{
// 		bufio.NewReader(os.Stdin),
// 	}
// }
// func (kr *KeyReader) ReadKey() (rune, int, error) {
// 	return kr.ReadRune()
// }

func restoreState() {
	term.Restore(int(os.Stdin.Fd()), globalState.oldState)
}

func exitTerm(err error) {
	if globalState.restoreTerm != nil {
		globalState.restoreTerm()
	}

	if err != nil {
		fmt.Println("error: "+err.Error())
		globalState.restoreTerm()
		os.Exit(1)
	}

	os.Exit(0)
}
