package main

import (
	"fmt"
)

var n int
var counter int // counter for stack
func main() {
	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&n)

}

type student struct {
	ID   int
	Name string
	GPA  int
	Pob  string // place of birth
}

// assignstudent for requirement 1 and 2
func requirement1() {
	var arr [100]students
	AssignStudent(&arr)
	arr[Maxgpa(arr)].Print()
	arr[Mingpa(arr)].Print()
}
func AssignStudent(s *[100]student) {
	for i := 0; i < n; i++ {
		fmt.Println("Please enter the Name of the student: ")
		fmt.Scanln(&s[i].Name)
		fmt.Println("Please enter the ID of the student: ")
		fmt.Scanln(&s[i].ID)
		for {
			fmt.Println("Please enter the GPA of the student: ")
			fmt.Scanln(&s[i].GPA)
			if s[i].GPA >= 0 && s[i].GPA <= 4 {
				break
			}
			fmt.Println("Error please enter a valid GPA")
		}
		fmt.Println("Please enter where the student was born: ")
		fmt.Scanln(&s[i].Pob)
	}

}
func (s students) Print() {
	fmt.Println(s.name)
	fmt.Println(s.GPA)
	fmt.Println(s.id)
	fmt.Println(s.pob)
}

// gpa calculation
func Maxgpa(s [100]students) int {
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

func Mingpa(s [100]students) int {
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

//assignstudent,  push and pop for req 3 and 4
func AssignStudentstack(s *[5]students) {
	fmt.Println("Please enter the information for 5 students")
	for i := 0; i < 5; i++ {
		s[i].Push()
	}
}

func (s *students) Push() {
	fmt.Println("Please enter the Name of the student: ")
	fmt.Scanln(&s.Name)
	fmt.Println("Please enter the ID of the student: ")
	fmt.Scanln(&s.ID)
	for {
		fmt.Println("Please enter the GPA of the student: ")
		fmt.Scanln(&s.GPA)
		if s.GPA >= 0 && s.GPA <= 4 {
			break
		}
		fmt.Println("Error please enter a valid GPA")
	}
	fmt.Println("Please enter where the student was born: ")
	fmt.Scanln(&s.Pob)
	counter++
}
func Pop(s *[5]students) students {
	counter--
	temp := s[counter]
	s[counter].Name = ""
	s[counter].ID = 0
	s[counter].GPA = 0
	s[counter].Pob = ""
	return temp
}

//
