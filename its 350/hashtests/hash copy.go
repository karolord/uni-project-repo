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
func (h *Hashmap) Inserthashsort(value string) {
	index := hashfunctionsort(value)
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
func hashfunctionsort(key string) int {
	sum := 0
	return sum
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

func main() {
	n = 10
	test := Initalizemap()
	fmt.Println(test)
	test.Inserthashsort("ab")
	test.Inserthashsort("a2we")
	test.Inserthashsort("abcc")
	test.Inserthashsort("abkkh")
	test.Inserthashsort("abdssadw")
	test.Inserthashsort("abjyjy")
	test.Inserthashsort("batu6ut5")
	test.Inserthashsort("cd")
	test.Inserthashsort("aa")
	test.Inserthashsort("as")
	fmt.Println(test)
}
