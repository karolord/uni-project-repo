package main

import (
	"fmt"
)

var numstudents int
var counter int // counter for stack
func main() {
	var arr2 [100]students
	fmt.Println(arr2[2].GPA)
	fmt.Println("Please enter the number of students: ")
	fmt.Scanln(&numstudents)
	var arr [100]students
	AssignStudent(&arr)
	arr[Maxgpa(arr)].Print()
	arr[Mingpa(arr)].Print()
	// end of first 2 requirements
	var stack [5]students
	AssignStudentstack(&stack)
	fmt.Println(Pop(&stack))
	fmt.Println(counter)
	fmt.Println(stack[counter])
	fmt.Println(Maxgpastack(&stack))

}

type students struct {
	id   int
	name string
	GPA  float64
	pob  string
}

func AssignStudent(s *[100]students) {
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
func (s students) Print() {
	fmt.Println(s.name)
	fmt.Println(s.GPA)
	fmt.Println(s.id)
	fmt.Println(s.pob)

}
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

func AssignStudentstack(s *[5]students) {
	for i := 0; i < 2; i++ {
		s[i].Push()
	}
}
func (s *students) Push() {
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
func Pop(s *[5]students) students {
	counter--
	temp := s[counter]
	s[counter].name = ""
	s[counter].id = 0
	s[counter].GPA = 0
	s[counter].pob = ""
	return temp
}
func Maxgpastack(s *[5]students) students {
	var temp [5]students
	for i := 0; i < 5; i++ {
		temp[i] = s[i]
	}
	var result students
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
func Mingpastack(s *[5]students) students {
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
	return result
}
func AssignStudentarr2(s *[100]students) {
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
