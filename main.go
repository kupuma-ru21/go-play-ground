package main

import (
	"fmt"
	"go-play-ground/even"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
