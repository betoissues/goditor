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

func processKeyPress(reader *KeyReader) {
	char, _, err := reader.ReadKey()

	if char == 'q' {
		fmt.Println("closing")
		exitTerm(nil)
	}

	if unicode.IsControl(char) {
		fmt.Printf("%d\r\n", char)
	} else {
		fmt.Printf("%d (%c)\r\n", char, char)
	}

	if err != nil {
		exitTerm(err)
	}

	fmt.Print(string(char)+"\n")
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
