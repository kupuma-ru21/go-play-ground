package main

import (
	"fmt"
	"time"
)

type myTime time.Time

func (t myTime) first3Chars() {
	fmt.Println(t)
}

func main() {
	m := myTime(time.Now())
	//calling existing String method on anonymous Time field
	// fmt.Println("Full time now:", m.String())
	m.first3Chars()
}
