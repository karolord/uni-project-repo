package main

import "fmt"

var queue Queue      // 1
var slice []int      // 1
var slicecounter int // 1
var x int            // 1

func main() {
	Requirement1()
	Requirement4()
	Requirement5()
	Requirement6()
	Requirement7()
}

type Queue struct {
	value   []int
	counter int
}

func (q *Queue) Enqueue(input int) {
	q.value[q.counter] = input // 3
	q.counter++                // 3
}

func (q *Queue) Dequeue() int {
	tmp := q.value[0]                     // 2
	for i := 0; i < len(q.value)-1; i++ { // 1, 5n + 5, 3n
		q.value[i] = q.value[i+1] // 5n
	}
	q.counter-- // 3
	return tmp  // 1
}

func Requirement1() {
	fmt.Println("Please enter a number:") // 1
	fmt.Scanln(&x)                        // 2
	slice = make([]int, x, x)             // 5
	queue.value = make([]int, x, x)       // 5
	var grade int                         // 1
	for i := 0; i < x; i++ {              // 1, 3n + 3, 3n
		fmt.Println("Enter the grade:") // 1n
		fmt.Scanln(&grade)              // 2n
		for grade < 0 || grade > 100 {  //  5n^2
			fmt.Println("Error invalid input") // 1n^2
			fmt.Println("Enter the grade:")    // 1n^2
			fmt.Scanln(&grade)                 // 2n^2
		}
		Requirement2(grade)
		Requirement3(grade)
	}
}

func Requirement2(grade int) {
	slice[slicecounter] = grade // 2
	slicecounter++              // 3
}

func Requirement3(grade int) {
	queue.Enqueue(grade)
}

func Requirement4() {
	copyslice := slice
	fmt.Println("Requirement 4 (Merge sort):")
	fmt.Println(Mergesort(copyslice))
}

func Requirement5() {
	queue.Insertionsort()
}

func (q *Queue) Insertionsort() {
	copyqueue := q             // 2
	tmpslice := make([]int, x) // 4
	for i := 0; i < x; i++ {   // 1, 3n + 3, 3n
		tmpslice[i] = copyqueue.Dequeue() // 3n
	}
	for i := 0; i < x; i++ { // 1, 3n + 3, 3n
		for j := 0; j < i; j++ { // n, 3n^2 + 3^n, 3n^2
			if tmpslice[i] > tmpslice[j] { // 5n^2
				tmpslice[i], tmpslice[j] = tmpslice[j], tmpslice[i] // 6n^2
			}
		}
	}
	for i := 0; i < x; i++ { // 1, 3n + 3, 3n
		copyqueue.Enqueue(tmpslice[i]) // 2n
	}
	fmt.Println("Requirement 5 (Insertion Sort):") // 1
	fmt.Println(copyqueue.value[0:])               // 2
}

func Merge(left []int, right []int) []int {
	result := make([]int, len(right)+len(left))
	i := 0
	j := 0
	z := 0
	for z < len(result) {
		if i > len(left)-1 && j < len(right) {
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

func Mergesort(input []int) []int {
	if len(input) == 1 {
		return input
	}
	mid := len(input) / 2
	left := input[:mid]
	right := input[mid:]
	return Merge(Mergesort(left), Mergesort(right))
}

func Requirement6() {
	if len(queue.value) < 5 { // 4
		return // 0
	}
	var copyqueue Queue                // 1
	copyqueue.value = make([]int, x-4) // 5
	mid := len(queue.value) / 2        // 4
	if len(queue.value)%2 == 1 {       // 4
		for i := 0; i < x; i++ { // 1, 3n + 3, 3n
			tmp := queue.Dequeue()                                    // 2n
			if i == mid+1 || i == mid+2 || i == mid-1 || i == mid-2 { // 19n
				continue // 0n
			}
			copyqueue.Enqueue(tmp) // 1n
		}
	} else {
		mid--                    // 3
		for i := 0; i < x; i++ { // 1, 3n + 3, 3n
			tmp := queue.Dequeue()                                    // 2n
			if i == mid+2 || i == mid+3 || i == mid-1 || i == mid-2 { // 19n
				continue // 0n
			}
			copyqueue.Enqueue(tmp) // 1n
		}
	}
	fmt.Println("Requirement 6:")    // 1
	fmt.Println(copyqueue.value[0:]) // 2
}

func Requirement7() {
	tmp := make([]int, x+4) // 5
	mid := len(slice) / 2   // 4
	counters := 0           // 1
	if x%2 == 1 {           // 3
		for i := 0; i < x; i++ { // 1, 3n + 3, 3n
			if i == mid+1 || i == mid-1 { // 9n
				tmp[counters] = slice[i] // 4n
				counters++               // 3n
				tmp[counters] = slice[i] // 4n
				counters++               // 3n
			}
			tmp[counters] = slice[i] // 4n
			counters++               // 3n
		}
	} else {
		mid--                    // 3
		for i := 0; i < x; i++ { // 1, 3n + 3, 3n
			if i == mid+2 || i == mid-1 { // 9n
				tmp[counters] = slice[i] // 4n
				counters++               // 3n
				tmp[counters] = slice[i] // 4n
				counters++               // 3n
			}
			tmp[counters] = slice[i] // 4n
			counters++               // 3n
		}
	}
	fmt.Println("Requirement 7:") // 1
	fmt.Println(tmp[0:])          // 2
}
