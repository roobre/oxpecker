package main

import (
	"fmt"
	"os"
	"roob.re/oxpecker"
)

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s <path/to/config>", os.Args[0])
		os.Exit(1)
	}

	ox, err := oxpecker.NewFromTomlFile(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	_ = ox.Run()
}
