package main

import (
	"bufio"
	"fmt"
	"os"
)

var keyReader = bufio.NewReader(os.Stdin)

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
				return 'w'
			case 'B':
				return 's'
			case 'C':
				return 'd'
			case 'D':
				return 'a'
			}
		}

		return key
	} else {
		return key
	}
}

func editorMoveCursor(key rune) {
	switch key {
	case 'w':
		E.cy--
	case 'a':
		E.cx--
	case 's':
		E.cy++
	case 'd':
		E.cx++
	}
}

func processKeyPress() {
	key := editorReadKey()

	switch key {
	case ctrlKey('q'):
		fmt.Println("closing")
		exitTerm(nil)
	case 'w', 'a', 's', 'd':
		editorMoveCursor(key)
	}
}
