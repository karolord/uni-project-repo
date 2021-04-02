package main

import "fmt"
var slice []int
var queue []int
var counter int
var x int
func main() {
	Requirement1()
	fmt.Println(mergeSort(slice))
}
func enqueue(x int){
	slice = append(slice,x)
}
func dequeue() int{
	value := slice[0]
	slice = slice[1:]
	return value
}

func Requirement1() {
	fmt.Println("Please enter a number:")
	fmt.Scanln(&x)
	slice = make([]int,x,x)
	fmt.Println(len(slice))
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
		
		
	}
}

func Requirement2(grade int) {
	slice[counter] = grade
}

func Requirement3(grade int) {
	enqueue(grade)
}
}
func mergeSort(items []int) []int {
    var num = len(items)
     
    if num == 1 {
        return items
    }
     
    middle := int(num / 2)
    var (
        left = make([]int, middle)
        right = make([]int, num-middle)
    )
    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
     
    return merge(mergeSort(left), mergeSort(right))
}
 
func merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
     
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
     
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
     
    return
}