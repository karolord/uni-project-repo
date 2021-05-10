package main

import "fmt"

var n int

type List struct {
	Tail *Node
	Head *Node
}

type Node struct {
	key  int
	Next *Node
	Prev *Nodes
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

func binarysearch(key int, node *Node) *Node {
	size := n / 2
	middlenode := NodeReturntAfter(node, size)
	for middlenode.key != key {
		if key > middlenode.key {
			size = size / 2
			tmp := NodeReturntAfter(middlenode, size)
			if tmp.key == middlenode.key {
				break
			}
			middlenode = tmp
		} else if key < middlenode.key {
			size = size / 2
			tmp := NodeReturntPrev(middlenode, size)
			if tmp.key == middlenode.key {
				break
			}
			middlenode = tmp
		}
	}
	if middlenode.key != key {
		return &Node{key: -1}
	}
	return middlenode
}
func NodeReturntAfter(node *Node, size int) *Node {
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
func NodeReturntPrev(node *Node, size int) *Node {
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
func main() {
	var list List
	list.insertNode(&Node{key: 1})
	list.insertNode(&Node{key: 2})
	list.insertNode(&Node{key: 3})
	list.insertNode(&Node{key: 4})
	list.insertNode(&Node{key: 5})
	list.insertNode(&Node{key: 7})
	n = 6
	fmt.Println(binarysearch(6, list.Head))
}
