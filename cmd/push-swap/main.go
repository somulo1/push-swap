package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	numbers, err := parseArguments(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if len(numbers) == 0 {
		return
	}

	instructions := pushSwapAlgorithm(numbers)
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
