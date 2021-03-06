package main

import (
	"fmt"
)

//citation https://stackoverflow.com/questions/2439453/using-a-pointer-to-array
var n int
var counter int // counter for stack
func main() {
	numberassignment()
	Requirement1()
	numberassignmentV2()
	Requirement2()
	numberassignment()
	Requirement3()
}

type student struct {
	ID   int
	Name string
	GPA  float64
	Pob  string // place of birth
}

func numberassignment() {
	for {
		fmt.Println("Please enter the number of students: ")
		fmt.Scanln(&n)
		if n <= 100 {
			break
		}
		fmt.Println("Error please enter a value less then or equal to 100")
	}
}

func numberassignmentV2() {
	for {
		fmt.Println("Please enter the number of students: ")
		fmt.Scanln(&n)
		if n <= 5 {
			break
		}
		fmt.Println("Error please enter a value less then or equal to 5")
	}
}

// assignstudent for requirement 1 and 2
func Requirement1() {
	var arr [100]student
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
	fmt.Println("The student name is ", s.Name)
	fmt.Println("The student GPA is", s.GPA)
	fmt.Println("The student ID is", s.ID)
	fmt.Println("The student Place of Birth is ", s.Pob)
}

// gpa calculation
func Maxgpa(s [100]student) {
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

// the complexity is 12 + 11n big o us O(n)

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
	s[index].Print()
}

// the complexity is 12 +11n big o is O(n)
//assignstudent,  push and pop for req 3 and 4
func Requirement2() {
	var stack [5]student
	AssignStudentstack(&stack)
	Maxgpastack(&stack)
	Mingpastack(&stack)
}
func AssignStudentstack(s *[5]student) {
	for i := 0; i < n; i++ {
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
func Pop(s *[5]student) student {
	counter--
	temp := s[counter]
	s[counter].Name = ""
	s[counter].ID = 0
	s[counter].GPA = 0
	s[counter].Pob = ""
	return temp
}

// time complexity of pop is 8 big o is o(1)
func Maxgpastack(s *[5]student) {
	var temp [5]student
	for i := 0; i < n; i++ {
		temp[i] = s[i]
	}
	var result student
	tempcount := counter
	for i := tempcount - 1; i >= 0; i-- {
		compare := Pop(&temp)
		if compare.GPA > result.GPA {
			result = compare
		}
	}
	counter = tempcount
	fmt.Println("The student with the highest GPA is ")
	result.Print()
}

// time complexity is 22 + 29n big o is o(1)
func Mingpastack(s *[5]student) {
	var temp [5]student
	for i := 0; i < n; i++ {
		temp[i] = s[i]
	}
	result := temp[4]
	tempcount := counter
	for i := tempcount - 1; i >= 0; i-- {
		compare := Pop(&temp)
		if compare.GPA < result.GPA {
			result = compare
		}
	}
	counter = tempcount
	fmt.Println("The student with the lowest GPA is ")
	result.Print()
}

// time complexity is 22 + 29n big o is o(1)

// Requirement 5 and 6
func Requirement3() {
	var arr2 [100]requirement5
	AssignStudentreq5(&arr2)
	Maxgpaarr2(arr2)
	Mingpaarr2(arr2)
}

type requirement5 struct {
	students student
}

// new struct so users dont have direct access to student struct
//the assign will do only n numbers of students
func AssignStudentreq5(r *[100]requirement5) {
	for i := 0; i < n; i++ {
		fmt.Println("Please enter the Name of the student: ")
		fmt.Scanln(&r[i].students.Name)
		fmt.Println("Please enter the ID of the student: ")
		fmt.Scanln(&r[i].students.ID)
		for {
			fmt.Println("Please enter the GPA of the student: ")
			fmt.Scanln(&r[i].students.GPA)
			if r[i].students.GPA >= 0 && r[i].students.GPA <= 4 {
				break
			}
			fmt.Println("Error please enter a valid GPA")
		}
		fmt.Println("Please enter where the student was born: ")
		fmt.Scanln(&r[i].students.Pob)
	}
}

func Maxgpaarr2(r [100]requirement5) {
	var Index int
	temp := r[0].students.GPA
	for i := 0; i < n; i++ {
		if r[i].students.GPA > temp {
			temp = r[i].students.GPA
			Index = i
		}
	}
	fmt.Println("The student with the highest GPA is ")
	r[Index].students.Print()

}

// the complexity is 12 +11n big o is O(n)

func Mingpaarr2(r [100]requirement5) {
	var Index int
	temp := r[0].students.GPA
	for i := 0; i < n; i++ {
		if r[i].students.GPA < temp {
			temp = r[i].students.GPA
			Index = i
		}
	}
	fmt.Println("The student with the lowest GPA is ")
	r[Index].students.Print()
}

// the complexity is 12 +11n big o is O(n)
