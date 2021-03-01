package main

var arr [100]int

func push(x int) {

	for i := 0; i < 100; i++ {
		if arr[i] == 0 {
			arr[i] = x
		}
	}
}

func pop() int {
	for i := 99; i > -1; i-- {
		if arr[i] != 0 {
			x := arr[i]
			arr[i] = 0
			return x
		}
	}
}

func main() {

}
