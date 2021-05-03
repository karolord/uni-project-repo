package main

import "fmt"

//Karo Rasool kk19046@auis.edu.krd*/

type tree struct {
	value  int
	left   *tree
	right  *tree
	parent *tree
	height int
}

func (t *tree) Insert(k int) {
	x := &tree{k, nil, nil, nil, 0}
	if t.value < k {
		if t.right == nil {
			x.parent = t
			t.right = x
			if t.left == nil {
				calcheight(x)
			}
		} else {
			t.right.Insert(k)
		}
	} else if t.value > k {
		if t.left == nil {
			x.parent = t
			t.left = x
			if t.right == nil {
				calcheight(x)
			}
		} else {
			t.left.Insert(k)
		}
	}

}
func calcheight(t *tree) {
	c := t.parent
	X := 0
	for c != nil {
		X++
		if X+1 > c.height {
			c.height = c.height + 1
		}
		c = c.parent
	}
}
func main() {
	x := &tree{5, nil, nil, nil, 0}
	x.Insert(555)
	x.Insert(554)
	x.Insert(553)
	x.Insert(552)
	x.Insert(551)
	x.Insert(550)
	x.Insert(556)
	x.Insert(557)
	x.Insert(558)
	x.Insert(3)
	x.Insert(2)
	fmt.Println(x.height)
}
