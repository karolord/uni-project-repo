package main

import "fmt"

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
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % n
}
func (h *Hashmap) searchhash(value string) bool {
	return false
}
func Initalizemap() *Hashmap {
	result := &Hashmap{}
	result.hmap = make([]*Linkedlist, n)
	for i := range result.hmap {
		result.hmap[i] = &Linkedlist{}
	}
	return result
}

//functions for linked linked
type Linkedlist struct {
	head *Node
	tail *Node
}
type Node struct {
	name string
	next *Node
}

func (l *Linkedlist) insertLinkedlist(value string) {
	newNode := &Node{name: value}
	newNode.next = l.head
	l.head = newNode
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

func (h *Hashmap) sortlist() {
	for i := range h.hmap {
		if h.hmap[i].head != nil {
			prevNode := h.hmap[i].head
			for prevNode.next != nil {
				if int(prevNode.next.name[0]) <= int(prevNode.name[0]) {
					*prevNode, *prevNode.next = *prevNode.next, *prevNode
				}
				prevNode = prevNode.next
			}
		}
	}
}
func main() {
	n = 4
	test := Initalizemap()
	fmt.Println(test)
	test.Inserthash("aB")
	test.Inserthash("bA")
	test.Inserthash("aD")
	test.Inserthash("bC")
	test.sortlist()
	fmt.Println(test)
}
