package main

import (
	"fmt"
	"github.com/AlexiaVeronica/input"
	"reflect"
)

func main() {
	// Example usage
	num := input.IntInput("Enter an integer: ")
	fmt.Println("You entered:", num)
	// 判断类型
	fmt.Println(reflect.TypeOf(num))

	floatNum := input.FloatInput("Enter a float: ")
	fmt.Println("You entered:", floatNum)

	boolean := input.BoolInput("Enter a boolean (y/n): ")
	fmt.Println("You entered:", boolean)

	int64Num := input.Int64Input("Enter an int64: ")
	fmt.Println("You entered:", int64Num)

	uintNum := input.UintInput("Enter a uint: ")
	fmt.Println("You entered:", uintNum)
}
