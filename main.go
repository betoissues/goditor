package main

const (
	GODITOR_VERSION = "0.0.1"
)

func main() {
	makeRaw()
	initTerm()

	defer exitTerm(nil)

	for {
		editorRefreshScreen()
		processKeyPress()
	}

}
