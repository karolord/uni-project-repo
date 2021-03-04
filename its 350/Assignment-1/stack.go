package main

import (
	"fmt"
)

func main() {

	var s [5]student
	addID(s)
	fmt.Println(s[2].ID)

}

type student struct {
	ID int
}

func addID(s [5]*student) {
	for i := 0; i < len(s); i++ {
		s[i].ID = 2
	}
}
