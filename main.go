package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter some input: ")
	// scanner.Split(bufio.ScanRunes) // <-- scan runes (newline included)
	// scanner.Split(bufio.ScanWords) // <-- scan words (sequence of characters delimited by spaces)
	scanner.Split(bufio.ScanLines) // <-- the default: scan lines
	for scanner.Scan() {           // The for loop stops when a EOF is given
		fmt.Println(scanner.Text())
	}
}
