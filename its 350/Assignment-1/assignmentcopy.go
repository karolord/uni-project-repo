package main

import (
	"fmt"
)

var numstudents int

func main() {

	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&numstudents)
	fmt.Println(numstudents)
	var s [100]Students
}

type Students struct {
	id   int
	name string
	GPA  float64
	pob  string
}

func (s *Students) AssignStudent(n int) {
	fmt.Println("Please enter the Name of the student: ")
	fmt.Scanln(&s.name)
	fmt.Println("Please enter the ID of the student: ")
	fmt.Scanln(&s.id)
	fmt.Println("Please enter the GPA of the student: ")
	fmt.Scanln(&s.GPA)
	fmt.Println("Please enter where the student was born: ")
	fmt.Scanln(&s.pob)
}

func (s Students) Print() {
	fmt.Println(s.name)
	fmt.Println(s.GPA)
	fmt.Println(s.id)
	fmt.Println(s.pob)
}
func assignstudents(n int, a []Students) {
	for i := 0; i < n; i++ {
		a[i].AssignStudent()
	}
}
