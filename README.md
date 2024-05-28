# User Input Utility

This project provides a set of utilities to facilitate the reading and parsing of user input from the command line. It supports various types of input including strings, integers, floats, booleans, and unsigned integers.

## Introduction

Command-line applications often require user inputs, which need to be read and validated efficiently. This project aims to simplify this process by offering a straightforward API to read and parse different types of inputs. The core motivation behind this project is to reduce boilerplate code in command-line applications, making them easier to maintain and extend.

Key features of this utility include:
- **Ease of Use**: Simple functions to read various input types.
- **Error Handling**: Robust error handling and user prompts for incorrect inputs.
- **Type Safety**: Functions for specific data types ensuring type safety.

## Key Features

- **IntInput**: Read and parse integer inputs.
- **StringInput**: Read string inputs.
- **FloatInput**: Read and parse floating-point numbers.
- **BoolInput**: Read boolean inputs with support for "y/n" and "true/false".
- **Int64Input**: Read and parse 64-bit integer inputs.
- **UintInput**: Read and parse unsigned integer inputs.

## Installation and Deployment

To use the User Input Utility in your project, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/user-input-utility.git
   cd user-input-utility
   ```

2. **Install Dependencies**:
   This project does not have any external dependencies outside the Go standard library.

3. **Build the Project**:
   ```bash
   go build
   ```

4. **Run the Project**:
   You can run the provided examples or integrate the utility functions into your own projects.

## Usage

To use the User Input Utility, import the package and call the desired functions. Below are some typical usage examples:

```go
package main

import (
	"fmt"
	"input"
)

func main() {
	name := input.StringInput("Enter your name: ")
	age := input.IntInput("Enter your age: ")
	height := input.FloatInput("Enter your height (in cm): ")
	isStudent := input.BoolInput("Are you a student? (y/n): ")

	fmt.Printf("Name: %s, Age: %d, Height: %.2f cm, Student: %t\n", name, age, height, isStudent)
}
```

To run the above example, save it in a file (e.g., `main.go`), and execute:
```bash
go run main.go
```

### Running Tests

If the project includes test cases, you can run them using the following command:
```bash
go test ./...
```
This will execute all test cases and provide a report on the coverage.

## Dependencies

The project relies on the Go standard library:
- `bufio`
- `fmt`
- `os`
- `strconv`
- `strings`

Ensure you have Go installed on your machine. The project is compatible with Go versions 1.16 and above.

## Development Status

The project is currently in the active development phase, with plans for future enhancements including support for additional data types and more sophisticated validation mechanisms. Contributions are welcome!

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
 

Thank you for using the User Input Utility! We hope it makes your command-line applications more robust and user-friendly.