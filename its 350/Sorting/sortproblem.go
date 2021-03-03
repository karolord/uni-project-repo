package main

import (
	"fmt"
)

var arr []int

func main() {
	fmt.Println("yes")
}

func bubblesort() {
	var swapped bool = true
	for i := 1; i < len(arr) && !swapped; i++ {
		swapped = false
		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			swapped = true
		}
	}

}
