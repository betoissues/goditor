package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var keyReader = bufio.NewReader(os.Stdin)

func ctrlKey(key rune) rune {
	return key & 0x1f
}

func processKeyPress() {
	key, _, err := keyReader.ReadRune()

	if err != nil {
		exitTerm(err)
	}

	switch key {
	case ctrlKey('q'):
		fmt.Println("closing")
		exitTerm(nil)
	default:
		if unicode.IsControl(key) {
			fmt.Printf("%d\r\n", key)
		} else {
			fmt.Printf("%d (%c)\r\n", key, key)
		}
	}

}
