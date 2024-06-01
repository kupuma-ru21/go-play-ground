package main

import "fmt"

// make struct here
type MytStruct struct {
	namedField float32
	int
	string
}

func main() {
	mystructValue := MytStruct{7.5, 10, "test"}
	fmt.Println(mystructValue)
}
