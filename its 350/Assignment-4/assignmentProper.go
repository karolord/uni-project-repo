package main

import "fmt"

var n int
var GeneralRoot GeneralTree
var BinaryRoot BinaryTree
var doublelist DoubleLinkedList

type student struct {
	Name    string
	ID      int
	Address string
}

type GeneralTree struct {
	Key    *student
	Left   *GeneralTree
	Right  *GeneralTree
	Middle *GeneralTree
}

func (g *GeneralTree) InsertGeneralNode(s *student) {
	if g.Key == nil {
		g.Key = s
		return
	}
	if s.Name[0] == g.Key.Name[0] {
		if g.Middle == nil {
			g.Middle = &GeneralTree{Key: s}
		} else {
			g.Middle.InsertGeneralNode(s)
		}
	} else if s.Name[0] > g.Key.Name[0] {
		if g.Right == nil {
			g.Right = &GeneralTree{Key: s}
		} else {
			g.Right.InsertGeneralNode(s)
		}
	} else if s.Name[0] < g.Key.Name[0] {
		if g.Left == nil {
			g.Left = &GeneralTree{Key: s}
		} else {
			g.Left.InsertGeneralNode(s)
		}
	}
}

type BinaryTree struct {
	Key   *student
	Left  *BinaryTree
	Right *BinaryTree
}

func (B *BinaryTree) InsertBinaryNode(s *student) {
	if B.Key == nil {
		B.Key = s
		return
	}
	if s.ID > B.Key.ID {
		if B.Right == nil {
			B.Right = &BinaryTree{Key: s}
		} else {
			B.Right.InsertBinaryNode(s)
		}
	} else if s.ID < B.Key.ID {
		if B.Left == nil {
			B.Left = &BinaryTree{Key: s}
		} else {
			B.Left.InsertBinaryNode(s)
		}
	}
}

type DoubleLinkedList struct {
	Tail *Nodelist
	Head *Nodelist
}
type Nodelist struct {
	key  *student
	Prev *Nodelist
	Next *Nodelist
}

func (l *DoubleLinkedList) insertNode(s *student) {
	newNode := &Nodelist{key: s}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		newNode.Prev = currentNode
		currentNode.Next = newNode
		l.Tail = newNode
	}
}

func main() {
	Requirement1()
	Requirement4()
	Requirement5()
}
func Requirement1() {
	var Name string
	var ID int
	var Address string
	fmt.Println("Please enter the number of students you would like to Enter: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Please enter the name of the student: ")
		fmt.Scanln(&Name)
		fmt.Println("Please enter the ID of the student: ")
		fmt.Scanln(&ID)
		fmt.Println("Please enter the Address of the student: ")
		fmt.Scanln(&Address)
		Requirement2(&student{Name, ID, Address})
		Requirement3(&student{Name, ID, Address})
	}
}
func Requirement2(s *student) {
	tmp := &GeneralRoot
	tmp.InsertGeneralNode(s)
}
func Requirement3(s *student) {
	tmp := &BinaryRoot
	tmp.InsertBinaryNode(s)
}

//Converting Binary Tree To a Doubly Linkedlist
func Requirement4() {
	tmp := &BinaryRoot
	tmp.bsttodll(&doublelist)
}

// in-order transverse over the binary search tree and transfer to doubly Linkedlist
func (b *BinaryTree) bsttodll(Head *DoubleLinkedList) {
	if b != nil {
		if b.Left != nil {
			b.Left.bsttodll(Head)
		}
		Head.insertNode(b.Key)
	}
	if b.Right != nil {
		b.Right.bsttodll(Head)
	}
}

func Requirement5() {
	var id int
	fmt.Println("Please Enter an id you want to search for:")
	fmt.Scanln(&id)
	data := BinarySearch(id, doublelist.Head)
	fmt.Println(data)
}
func BinarySearch(ID int, head *Nodelist) string {
	size := n / 2
	middle := NodeReturnAfter(head, size)
	for middle.key.ID != ID {
		if ID > middle.key.ID {
			size = size / 2
			tmp := NodeReturnAfter(middle, size)
			if tmp.key.ID == middle.key.ID {
				break
			}
			middle = tmp
		} else {
			size = size / 2
			tmp := NodeReturnPrev(middle, size)
			if tmp.key.ID == middle.key.ID {
				break
			}
			middle = tmp
		}
	}
	if middle.key.ID != ID {
		return "ERROR ID NOT FOUND"
	}
	return "The Name is " + middle.key.Name + " The Address is " + middle.key.Address

}

func NodeReturnAfter(node *Nodelist, size int) *Nodelist {
	tmp := node
	if size == 0 {
		size += 1
	}
	for i := 0; i < size; i++ {
		if tmp.Next != nil {
			tmp = tmp.Next
		}
	}
	return tmp
}
func NodeReturnPrev(node *Nodelist, size int) *Nodelist {
	tmp := node
	if size == 0 {
		size += 1
	}
	for i := 0; i < size; i++ {
		if tmp.Prev != nil {
			tmp = tmp.Prev
		}
	}
	return tmp
}
