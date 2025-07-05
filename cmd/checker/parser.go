package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseArguments(args []string) ([]int, error) {
	if len(args) == 0 {
		return []int{}, nil
	}

	var numbers []int
	seen := make(map[int]bool)

	for _, arg := range args {
		// Split by spaces to handle quoted arguments like "2 1 3"
		parts := strings.Fields(arg)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid integer: %s", part)
			}
			
			if seen[num] {
				return nil, fmt.Errorf("duplicate number: %d", num)
			}
			seen[num] = true
			numbers = append(numbers, num)
		}
	}

	return numbers, nil
}

func createStackFromNumbers(numbers []int) *Stack {
	stack := NewStack()
	// Numbers are added in reverse order because stack is LIFO
	// First argument should be at the top
	for i := len(numbers) - 1; i >= 0; i-- {
		stack.Push(numbers[i])
	}
	return stack
}

func isAlreadySorted(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return false
		}
	}
	return true
}
