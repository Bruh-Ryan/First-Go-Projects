package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter +, -, *, / for basic calculation")
	fmt.Println("Enter 'scientific' or 'sci' for scientific calculation")
	fmt.Print("Enter operator: ")
	fmt.Scanln(&operator)

	if operator == "scientific" || operator == "sci" {
		var sciOp string
		fmt.Println("Scientific operations: sqrt, sin, cos, log, pow")
		fmt.Print("Enter scientific operation: ")
		fmt.Scanln(&sciOp)

		switch sciOp {
		case "sqrt":
			fmt.Print("Enter value: ")
			fmt.Scanln(&num1)
			fmt.Printf("sqrt(%.2f) = %.4f\n", num1, math.Sqrt(num1))
		case "sin":
			fmt.Print("Enter angle in degrees: ")
			fmt.Scanln(&num1)
			fmt.Printf("sin(%.2f) = %.4f\n", num1, math.Sin(num1*math.Pi/180))
		case "cos":
			fmt.Print("Enter angle in degrees: ")
			fmt.Scanln(&num1)
			fmt.Printf("cos(%.2f) = %.4f\n", num1, math.Cos(num1*math.Pi/180))
		case "log":
			fmt.Print("Enter value: ")
			fmt.Scanln(&num1)
			fmt.Printf("ln(%.2f) = %.4f\n", num1, math.Log(num1))
		case "pow":
			fmt.Print("Enter base: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter exponent: ")
			fmt.Scanln(&num2)
			fmt.Printf("%.2f ^ %.2f = %.4f\n", num1, num2, math.Pow(num1, num2))
		default:
			fmt.Println("Invalid scientific operation.")
		}
		return
	}

	// Basic operations
	fmt.Print("Enter number 1: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter number 2: ")
	fmt.Scanln(&num2)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator.")
		return
	}

	fmt.Println("Result:", result)
}
