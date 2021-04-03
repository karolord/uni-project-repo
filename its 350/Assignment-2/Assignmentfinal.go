package main

import "fmt"

var queue Queue
var slice []int
var slicecounter int
var x int

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
	q.value[q.counter] = input
	q.counter++
}

func (q *Queue) Dequeue() int {
	tmp := q.value[0]
	for i := 0; i < len(q.value)-1; i++ {
		q.value[i] = q.value[i+1]
	}
	q.counter--
	return tmp
}

func Requirement1() {
	fmt.Println("Please enter a number:")
	fmt.Scanln(&x)
	slice = make([]int, x, x)
	queue.value = make([]int, x, x)
	var grade int
	for i := 0; i < x; i++ {
		fmt.Println("Enter the grade:")
		fmt.Scanln(&grade)
		for grade < 0 || grade > 100 {
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

func (q *Queue) Insertionsort(){
	copyqueue := q
	tmpslice := make([]int, x)
	for i := 0; i < x; i++ {
		tmpslice[i] = copyqueue.Dequeue()
	}
	for i := 0; i < x; i++ {
		for j := 0; j < i; j++ {
			if tmpslice[i] > tmpslice[j] {
				tmpslice[i], tmpslice[j] = tmpslice[j], tmpslice[i]
			}
		}
	}
	for i := 0; i < x; i++ {
		copyqueue.Enqueue(tmpslice[i])
	}
	fmt.Println("Requirement 5 (Insertion Sort):")
	fmt.Println(copyqueue.value[0:])
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
	if len(queue.value) < 5 {
		return
	}
	var copyqueue Queue
	copyqueue.value = make([]int, x-4)
	mid := len(queue.value) / 2
	if len(queue.value)%2 == 1 {
		for i := 0; i < x; i++ {
			tmp := queue.Dequeue()
			if i == mid+1 || i == mid+2 || i == mid-1 || i == mid-2 {
				continue
			}
			copyqueue.Enqueue(tmp)
		}
	} else {
		mid--
		for i := 0; i < x; i++ {
			tmp := queue.Dequeue()
			if i == mid+2 || i == mid+3 || i == mid-1 || i == mid-2 {
				continue
			}
			copyqueue.Enqueue(tmp)
		}
	}
	fmt.Println("Requirement 6:")
	fmt.Println(copyqueue.value[0:])
}

func Requirement7() {
	tmp := make([]int, x+4)
	mid := len(slice) / 2
	counters := 0
	if x%2 == 1 {
		for i := 0; i < x; i++ {
			if i == mid+1 || i == mid-1 {
				tmp[counters] = slice[i]
				counters++
				tmp[counters] = slice[i]
				counters++
			}
			tmp[counters] = slice[i]
			counters++
		}
	} else {
		mid--
		for i := 0; i < x; i++ {
			if i == mid+2 || i == mid-1 {

				tmp[counters] = slice[i]
				counters++
				tmp[counters] = slice[i]
				counters++
			}
			tmp[counters] = slice[i]
			counters++
		}
	}
	fmt.Println("Requirement 7:")
	fmt.Println(tmp[0:])
}
