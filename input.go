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

func InputInt(prompt string) int {
	for {
		input, err := ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}
		return number
	}
}

func InputString(prompt string) string {
	for {
		input, err := ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		return input
	}
}

func InputFloat(prompt string) float64 {
	for {
		input, err := ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}
		return number
	}
}

func InputBool(prompt string) bool {
	for {
		input, err := ReadInput(prompt)
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		lowerInput := strings.ToLower(input)
		if lowerInput == "y" || lowerInput == "true" {
			return true
		} else if lowerInput == "n" || lowerInput == "false" {
			return false
		} else {
			fmt.Println("Invalid input, please enter y or n.")
		}
	}
}

func main() {
	// Example usage
	fmt.Println("InputInt:", InputInt("Enter an integer: "))
	fmt.Println("InputString:", InputString("Enter a string: "))
	fmt.Println("InputFloat:", InputFloat("Enter a float: "))
	fmt.Println("InputBool:", InputBool("Enter a boolean (y/n): "))
}
