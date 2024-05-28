package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadInput(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func parseInput[T any](prompt string, parseFunc func(string) (T, error), errMsg string) T {
	for {
		input, err := ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		value, err := parseFunc(input)
		if err != nil {
			fmt.Println(errMsg)
			continue
		}
		return value
	}
}

func IntInput(prompt string) int {
	return parseInput(prompt, strconv.Atoi, "Invalid input, please enter a valid number.")
}

func StringInput(prompt string) string {
	return parseInput(prompt, func(input string) (string, error) { return input, nil }, "Error reading input, please try again.")
}

func FloatInput(prompt string) float64 {
	return parseInput(prompt, func(input string) (float64, error) { return strconv.ParseFloat(input, 64) }, "Invalid input, please enter a valid number.")
}

func BoolInput(prompt string) bool {
	return parseInput(prompt, func(input string) (bool, error) {
		lowerInput := strings.ToLower(input)
		if lowerInput == "y" || lowerInput == "true" {
			return true, nil
		} else if lowerInput == "n" || lowerInput == "false" {
			return false, nil
		}
		return false, fmt.Errorf("invalid input")
	}, "Invalid input, please enter y or n.")
}

func Int64Input(prompt string) int64 {
	return parseInput(prompt, func(input string) (int64, error) {
		return strconv.ParseInt(input, 10, 64)
	}, "Invalid input, please enter a valid number.")
}

func UintInput(prompt string) uint {
	return parseInput(prompt, func(input string) (uint, error) {
		value, err := strconv.ParseUint(input, 10, 64)
		return uint(value), err
	}, "Invalid input, please enter a valid number.")
}
