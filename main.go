package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Left  *Node
	Value interface{}
	Right *Node
}

func main() {
	b := []byte(`{"Value": "Father", "Left": {"Value": "Left child"}, "Right": {"Value": "Right child"}}`)
	var f Node
	json.Unmarshal(b, &f)
	fmt.Println(f, f.Left, f.Right)
}
