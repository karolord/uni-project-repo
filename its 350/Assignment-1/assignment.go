package main

import (
	"fmt"
)

var numstudents int

func main() {

	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&numstudents)
	fmt.Println(numstudents)
	var s Students
	s.AssignStudent(numstudents)
	fmt.Println(s.maxgpa())
}

type Students struct {
	id   [100]int
	name [100]string
	GPA  [100]float64
	pob  [100]string
}

func (s *Students) AssignStudent(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Please enter the Name of the student: ")
		fmt.Scanln(&s.name[i])
		fmt.Println("Please enter the ID of the student: ")
		fmt.Scanln(&s.id[i])
		fmt.Println("Please enter the GPA of the student: ")
		fmt.Scanln(&s.GPA[i])
		fmt.Println("Please enter where the student was born: ")
		fmt.Scanln(&s.pob[i])
	}
}
func (s Students) Print() {
	for i := 0; i < numstudents; i++ {
		fmt.Println(s.name[i])
		fmt.Println(s.GPA[i])
		fmt.Println(s.id[i])
		fmt.Println(s.pob[i])
	}
}
func (s Students) maxgpa() int {
	var maxGPA float64
	var index int
	for k := 0; k < numstudents; k++ {
		if s.GPA[k] > maxGPA {
			maxGPA = s.GPA[k]
			index = k
		}
	}
	return index
}
func (s Students) mingpa() int {
	var minGPA float64
	var index int
	for k := 0; k < numstudents; k++ {
		if s.GPA[k] < minGPA {
			maxGPA = s.GPA[k]
			index = k
		}
	}
	return index
}
