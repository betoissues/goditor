package main

import (
	"bufio"
	"fmt"
	"unicode"
	"os"
)

var keyReader = bufio.NewReader(os.Stdin)

func ctrlKey(key rune) rune {
	return key & 0x1f
}

func processKeyPress() {
	key, _, err := keyReader.ReadRune()

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
