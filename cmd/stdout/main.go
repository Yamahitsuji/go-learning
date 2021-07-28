package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("This is stdout.\n")
	fmt.Fprintf(os.Stderr, "This is stderr by fmt.Fprintf.\n")
	os.Stderr.WriteString("This is stderr by os.Stderr.WriteString.\n")
}
