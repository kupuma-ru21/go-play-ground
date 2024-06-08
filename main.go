package main

import (
	"fmt"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("output/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}
	defer outputFile.Close()
	for i := 0; i < 10; i++ {
		fmt.Fprintf(outputFile, "Some test data.\n")
	}
}
