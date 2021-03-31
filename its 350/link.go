package main

// homework create print similar to this but print name of the forth student
// pointer tail
type Student struct {
	Name string
	GPA  float64
	ID   int
	Node *Student
}

var head *Student

func addtoend(n string, g float64, i int) {
	s := Student{n, g, i, nil}
	if head == nil {
		head = &s
	} else {
		c := head
		for c != nil {
			c = c.Node
		}
		c.Node = &s
	}
}
func main() {
	addtoend("ali", 3.2, 155555)
	addtoend("karo", 3.22, 55555)

}
func addtostart(n string, g float64, i int) {
	s := Student{n, g, i, nil}
	s.Node = head
	head = &s
}
func addafter(n string, g float64, i int, std *Student) {
	s := Student{n, g, i, nil}
	s.Node = std.Node
	std.Node = &s

}
func removefirst() {
	head = head.Node

}
func removeafter(std *Student) {
	if std != nil {
		std.Node = std.Node.Node
	}
}
func removestd(std *Student) {
	temp := head
	for temp.Node != std {
		temp = temp.Node
	}
	temp.Node = std.Node

}
