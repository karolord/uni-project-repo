package main

import (
	"fmt"
)

var numstudents int
var counter int // counter for struct
func main() {

	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&numstudents)
	fmt.Println(numstudents)
	var arr [100]Students
	assignstudents(&arr)
	fmt.Println(arr[5].GPA)

}

type Students struct {
	id   int
	name string
	GPA  float64
	pob  string
}

func assignstudents(s *[100]Students) {
	for i := 0; i < len(s); i++ {
		s[i].GPA = 5
	}
}
