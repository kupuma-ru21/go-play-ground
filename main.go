package main

import (
	"encoding/gob"
	"log"
	"os"
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
	var vc VCard
	file, _ := os.OpenFile("output/vcard.gob", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err := decoder.Decode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	log.Println(vc)
}
