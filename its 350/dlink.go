package main

type Student struct {
	grade    int
	children [2]*Student
	parent   *Student
	height   int
}

var root *Student

func calcHeight(s *Student) {
	c := s.parent
	counter := 0
	if c.children[1] == nil && c.children[0] == nil {
		for c != nil {
			counter++
			if counter >= c.height {
				c.height = counter
			}
			c = c.parent
		}
	}
}

func add(g int, target *Student, dir int) *Student {
	s := Student{g, nil, nil, nil, 0}
	if root == nil {
		root = &s
		return &s
	}

	target.children[dir] = &s

	s.parent = target
	calcHeight(&s)
	return &s
}

func main() {
	s1 := add(100, nil, 0)
	s2 := add(85, s1, 0)
	s3 := add(70, s1, 1)
	s4 := add(75, s1, 2)
	s5 := add(20, s3, 0)
}
