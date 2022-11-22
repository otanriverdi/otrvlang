package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/otanriverdi/otrvlang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("otrvlang (%s)\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
