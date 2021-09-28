package main

import (
	"io/ioutil"
	"errors"
)

// ConvertContent goes through each file, extracts it's content
// and ovewrites it back after encoding.
func ConvertContent(files []string, N_range int) error {
	for _, file := range files {
		_, err := ioutil.ReadFile(file)
		if err != nil {
			err := errors.New("Error occured when accesing file")
			return err
		}

		// TODO(Andrei): Call rotn module function.

		message := []byte("Hello")
		if ioutil.WriteFile(file, message, 0644) != nil {
			err := errors.New("Error occured when overwriting file")
			return err
		}
	}
	return nil
}
