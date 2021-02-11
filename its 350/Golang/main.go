package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

// variables in go
// string
//int
// byte for char/ int8
//int16 short
//int32 normal int
//int64 long
//uint32 refers to unsigned meaning only positive values
//float32 float / float64 double
// in go int changes size based on value

func main() {
	fmt.Println("Hello world")
	// variables without declaration
	test := 5
	num2 := 8
	sum := test + num2
	fmt.Println(sum)
	// if condition
	if test != 5 {
		fmt.Println("Hello")
	}

	//concatination in println
	fmt.Println("yes")

	//another type of concatination assigning a string the value of sprintf
	s := fmt.Sprintf("%d is 5 and %d is 8\n", test, num2)

	io.WriteString(os.Stdout, s)

	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println("test")

	}
	i := 5
	// while in for form
	for i < 6 {
		fmt.Println(sum)
		i++
	}

	// Arrays declaring and intitalizing with values
	var a [5]int
	a[2] = 5
	//short syntax
	b := [5]int{5, 34, 3, 2, 1}
	fmt.Println(b)
	fmt.Println(a[2])
	// Dynamic array or slices
	c := []int{1, 2, 3}
	//appending an array modifies the size of array adnd adds the value 13
	c = append(c, 13)

	// iterating over an array
	data := []string{"yes", "no"}
	for index, value := range data {
		fmt.Println(index)
		fmt.Println(value)
	}

	//sqrt test
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

// function declaration public to all if first letter is capitalized
func Sum(x int, y int) int {
	return y + x
}

// function is public to this package if the first letter is capitalized
func sub(x int, y int) int {
	return x - y
}

// returning an error along with the value
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("go fk yourself")
	}

	return math.Sqrt(x), nil

}
