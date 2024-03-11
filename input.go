package main

import (
	"bufio"
	"fmt"
	"os"
)

var keyReader = bufio.NewReader(os.Stdin)

var editorKeys = struct {
	ARROW_LEFT  rune
	ARROW_RIGHT rune
	ARROW_UP    rune
	ARROW_DOWN  rune
}{
	ARROW_LEFT:  1000,
	ARROW_RIGHT: 1001,
	ARROW_UP:    1002,
	ARROW_DOWN:  1003,
}

func ctrlKey(key rune) rune {
	return key & 0x1f
}

func editorReadKey() rune {
	key, _, err := keyReader.ReadRune()

	if err != nil {
		exitTerm(err)
	}

	if key == rune(escSeq) {
		var keySize int
		seq := make([]rune, 3)

		// @TODO: without timeout for `ReadRune()`
		// no assurance this is an escape sequence command
		seq[0], keySize, err = keyReader.ReadRune()

		if keySize < 1 {
			return key
		}

		seq[1], keySize, err = keyReader.ReadRune()

		if keySize < 1 {
			return key
		}

		if seq[0] == '[' {
			switch seq[1] {
			case 'A':
				return editorKeys.ARROW_UP
			case 'B':
				return editorKeys.ARROW_DOWN
			case 'C':
				return editorKeys.ARROW_RIGHT
			case 'D':
				return editorKeys.ARROW_LEFT
			}
		}

		return key
	} else {
		return key
	}
}

func editorMoveCursor(key rune) {
	switch key {
	case editorKeys.ARROW_UP:
		E.cy--
	case editorKeys.ARROW_LEFT:
		E.cx--
	case editorKeys.ARROW_DOWN:
		E.cy++
	case editorKeys.ARROW_RIGHT:
		E.cx++
	}
}

func processKeyPress() {
	key := editorReadKey()

	switch key {
	case ctrlKey('q'):
		fmt.Println("closing")
		exitTerm(nil)
	case editorKeys.ARROW_UP,
		editorKeys.ARROW_LEFT,
		editorKeys.ARROW_DOWN,
		editorKeys.ARROW_RIGHT:
		editorMoveCursor(key)
	}
}
