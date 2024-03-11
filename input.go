package main

import (
	"bufio"
	"fmt"
	"os"
)

var keyReader = bufio.NewReader(os.Stdin)

const (
	keyLeft  = 1000
	keyRight = 1001
	keyUp    = 1002
	keyDown  = 1003
)

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
				return keyUp
			case 'B':
				return keyDown
			case 'C':
				return keyRight
			case 'D':
				return keyLeft
			}
		}

		return key
	} else {
		return key
	}
}

func editorMoveCursor(key rune) {
	switch key {
	case keyUp:
		E.cy--
	case keyLeft:
		E.cx--
	case keyDown:
		E.cy++
	case keyRight:
		E.cx++
	}
}

func processKeyPress() {
	key := editorReadKey()

	switch key {
	case ctrlKey('q'):
		fmt.Println("closing")
		exitTerm(nil)
	case keyUp, keyLeft, keyDown, keyRight:
		editorMoveCursor(key)
	}
}
