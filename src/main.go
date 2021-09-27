package main

import "fmt"

func main() {
	// Parse the commands and files given as input.
	_, _, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
	}
}
