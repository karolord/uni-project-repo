package main

import "fmt"
var slice []int
var queue []int
var counter int
var slicecounter int
var x int
func main() {
	Requirement1()
	Requirement6()

}
func enqueue(x int,queues []int) {
	queues[counter] = x
	counter++ 
}
func dequeue() int{
	value := queue[0]
	queue = queue[1:]
	return value
}

func Requirement1() {
	fmt.Println("Please enter a number:")
	fmt.Scanln(&x)
	slice = make([]int,x,x)
	queue = make([]int,x,x)

	var grade int
	for i := 0; i < x; i++ {
		fmt.Println("Enter the grade:")
		fmt.Scanln(&grade)
		for grade < 0 || grade > 100{
			fmt.Println("Error invalid input")
			fmt.Println("Enter the grade:")
			fmt.Scanln(&grade)
		}
		Requirement2(grade)
		Requirement3(grade)
		
	}
}

func Requirement2(grade int) {
	slice[slicecounter] = grade
	slicecounter++
}

func Requirement3(grade int) {
	enqueue(grade,queue)
}
func Requirement4() {
	copyslice := slice
	fmt.Println(mergesort(copyslice))
}
func merge(left []int, right []int) []int {
	result := make([]int,len(right)+len(left))
	i := 0
	j := 0
	z := 0
	for z < len(result){
		if i > len(left)-1 && j < len(right){
			result[z] = right[j]
			z++
			j++
		} else if i < len(left) && j > len(right)-1 {
			result[z] = left[i]
			z++
			i++
		} else if left[i] > right[j] {
			result[z] = left[i]
			z++
			i++
		} else {
			result[z] = right[j]
			z++
			j++
		}
	}
	return result 
}
func mergesort(input []int) []int  {
	if len(input) == 1{
		return input
	}
	mid := len(input)/2
	left := input[:mid]
	right := input[mid:]
	return merge(mergesort(left), mergesort(right))
}
func insertionsort(input []int){
	
}


func Requirement6(){
	if(len(queue) < 5) {
		return
	}
	counter = 0
	copyqueue := make([]int,x-4)
	mid := len(queue)/2
	if len(queue)%2 == 1 {	
		for i := 0; i <= len(queue); i++ {
			tmp := dequeue()
			if i == mid + 1 || i == mid + 2 || i == mid - 1 || i == mid - 2 {
				continue
			}
			enqueue(tmp,copyqueue)
		}
	} else {
		mid--
		for i := 0; i =< len(queue); i++ {
			tmp := dequeue()
			if i == mid + 2 || i == mid + 3 || i == mid - 1 || i == mid - 2 {
				continue
			}
			enqueue(tmp,copyqueue)
		}
	}
	fmt.Println(copyqueue)
}