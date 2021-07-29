package main

import (
	"fmt"
	"io"
	"os"
)

var w io.Writer

func init() {
	w = os.Stdout
}

func main() {
	run()
}

func run() {
	x, y := 3, 7
	fmt.Fprintf(w, "%d + %d = %d\n", x, y, x+y)
}
