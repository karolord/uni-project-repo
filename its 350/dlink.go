package main

import "fmt"

type stu struct {
	name    string
	ID      int
	address string
	left    *stu
	right   *stu
	mid     *stu
	parent  *stu
}

var n int
var name string
var ID int
var address string
var p int
var change *stu
var root1 *stu

///var s *stu

func main() {
	p = 1
	s1 := add1("name", 5, "address", nil)
	fmt.Println(root1)
	change = s1
	name = "karo"
	address = "karo"
	Add(4)

}

func requ1() {
	fmt.Print("Enter the n number of students: ")
	fmt.Scan(&n)

	for x := 1; x <= n; x++ {
		fmt.Printf("Enter name %d: ", x)
		fmt.Scan(&name)
		fmt.Printf("Enter ID %d: ", x)
		fmt.Scan(&ID)
		fmt.Printf("Enter address %d: ", x)
		fmt.Scan(&address)
		fmt.Println()

		if p == 1 {
			s1 := add1(name, ID, address, nil)
			change = s1
			p++

		} else {
			Add(ID)

		}

	}
}

func Add(idi int) {
	c := root1
	for c != nil {
		if idi < c.ID {
			//go left
			left := c.left
			if left == nil {
				add1(name, idi, address, left)
				return
			}
			c = left
		} else {
			///go right
			right := c.right
			if right == nil {
				add1(name, idi, address, right)
				break
			}
			c = right
		}

	}
}
func add1(na string, id int, addr string, target *stu) *stu {
	s := stu{na, id, addr, nil, nil, nil, nil}
	fmt.Println(target)
	if root1 == nil {
		root1 = &s
		return &s
	} else {

		if s.ID < target.ID {
			target.left = &s
		} else if s.ID > target.ID {
			target.right = &s
		}

		s.parent = target

		return &s
	}
}
