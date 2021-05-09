package main

type Student struct {
	Name    string
	ID      int
	Address string
}

type GeneralTree struct {
	value  *Student
	Right  *GeneralTree
	Left   *GeneralTree
	Middle *GeneralTree
}

func (g *GeneralTree) InsertGeneralTree(S *Student) {
	if g.value == nil {
		g.value = S
		return
	}
	if S.Name[0] < g.value.Name[0] {
		if g.Left == nil {
			g.Left = &GeneralTree{value: S}
		} else {
			g.Left.InsertGeneralTree(S)
		}
	} else if S.Name[0] > g.value.Name[0] {
		if g.Right == nil {
			g.Right = &GeneralTree{value: S}
		} else {
			g.Right.InsertGeneralTree(S)
		}
	} else if S.Name[0] == g.value.Name[0] {
		if g.Middle == nil {
			g.Middle = &GeneralTree{value: S}
		} else {
			g.Middle.InsertGeneralTree(S)
		}
	}

}

func main() {

}
