package main

import (
	"fmt"
)

var n int

type Hashmap struct {
	hmap []*Linkedlist
}

//functions for hashmap
func (h *Hashmap) Inserthash(value string) {
	index := hashfunction(value)
	h.hmap[index].insertLinkedlist(value)
}
func (h *Hashmap) Deletehash(value string) {
	index := hashfunction(value)
	h.hmap[index].deleteLinkedlist(value)
}
func hashfunction(key string) int {
	sum := int(key[len(key)-2])
	return (sum - 19) % 26
}
func Initalizemap(size int) *Hashmap {
	result := &Hashmap{}
	result.hmap = make([]*Linkedlist, size)
	for i := range result.hmap {
		result.hmap[i] = &Linkedlist{}
	}
	return result
}

//functions for linked linked
type Linkedlist struct {
	head   *Node
	length int
	tail   *Node
}

type Node struct {
	name string
	next *Node
}

func (l *Linkedlist) insertLinkedlist(value string) {
	newNode := &Node{name: value}
	if l.head == nil {
		newNode.next = l.head
		l.head = newNode
	} else if newNode.name[len(newNode.name)-1] < l.head.name[len(l.head.name)-1] {
		newNode.next = l.head
		l.head = newNode
	} else {
		tmp := l.head
		for tmp.next != nil {
			if newNode.name[len(newNode.name)-1] < tmp.next.name[len(l.head.name)-1] {
				newNode.next = tmp.next
				tmp.next = newNode
				break
			} else if tmp.next.next == nil {
				tmp.next.next = newNode
				break
			}
			tmp = tmp.next
		}
	}

}

func (l *Linkedlist) deleteLinkedlist(value string) {
	if l.head.name == value {
		l.head = l.head.next
		return
	}

	prevNode := l.head
	for prevNode.next != nil {
		if prevNode.next.name == value {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func main() {
	testmap := Initalizemap(26)
	fmt.Println(tolower("KKKKARaro"))
	testmap.Inserthash("kaao")
	testmap.Inserthash("kaac")
	testmap.Inserthash("kaax")
	testmap.Inserthash("kaab")
	testmap.Inserthash("kaal")
}
