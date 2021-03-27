package main

import (
	"fmt"
)

var counter int

func main() {
	// AUISDS.go
	var arr [1000000000]int
	var N int

	// Get a value from our custom DS
	fmt.Printf("Enter how long you want your array to be: ")
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("Enter an integer: ")
		fmt.Scanln(&arr[i])
		Set(i, arr[i], arr, N)
	}

	for i := 0; i < N; i++ {
		fmt.Print(Get(i, arr), " ")
	}

}

func Get(ix int, arr [1000000000]int) int {
	return arr[ix]
}

func Set(ix int, value int, arr [1000000000]int, N int) {
	if ix >= N {
		fmt.Println("NOT ALLOWED")
		return
	}
	push(value, arr, N)
}

func push(value int, arr [1000000000]int, N int) {
	for i := 0; i < N; i++ {
		if arr[i] == 0 {
			arr[i] = value
		}
	}
}

func pop(value int, arr [1000000000]int, N int) {
	for i := N; i => 0; i++ {
		if arr[i] != 0 {
			arr[i] = value
		}
	}
}