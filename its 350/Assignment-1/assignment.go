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
	AssignStudent(&arr)
	arr[Maxgpa(arr)].Print()
	arr[Mingpa(arr)].Print()
	// end of first 2 requirements

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
