package main

import (
	"fmt"
)

var numstudents int
var counter int // counter for stack
func main() {
	var stack [5]Students
	AssignStudentstack(&stack)
	fmt.Println(Maxgpastack(&stack))
	fmt.Println(Mingpastack(&stack))

}

type Students struct {
	id   int
	name string
	GPA  float64
	pob  string
}

func AssignStudentstack(s *[5]Students) {
	for i := 0; i < 5; i++ {
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
func Pop(s *[5]Students) Students {
	counter--
	temp := s[counter]
	s[counter].name = ""
	s[counter].id = 0
	s[counter].GPA = 0
	s[counter].pob = ""
	return temp
}
func Maxgpastack(s *[5]Students) Students {
	var temp [5]Students
	for i := 0; i < 5; i++ {
		temp[i] = s[i]
	}
	var result Students
	tempcount := counter
	for i := tempcount - 1; i > 0; i-- {
		compare := Pop(&temp)
		if compare.GPA > result.GPA {
			result = compare
		}
	}
	counter = tempcount
	return result
}
func Mingpastack(s *[5]Students) Students {
	var temp [5]Students
	for i := 0; i < 5; i++ {
		temp[i] = s[i]
	}
	var result Students
	tempcount := counter
	for i := tempcount; i > 0; i-- {
		compare := Pop(&temp)
		if compare.GPA < result.GPA {
			result = compare
		}
	}
	counter = tempcount
	return result
}
