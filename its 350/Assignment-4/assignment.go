/*Kosar N. Aziz kn18011@auis.edu.krd
Karo K. Rasool kk19046@auis.edu.krd*/
/*
Requirement 4: T(n) = 5n^2 + 35n, Big O = O(n^2)
Requirement 5: T(n) = 11nlog n + 8log n, Big O = O(nlog(n))
*/
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
	newNode := &Nodelist{key: s} // 4
	if l.Head == nil {           // 3
		l.Head = newNode // 2
		l.Tail = newNode // 2
	} else {
		currentNode := l.Head         // 2
		for currentNode.Next != nil { // 3n + 3
			currentNode = currentNode.Next // 2n
		}
		newNode.Prev = currentNode // 2
		currentNode.Next = newNode // 2
		l.Tail = newNode           // 2
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
	tmp := &BinaryRoot        // 2
	tmp.bsttodll(&doublelist) // 35 n + 5n^2
}

// in-order transverse over the binary search tree and transfer to doubly Linkedlist
/*the functon will call it self n times because it has to transverse the entire list and
it will also call the insertnode function n times because it has to transfer n amount of students
to the linked list so in total those two functons time complexity will be 13n(for the recursion) + (5n^2 + 22n) for the insertnode function
time complexity will be 5n^2 + 35n
*/
func (b *BinaryTree) bsttodll(Head *DoubleLinkedList) {
	if b != nil { // 3
		if b.Left != nil { // 3
			b.Left.bsttodll(Head) //2
		}
		Head.insertNode(b.Key) //2
	}
	if b.Right != nil { // 3
		b.Right.bsttodll(Head)
	}
}

func Requirement5() {
	var id int                                                // 0
	fmt.Println("Please Enter an id you want to search for:") // 1
	fmt.Scanln(&id)                                           // 2
	data := BinarySearch(id, doublelist.Head)                 // 3
	fmt.Println(data)                                         // 2
}
func BinarySearch(ID int, head *Nodelist) string {
	size := n / 2 // 3
	middle := NodeReturnAfter(head, size)
	for middle.key.ID != ID { // 3n + 3
		if ID > middle.key.ID { // 3n
			size = size / 2 // 3n
			tmp := NodeReturnAfter(middle, size)
			if tmp.key.ID == middle.key.ID { // 3n
				break // 0n
			}
			middle = tmp // 2n
		} else {
			size = size / 2 // 3n
			tmp := NodeReturnPrev(middle, size)
			if tmp.key.ID == middle.key.ID { // 3n
				break // 0n
			}
			middle = tmp // 2n
		}
	}
	if middle.key.ID != ID { // 3
		return "ERROR ID NOT FOUND" // 0
	}
	return "The Name is " + middle.key.Name + " The Address is " + middle.key.Address // 2

}

func NodeReturnAfter(node *Nodelist, size int) *Nodelist {
	tmp := node    // 2
	if size == 0 { // 2
		size += 1 // 3
	}
	for i := 0; i < size; i++ { // 1 + 3 + 3
		if tmp.Next != nil { // 2
			tmp = tmp.Next // 2
		}
	}
	return tmp // 1
}
func NodeReturnPrev(node *Nodelist, size int) *Nodelist {
	tmp := node    // 2
	if size == 0 { // 2
		size += 1 // 3
	}
	for i := 0; i < size; i++ { // 1 + 3 + 3
		if tmp.Prev != nil { // 2
			tmp = tmp.Prev // 2
		}
	}
	return tmp // 1
}

/*the functions NodeReturnPrev and NodeReturnAfter are gonna be ran log n times
the total time complexity will be n *log n because we have to transverse throught the entire list
each Node return is gonna take 11n + 8 then we multiply it by log n
time complexity will be 11n log n + 8 log n
*/
