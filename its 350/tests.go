package main

func main() {
	x(5)
	x("yes")
}

func x(s int) {
	s = 5
}
func x(s string) {
	s = "yes"
}
