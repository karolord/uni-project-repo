package main

import "fmt"

func main() {

	var x student
	add(x)

	
	type student struct {
}

	ID int
}

func add(s student) {
	s.ID = 5
}
