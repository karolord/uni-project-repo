/*Kosar N. Aziz kn18011@auis.edu.krd
Karo K. Rasool kk19046@auis.edu.krd*/
/*
Requirement 5: T(n) = 16n + 16, Big O = O(n)
Requirement 6: T(n) = 21n^3 + 49n^2 + 45n + 39, Big O = O(n^3)
*/
/*
Citations:
https://www.youtube.com/watch?v=zLnJcAt1aKs
*/

package main

import (
	"fmt"
)

var n int
var list Linkedlist

//functions for hashtable
type Hashtable struct {
	hmap []*Linkedlist
}

func (h *Hashtable) Inserthash(value string) {
	index := hashfunction(value, len(h.hmap)) // 6
	h.hmap[index].insertsorted(value) // 3
}
func hashfunction(key string, length int) int {
	sum := int(key[len(key)-2])
	return (sum - 19) % length
}
func Initalizemap(size int) *Hashtable {
	result := &Hashtable{} // 3
	result.hmap = make([]*Linkedlist, size) // 4
	for i := range result.hmap { // 3n
		result.hmap[i] = &Linkedlist{} // 3n
	}
	return result // 1
}
func (h *Hashtable) Printhashtable() {
	for i := 0; i < len(h.hmap); i++ { // 2, 4n + 4, 3n
		node := h.hmap[i].head // 4n
		for node != nil { // 2n^2
			fmt.Printf("%s ", node.name) // 2n^2
			node = node.next // 2n^2
		}
	}
}

//functions for linked linked
type Linkedlist struct {
	head   *Node
	length int
}

// node for data and pointer
type Node struct {
	name string
	next *Node
}

func (l *Linkedlist) insertLinkedlist(value string) {
	newNode := &Node{name: value}
	newNode.next = l.head
	l.head = newNode
}

//sorting for the hashtable
func (l *Linkedlist) insertsorted(value string) {
	newNode := &Node{name: value} // 5
	if l.head == nil { // 2
		newNode.next = l.head // 2
		l.head = newNode // 2
	} else if newNode.name[len(newNode.name)-1] <= l.head.name[len(l.head.name)-1] { // 9
		newNode.next = l.head // 2
		l.head = newNode // 2
	} else {
		tmp := l.head // 3
		for tmp != nil { // 2n
			if tmp.next == nil { // 2n
				tmp.next = newNode // 2n
				break // 0
			} else if newNode.name[len(newNode.name)-1] < tmp.next.name[len(tmp.next.name)-1] { // 9n
				newNode.next = tmp.next // 2n
				tmp.next = newNode // 2n
				break // 0
			}
			tmp = tmp.next // 2n
		}
	}
}

// Requirements
func requirement1() {
	fmt.Println("Please enter the amount of names you want:")
	fmt.Scanf("%d", &n)
	fmt.Scanln()
	requirement2()
	printlinklist()
}

func requirement2() {
	for i := 0; i < n; i++ {
		fmt.Println("Please enter the name:")
		var s string
		fmt.Scanln(&s)
		list.insertLinkedlist(s)
	}
}
func printlinklist() {
	fmt.Println("Linkedlist:")
	node := list.head
	for node != nil {
		fmt.Printf("%s ", node.name)
		node = node.next
	}
	fmt.Println()
}
func (hashtable *Hashtable) requirement3() {
	node := list.head
	for node != nil {
		hashtable.Inserthash(node.name)
		node = node.next
	}
}

func (hashtable *Hashtable) deleteletter(c string) {
	x := (int(c[0]) - 19) % 26
	hashtable.hmap[x].head = nil
}
func (hashtable *Hashtable) requirement4() {
	hashtable.deleteletter("r")
	hashtable.deleteletter("z")
	fmt.Println()
	fmt.Println("Deleted R and Z: ")
	hashtable.Printhashtable()
	fmt.Println()
}
func requirement5() {
	fmt.Println("the reversed link list:") // 1
	var prev *Node // 1
	var next *Node // 1
	current := list.head // 3
	for current != nil { // 2n + 2
		next = current.next // 2n
		current.next = prev // 2n
		prev = current // 2n
		current = next // 2n
	}
	list.head = prev // 2
	node := list.head // 3
	for node != nil { // 2n + 2
		fmt.Printf("%s ", node.name) // 2n
		node = node.next // 2n
	}
	fmt.Println() // 1
}
func (hashtable *Hashtable) requirement6() {
	if n < 3 { // 2
		fmt.Println("The key size is not divisible by 3") // 1
		return // 0
	}
	fmt.Println("Resized hashtable: ") // 1
	table6 := Initalizemap(n / 3) // 4 + 6n + 8
	for i := 0; i < len(hashtable.hmap); i++ { // 2, 4n + 4, 3n
		node := hashtable.hmap[i].head // 4n
		for node != nil { // 2n^2
			table6.Inserthash(node.name) // 2n^2 + 9n^2 + 21n^3 + 27n^2
			node = node.next // 2n^2
		}
	}
	table6.Printhashtable() // 1 + 17n + 12 
}
func main() {
	tablereq4 := Initalizemap(26)
	hashtable := Initalizemap(26)
	requirement1()
	fmt.Println("Hashtable:")
	hashtable.requirement3()
	hashtable.Printhashtable()
	tablereq4.requirement3()
	tablereq4.requirement4()
	requirement5()
	hashtable.requirement6()
}
