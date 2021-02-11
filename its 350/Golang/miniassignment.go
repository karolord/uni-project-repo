package main

import (
	"fmt"
	"math"
)

func main() {
	var v1 Variable
	var v2 Variable
	var op Operation
	fmt.Println("Thanks for using this amazing calculator")
	fmt.Println("List of operations\n1.Addition \n2.Subtraction \n3.Multiplication\n4.Natural Log\n5.Exponent")
	fmt.Printf("Please select one of the following operations:")
	fmt.Scanln(&op.value)
	fmt.Printf("Please enter the first number:")
	fmt.Scanln(&v1.value)
	fmt.Printf("Please enter the second number:")
	fmt.Scanln(&v2.value)
	switch op.value {
	case 1:
		Addition(v1.value, v2.value)
	case 2:
		Subtraction(v1.value, v2.value)
	case 3:
		Multiplication(v1.value, v2.value)

	case 4:
		Nlog(v1.value, v2.value)

	default:
		Exponent(v1.value, v2.value)

	}

}

type Variable struct {
	value float64
}

type Operation struct {
	value int
}

func Addition(v1 float64, v2 float64) {
	result := v1 + v2
	fmt.Println("The result is :", result)
}

func Subtraction(v1 float64, v2 float64) {
	result := v1 - v2
	fmt.Println("The result is :", result)
}

func Multiplication(v1 float64, v2 float64) {
	result := v1 * v2
	fmt.Println("The result is :", result)
}
func Nlog(v1 float64, v2 float64) {
	nlog1 := math.Log(v1)
	nlog2 := math.Log(v2)
	fmt.Println("The natural log of", v1, "is", nlog1)
	fmt.Println("The natural log of", v2, "is", nlog2)

}
func Exponent(v1 float64, v2 float64) {
	result := math.Pow(v1, v2)
	fmt.Println("The result is :", result)
}
