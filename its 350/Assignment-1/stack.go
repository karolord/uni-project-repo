package main

import (
	"fmt"
)

var stack [5]int
var counter int

func main() {
	push(5)
	fmt.Println(stack[0])
	fmt.Println(counter)

}

func push(x int) {
	stack[counter] = x
	counter++
}

func pop() (int, error) {
	if counter == 0 {
		return 0, errors.new("stack is empty cannot pop")
	} else {
		var x = stack[counter-1]
		stack[counter-1] = 0
		counter--
		return x, nil
	}
}
