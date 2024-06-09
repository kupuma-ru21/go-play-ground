package main

import "fmt"

func badCall() {
	a, b := 10, 0
	n := a / b // it will cause panic as it's a division by 0
	fmt.Println(n)
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\n", e)
		}
	}()
	badCall()
	fmt.Println("After bad call")
}

func main() {
	fmt.Println("Starting the program")
	test()
	fmt.Println("Ending the program")
}
