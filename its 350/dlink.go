package main

// homework create print similar to this but print name of the forth student
// pointer tail
type Student struct {
	Name string
	GPA  float64
	ID   int
	Node *Student
	prev *Student
}

var head *Student

func main() {

}

func addafter(n string, g float64, i int, std *Student) {
	s := Student{n, g, i, nil, nil}
	s.Node = std.Node
	s.prev = std
	std.Node = &s
	s.Node.prev = &s

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
	std.prev.Node = std.Node
	std.Node.prev = std.prev
}
