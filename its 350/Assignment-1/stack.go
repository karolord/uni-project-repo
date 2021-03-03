package main

import (
	"errors"
)

var stack [5]int
var counter int

func main() {

}

func push(x int) {
	stack[counter] = x
	counter++
}

func pop() (int, error) {
	if counter == 0 {
		return 0, errors.New("stack is empty cannot pop")
	} else {
		var x = stack[counter-1]
		stack[counter-1] = 0
		counter--
		return x, nil
	}
}
func maxGPAs() {
	var maxGPA int
	for i := 5; i > 0; i-- {
	}

}
