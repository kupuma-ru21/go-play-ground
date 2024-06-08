package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() {
	filename := p.Title + ".txt"
	err := os.WriteFile(filename, p.Body, 0600)
	if err != nil {
		fmt.Println("Error saving file: ", err)
	}
}

func load(title string) *Page {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error loading file: ", err)
	}
	return &Page{Title: title, Body: body}
}

func main() {
	p1 := &Page{Title: "SamplePage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2 := load(p1.Title)
	fmt.Println(string(p2.Body))
}
