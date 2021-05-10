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
	requ1()

}

func requ1() {
	n = 5
	p = 1
	for x := 1; x <= n; x++ {
		name = "karo"
		ID = 6
		address = "karo"
		if p == 1 {
			s1 := add1(name, ID, address, nil)
			change = s1
			p++

		} else {
			s2 := add1(name, ID, address, change)
			change = s2

		}

	}
}
func add1(na string, id int, addr string, target *stu) *stu {
	s := stu{na, id, addr, nil, nil, nil, nil}
	if root1 == nil {
		root1 = &s
		return &s
	}
	fmt.Printf("error check")
	if s.ID < target.ID {
		target.left = &s
	} else if s.ID > target.ID {
		target.right = &s
	} else {
		target.mid = &s
	}

	s.parent = target

	return &s
}
