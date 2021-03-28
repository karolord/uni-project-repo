package main
import ("fmt")

func main() {
	var a auis
	a.Set(5,4)
	fmt.Println()

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
     s.a.set(s.counter,x)
	s.counter++
}

func (s *stack) pop() int{
	s.counter--
	x := s.a.get[s.counter]
	s.a.set(s.counter,0)
	return x
}