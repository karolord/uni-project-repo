package main

type Student struct {
	Name    string
	ID      int
	Address string
}

type NodeGeneral struct {
	Key    *Student
	Left   *NodeGeneral
	Right  *NodeGeneral
	Middle *NodeGeneral
}

func (g *NodeGeneral) InsertNodeGeneral(S *Student) {
	if g.Key == nil {
		g.Key = S
		return
	}
	if S.Name[0] == g.Key.Name[0] {
		if g.Middle == nil {
			g.Middle = &NodeGeneral{Key: S}
		} else {
			g.Middle.InsertNodeGeneral(S)
		}
	} else if S.Name[0] > g.Key.Name[0] {
		if g.Right == nil {
			g.Right = &NodeGeneral{Key: S}
		} else {
			g.Right.InsertNodeGeneral(S)
		}
	} else if S.Name[0] < g.Key.Name[0] {
		if g.Left == nil {
			g.Left = &NodeGeneral{Key: S}
		} else {
			g.Left.InsertNodeGeneral(S)
		}
	}
}
func main() {
}
