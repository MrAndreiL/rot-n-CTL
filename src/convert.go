package main

import (
	"io/ioutil"
	"errors"

	"github.com/MrAndreiL/rot-n-CTL/tree/master/src/rotn"
)

// ConvertContent goes through each file, extracts it's content
// and ovewrites it back after encoding.
func ConvertContent(files []string, N_range int) error {
	for _, file := range files {
		buffer, err := ioutil.ReadFile(file)
		if err != nil {
			err := errors.New("Error occured when accesing file")
			return err
		}

		// Call rotn encoding and write the result in the given file.
		encoding := rotn.NRotation(buffer, N_range)
		if ioutil.WriteFile(file, encoding, 0644) != nil {
			err := errors.New("Error occured when overwriting file")
			return err
		}
	}
	return nil
}
