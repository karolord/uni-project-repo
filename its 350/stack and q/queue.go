package main

var arr [100]int

func inqueue(x int) {

	for i := 0; i < 100; i++ {
		if arr[i] == 0 {
			arr[i] = x
		}
	}
}

func dequeue() int {
	y := arr[0]
	for i := 0; i < 99; i++ {
		arr[i] = arr[i+1]
	}
	return y
}

func main() {

}
