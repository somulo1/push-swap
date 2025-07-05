package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	a := createStackFromNumbers(numbers)
	b := NewStack()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		err := executeOperation(line, a, b)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	// Check if stack a is sorted and stack b is empty
	if a.IsSorted() && b.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
