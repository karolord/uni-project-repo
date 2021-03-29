package main
import("fmt")

// homework create print similar to this but print name of the forth student
// pointer tail
type Student struct {
	Name string
	GPA float64
	ID int
	Node *Student
}
var head *Student
func add(n string,g float64, i int) {
	s:=Student{n,g,i,nil}
	if head == nil {
		head = &s
		}else{
			c := head
			for c != nil{
				c = c.Node
			}
			c.Node = &s
		}
	}
	func main() {
		add("ali",3.2,155555 )
		add("karo",3.22,55555 )
		fmt.Println(head)
		fmt.Println(head.Node)
	
		
	}