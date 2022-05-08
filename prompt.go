package main

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

func main() {
	fmt.Print("Password" + ":" + " ")

	input, exception := term.ReadPassword(int(os.Stdin.Fd()))

	if exception != nil {
		os.Exit(1)
	}

	fmt.Print("\n", "User-Input", ":", " ", string(input), "\n")
}
