package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file")
		os.Exit(1)
	}

	fileName := os.Args[1]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("File Contents: \n%s\n", content)
}
