package main

import (
	"fmt"
)

var n int
var counter int // counter for stack
func main() {
	numberassignment()
	Requirement1()
	requirement2()

}

type student struct {
	ID   int
	Name string
	GPA  int
	Pob  string // place of birth
}
func numberassignment(){
	for true {
	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&n)
	if n =< 100 {break}
	fmt.Println("Error please enter a value less then or equal than 100")
	}
}
// assignstudent for requirement 1 and 2
func requirement1() {
	var arr [100]students
	AssignStudent(&arr)
	Maxgpa(arr)
	Mingpa(arr)
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
func (s student) Print() {
	fmt.Println("The student name is ",s.Name)
	fmt.Println("The student GPA is", s.GPA)
	fmt.Println("The student ID is" s.ID)
	fmt.Println("The student Place of Birth is ",s.Pob)
}

// gpa calculation
func Maxgpa(s [100]student)  {
	x := s[0].GPA
	index := 0
	for i := 1; i < n; i++ {
		if s[i].GPA > x {
			x = s[i].GPA
			index = i
		}
	}
	fmt.Println("The student with the highest GPA is ")
	s[index].Print()
}

func Mingpa(s [100]student) {
	x := s[0].GPA
	index := 0
	for i := 1; i < n; i++ {
		if s[i].GPA < x {
			x = s[i].GPA
			index = i
		}
	}
	fmt.Println("The student with the lowest GPA is ")
	s[index].Print()}

//assignstudent,  push and pop for req 3 and 4
func Requirement2(){
	var stack [5]student
	AssignStudent(&stack)
	Maxgpastack(&stack)
	Mingpastack(&stack)
}
func AssignStudentstack(s *[5]student) {
	fmt.Println("Please enter the information for 5 students")
	for i := 0; i < 5; i++ {
		s[i].Push()
	}
}

func (s *student) Push() {
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
func Pop(s *[5]student) students {
	counter--
	temp := s[counter]
	s[counter].Name = ""
	s[counter].ID = 0
	s[counter].GPA = 0
	s[counter].Pob = ""
	return temp
}
func Maxgpastack(s *[5]student)  {
	var temp [5]student
	for i := 0; i < 5; i++ {
		temp[i] = s[i]
	}
	var result student
	tempcount := counter
	for i := tempcount - 1; i > 0; i-- {
		compare := Pop(&temp)
		if compare.GPA > result.GPA {
			result = compare
		}
	}
	counter = tempcount
	fmt.Println("The student with the highest GPA is ")
	result.Print()
}
func Mingpastack(s *[5]student)  {
	var temp [5]students
	for i := 0; i < 5; i++ {
		temp[i] = s[i]
	}
	var result students
	tempcount := counter
	for i := tempcount; i > 0; i-- {
		compare := Pop(&temp)
		if compare.GPA < result.GPA {
			result = compare
		}
	}
	counter = tempcount
	fmt.Println("The student with the lowest GPA is ")
	result.Print()
}

// Requirement 5 and 6
