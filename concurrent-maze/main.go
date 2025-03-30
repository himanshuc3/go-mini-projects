package main

import (
	"fmt"
	"maze-solver/internal/solver"
	"os"
)

// NOTE:
// 1. go doc image

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFilePath := os.Args[1]
	// outputFilePath := os.Args[2]

	// NOTE:
	// 1. log vs fmt
	// log is thread-safe where fmt is not
	// 2. %q - for printing strings in double quotes
	mazeSolver, err := solver.New(inputFilePath)

	fmt.Printf("The solver object: %T %v", mazeSolver, mazeSolver)
}

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver input.png output.png")
	os.Exit(1)
}
