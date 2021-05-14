/////// as17176   Akar Sherko Abuabkir
/////// ra17192   Rawa Ali Mohammed

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
var change *stu
var root1 *stu

///var s *stu

func main() {
	requ1()
	fmt.Println(root1)
	fmt.Println(root1.mid)
	fmt.Println(root1.mid.mid)

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
		Add(ID, root1)
		Add2(ID, root2)
	}
}

func Add(idi int, treeroot *stu) {
	s := stu{name, ID, address, nil, nil, nil, nil}
	if root1 == nil {
		root1 = &s
		return
	}
	c := treeroot
	if idi < c.ID {
		// GO Left
		if c.left == nil {
			c.left = &stu{name, ID, address, nil, nil, nil, nil}
		} else {
			Add(idi, c.left)
		}
	} else if idi > c.ID {
		// GO Left
		if c.right == nil {
			c.right = &stu{name, ID, address, nil, nil, nil, nil}
		} else {
			Add(idi, c.right)
		}
	}
}
func Add2(idi int, treeroot *stu) {
	s := stu{name, ID, address, nil, nil, nil, nil}
	if root2 == nil {
		root2 = &s
		return
	}
	c := treeroot
	if idi < c.ID {
		// GO Left
		if c.left == nil {
			c.left = &stu{name, ID, address, nil, nil, nil, nil}
		} else {
			Add2(idi, c.left)
		}
	} else if idi > c.ID {
		// GO Left
		if c.right == nil {
			c.right = &stu{name, ID, address, nil, nil, nil, nil}
		} else {
			Add2(idi, c.right)
		}
	} else {
		if c.mid == nil {
			c.mid = &stu{name, ID, address, nil, nil, nil, nil}
		} else {
			Add2(idi, c.mid)
		}
	}
}
