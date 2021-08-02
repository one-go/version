package main

import (
	"flag"
	"fmt"
	"os"

	"git.woa.com/rio/version"
)

func main() {
	ver := flag.Bool("version", false, "print version")
	flag.Parse()

	if *ver {
		fmt.Print(version.String())
		os.Exit(0)
	}
}
