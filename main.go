package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc1 := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc1)
	fmt.Printf("JSON format: %s", js)
	vc2 := VCard{}
	err := json.Unmarshal(js, &vc2)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Struct format: %v", vc2)
}
