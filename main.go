package main

import (
	"fmt"
	"os"
)

func main() {
	inputFile := "products.txt"
	buf, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
}
