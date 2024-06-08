package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("output/target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(err)
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()
	io.Copy(dst, src)
}
