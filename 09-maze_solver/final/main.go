package main

import (
	"fmt"
	"log/slog"
	"os"

	"tiny-go-projects/09-maze_solver/final/internal/solver"
)

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	inputFile := os.Args[0]
	outputFile := os.Args[1]

	slog.Info(fmt.Sprintf("Solving maze %s and saving it as %s", inputFile, outputFile))
	s, err := solver.New(inputFile)
	if err != nil {
		panic(err)
	}

	err = s.Solve()
	if err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	err = s.SaveSolution(outputFile)
	if err != nil {
		panic(err)
	}
}

func usage() {
	_, _ = fmt.Fprintln(os.Stderr, "Usage: maze_solver path/to/input.png path/to/output.png")
	os.Exit(1)
}
