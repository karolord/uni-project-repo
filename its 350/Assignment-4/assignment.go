package main

type student struct {
	Name    string
	ID      int
	Address string
}

type Node struct {
	Key    *student
	Left   *Node
	Right  *Node
	Middle *Node
}

func (g *Node) InsertNode(s *student) {
	if g.Key == nil {
		g.Key = s
		return
	}
	if s.Name[0] == g.Key.Name[0] {
		if g.Middle == nil {
			g.Middle = &Node{Key: s}
		} else {
			g.Middle.InsertNode(s)
		}
	} else if s.Name[0] > g.Key.Name[0] {
		if g.Right == nil {
			g.Right = &Node{Key: s}
		} else {
			g.Right.InsertNode(s)
		}
	} else if s.Name[0] < g.Key.Name[0] {
		if g.Left == nil {
			g.Left = &Node{Key: s}
		} else {
			g.Left.InsertNode(s)
		}
	}
}
func main() {
	Root := &Node{}
	s1 := &student{"karo", 5, "aaaaaa"}
	s2 := &student{"kosar", 5, "aaaaaa"}
	s3 := &student{"kahmad", 5, "aaaaaa"}
	s4 := &student{"zebra", 5, "aaaaaa"}
	s5 := &student{"black", 5, "aaaaaa"}
	s6 := &student{"yellow", 5, "aaaaaa"}
	Root.InsertNode(s1)
	Root.InsertNode(s2)
	Root.InsertNode(s3)
	Root.InsertNode(s4)
	Root.InsertNode(s5)
	Root.InsertNode(s6)

}
