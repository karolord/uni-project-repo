package main
import ("fmt")

func main() {
}
// original in go
type auis struct{
     arr [100000]int
	 n int
}
func (a *auis) Set(n int,x int){
	if n > a.n {
		fmt.Println("Invalid")
		return 		
	}
	a.arr[n] = x
}
func (a *auis) Get(n int) int{
	return a.arr[n]
}

type stack struct{
	a auis
	counter int	
}
func (s *stack) push(x int){

     s.set(s.counter,x)
	s.counter++
}

func (s *stack) pop() int{
	x := s.a.get[s.counter-1]
	s.a.set(s.counter-1,0)
	s.counter--
	return x
}