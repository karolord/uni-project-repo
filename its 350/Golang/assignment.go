package main

import "fmt"

var arr [10]int

func main() {
	add(50, 4)
	fmt.Println(arr[4])
	addbegin(30)
	fmt.Println(arr[0])
	delete(0)
	fmt.Println(arr[0])
}
func addbegin(value int) {
	add(value, 0)
}

func add(value int, index int) {
	for i := len(arr) - 1; i >= index; i-- {
		if i == index {
			arr[i] = value
		} else {
			arr[i] = arr[i-1]
		}
	}
}

func delete(index int) {
	for i := index; i < len(arr); i++ {
		if i == len(arr)-1 {
			arr[i] = 0
		} else {
			arr[i] = arr[i+1]
		}

	}
}
