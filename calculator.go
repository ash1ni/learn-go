package main

import "fmt"

func main(){
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Enter operator (+,-,*,/); ")
	fmt.Scanln(&operator)

	var result float64

	switch operator{
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1/num2
	default: 
		fmt.Println("Invalid operator")
		return 

	}
	fmt.Printf("%.2f %s %.2f = %.2f", num1, operator, num2, result)
}