package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/one-go/version"
)

func main() {
	ver := flag.Bool("version", false, "print version")
	flag.Parse()

	if *ver {
		fmt.Print(version.String())
		os.Exit(0)
	}
}
