package main

type bst struct {
	x     int
	left  *bst
	right *bst
}
type List struct {
	Tail *Node
	Head *Node
}

type Node struct {
	key  int
	Next *Node
	Prev *Node
}

func (l *List) insertNode(x int) {
	newNode := &Node{key: x}
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

func (b *bst) insertbst(x int) {
	if x < b.x {
		if b.left == nil {
			b.left = &bst{x: x}
		} else {
			b.left.insertbst(x)
		}
	} else if x > b.x {
		if b.right == nil {
			b.right = &bst{x: x}
		} else {
			b.right.insertbst(x)
		}
	}

}

func (b *bst) bsttodll(head *List) {
	if b != nil {

		if b.left != nil {
			b.left.bsttodll(head)
		}
		head.insertNode(b.x)
	}
	if b.right != nil {
		b.right.bsttodll(head)
	}

}
func main() {
	binarysearch := &bst{x: 5}
	l := &List{}
	binarysearch.insertbst(6)
	binarysearch.insertbst(2)
	binarysearch.insertbst(4)
	binarysearch.insertbst(3)
	binarysearch.insertbst(1)
	binarysearch.insertbst(7)
	binarysearch.insertbst(8)
	binarysearch.bsttodll(l)

}
