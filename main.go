package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to open file")

	}

	str := "write this file by Golang!"
	data := []byte(str)
	count, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("fail to write file")
	}

	fmt.Printf("write %d bytes\n", count)
}
