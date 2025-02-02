package main

import "fmt"

func main() {
	fmt.Println("calculator")
	fmt.Println("--------------------")

	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("number 1:", num1, "\nnumber 2:", num2)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

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
			fmt.Println("Cannot divide by zero")
			return
		}else{
		result = num1 / num2
		}
	default:
		fmt.Println("Invalid operator")
		return

	}


fmt.Println(num1, operator, num2, "=",result);

}