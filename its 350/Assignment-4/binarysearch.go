package main

type List struct {
	Tail *Node
	Head *Node
}

type Node struct {
	key  int
	Next *Node
	Prev *Node
}

func (l *List) insertNode(newNode *Node) {
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

}
