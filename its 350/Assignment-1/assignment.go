package main

import (
	"fmt"
)

var numstudents int
var counter int // counter for stack
func main() {
	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&numstudents)
	var arr [100]Students
	AssignStudent(&arr)
	arr[Maxgpa(arr)].Print()
	arr[Mingpa(arr)].Print()
	// end of first 2 requirements
	var stack [5]Students
	AssignStudentstack(&stack)
	Pop(&stack)
	fmt.Println(counter)
	fmt.Println(stack[0])

}

type Students struct {
	id   int
	name string
	GPA  float64
	pob  string
}

func AssignStudent(s *[100]Students) {
	for i := 0; i < numstudents; i++ {
		fmt.Println("Please enter the Name of the student: ")
		fmt.Scanln(&s[i].name)
		fmt.Println("Please enter the ID of the student: ")
		fmt.Scanln(&s[i].id)
		for {
			fmt.Println("Please enter the GPA of the student: ")
			fmt.Scanln(&s[i].GPA)
			if s[i].GPA >= 0 && s[i].GPA <= 4 {
				break
			}
			fmt.Println("Error please enter a valid GPA")
		}
		fmt.Println("Please enter where the student was born: ")
		fmt.Scanln(&s[i].pob)
	}
}
func (s Students) Print() {
	fmt.Println(s.name)
	fmt.Println(s.GPA)
	fmt.Println(s.id)
	fmt.Println(s.pob)

}
func Maxgpa(s [100]Students) int {
	x := s[0].GPA
	index := 0
	for i := 1; i < numstudents; i++ {
		if s[i].GPA > x {
			x = s[i].GPA
			index = i
		}
	}
	return index
}

func Mingpa(s [100]Students) int {
	x := s[0].GPA
	index := 0
	for i := 1; i < numstudents; i++ {
		if s[i].GPA < x {
			x = s[i].GPA
			index = i
		}
	}
	return index
}
func AssignStudentstack(s *[5]Students) {
	for i := 0; i < 1; i++ {
		s[i].Push()
	}
}
func (s *Students) Push() {
	fmt.Println("Please enter the Name of the student: ")
	fmt.Scanln(&s.name)
	fmt.Println("Please enter the ID of the student: ")
	fmt.Scanln(&s.id)
	for {
		fmt.Println("Please enter the GPA of the student: ")
		fmt.Scanln(&s.GPA)
		if s.GPA >= 0 && s.GPA <= 4 {
			break
		}
		fmt.Println("Error please enter a valid GPA")
	}
	fmt.Println("Please enter where the student was born: ")
	fmt.Scanln(&s.pob)
	counter++
}
func Pop(s *[5]Students) {
	fmt.Println("yes")
	s[counter-1].name = "karo"
	fmt.Println(s[counter].name)
	s[counter-1].id = 0
	s[counter-1].GPA = 0
	s[counter].pob = " "
	counter--

}
