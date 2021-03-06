package main

import (
	"fmt"
)

var n int
var list Linkedlist
var hashtable Hashmap

type Hashmap struct {
	hmap []*Linkedlist
}

//functions for hashmap
func (h *Hashmap) Inserthash(value string) {
	index := hashfunction(value, len(h.hmap))
	h.hmap[index].insertsorted(value)
}
func (h *Hashmap) Deletehash(value string) {
	index := hashfunction(value, len(h.hmap))
	h.hmap[index].deleteLinkedlist(value)
}
func hashfunction(key string, length int) int {
	sum := int(key[len(key)-2])
	return (sum - 19) % length
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
func (l *Linkedlist) insertsorted(value string) {
	newNode := &Node{name: value}
	if l.head == nil {
		newNode.next = l.head
		l.head = newNode
	} else if newNode.name[len(newNode.name)-1] <= l.head.name[len(l.head.name)-1] {
		newNode.next = l.head
		l.head = newNode
	} else {
		tmp := l.head
		for tmp != nil {
			if tmp.next == nil {
				tmp.next = newNode
				break
			} else if newNode.name[len(newNode.name)-1] < tmp.next.name[len(tmp.next.name)-1] {
				newNode.next = tmp.next
				tmp.next = newNode
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
	hashtable := Initalizemap(26)
	n = 9
	hashtable.Inserthash("karo")
	hashtable.Inserthash("kosar")
	hashtable.Inserthash("kerm")
	hashtable.Inserthash("kerman")
	hashtable.Inserthash("ahmad")
	hashtable.Inserthash("omer")
	//hashtable.requirement3()
	//hashtable.requirement4()
	hashtable.requirement6()

}
func requirement1() {
	fmt.Println("Please enter the amount of names you want:")
	fmt.Scanf("%d", &n)
	fmt.Scanln()
	requirement2()
}

func requirement2() {
	for i := 0; i < n; i++ {
		fmt.Println("Please enter the name:")
		var s string
		fmt.Scanln(&s)
		list.insertLinkedlist(s)
	}
}
func (h *Hashmap) requirement3() {
	node := list.head
	for node != nil {
		h.Inserthash(node.name)
		node = node.next
	}
	h.Printhashtable()
}

func (h *Hashmap) Printhashtable() {
	for i := 0; i < len(h.hmap); i++ {
		node := h.hmap[i].head
		for node != nil {
			fmt.Printf("%s ", node.name)
			node = node.next
		}
	}
}

func (h *Hashmap) deleteletter(c string) {
	x := (int(c[0]) - 19) % 26
	h.hmap[x].head = nil
}

func (h *Hashmap) requirement4() {
	h.deleteletter("r")
	h.deleteletter("z")
	h.Printhashtable()
}
func requirement5() {
	node := list.head
	for node != nil {
		fmt.Printf("%s ", node.name)
		node = node.next
	}
	fmt.Println()
	var prev *Node
	var next *Node
	current := list.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
	node = list.head
	for node != nil {
		fmt.Printf("%s ", node.name)
		node = node.next
	}
}
func (h *Hashmap) requirement6() {
	table6 := Initalizemap(n / 3)
	for i := 0; i < len(h.hmap); i++ {
		node := h.hmap[i].head
		for node != nil {
			table6.Inserthash(node.name)
			node = node.next
		}
	}
	table6.Printhashtable()
}
