package main

import "fmt"

func main() {
	// Parse the commands and files given as input.
	files, N_range, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
	}

	// Call ConvertContent to perform the encoding and overwrite the files.
	if err := ConvertContent(files, N_range); err != nil {
		fmt.Println(err)
	}
}
