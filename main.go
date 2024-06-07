package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to open file")

	}

	data := make([]byte, 1024)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to read file")
	}

	fmt.Printf("read %d bytes:\n", count)
	fmt.Println(string(data[:count]))
}
