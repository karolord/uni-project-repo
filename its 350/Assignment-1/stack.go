package main

import "fmt"

func main() {
	var n int32
	fmt.Scanf(&n)
	var arr [n]student
}

type student struct {
	ID int
}

func add(s *student) {
	s.ID = 5
}
