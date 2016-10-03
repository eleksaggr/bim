package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zillolo/bim/builtin"
)

func initRepo() {
	fmt.Printf("Called init.\n")
	if len(flag.Args()) != 2 {
		flag.Usage()
		os.Exit(1)
	}
	name := flag.Args()[1]
	fmt.Printf("Creating project directory \"%v\" ...\n", name)
	builtin.Init(name)
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	command := flag.Args()[0]
	switch command {
	case "init":
		initRepo()
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command: %v\n", command)
	}
}
