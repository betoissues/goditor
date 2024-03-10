package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"unicode"
)

type KeyReader struct {
	*bufio.Reader
}

var globalState = struct {
	restoreTerm func()
	oldState *term.State
}{
	nil,
	nil,
}

func NewKeyReader() *KeyReader {
	return &KeyReader{
		bufio.NewReader(os.Stdin),
	}
}
func (kr *KeyReader) ReadKey() (rune, int, error) {
	return kr.ReadRune()
}

func ctrlKey(key rune) rune {
	return key & 0x1f
}

func processKeyPress(reader *KeyReader) {
	key, _, err := reader.ReadKey()

	if unicode.IsControl(key) {
		fmt.Printf("%d\r\n", key)
	} else {
		fmt.Printf("%d (%c)\r\n", key, key)
	}

	if key == ctrlKey('q') {
		fmt.Println("closing")
		exitTerm(nil)
	}

	if err != nil {
		exitTerm(err)
	}
}

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

func main() {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		exitTerm(err)
	}
	
	globalState.oldState = oldState
	globalState.restoreTerm = restoreState

	defer exitTerm(nil)

	reader := NewKeyReader()
	for {
		processKeyPress(reader)
	}

}
