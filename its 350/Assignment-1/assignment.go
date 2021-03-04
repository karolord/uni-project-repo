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
	for i := 0; i < numstudents; i++ {
		arr[i].AssignStudent()
	}
	arr[Maxgpa(arr)].Print()
	arr[Mingpa(arr)].Print()
	// end of first 2 requirements
	var struct [5]Students

}

type Students struct {
	id   int
	name string
	GPA  float64
	pob  string
}

func (s *Students) AssignStudent() {
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

func (s *Student) push(){

}


func (s *Student) pop(){


}