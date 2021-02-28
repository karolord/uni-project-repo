package main

import (
	"fmt"
)

var arr [10]int // wll cost 1 time t
var counter int

func main() {

	insertL(10)
	insertL(20)
	insertL(30)

}

func insertL(a int) {
	arr[counter] = a
	counter++
}

func printY() {
	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])

	}

}

func delete(a int) {

	for i := a; i < 10; i++ {
		arr[i] = arr[i+1]
	}

}

func add(a int, b int) {
	for i := 10; i > a; i++ {
		arr[i] = arr[i-1]

	}
}
