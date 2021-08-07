package main

import "fmt"

func main() {
	fmt.Println("=========== test ===========")
	return

	// unreachable error (govet)
	fmt.Println("-------")
}

// unused error
func f() {
	fmt.Println("function f")
}
