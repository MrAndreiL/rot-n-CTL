package main

import (
	"os"
	"errors"
)

// ParseArgs goes over the commands given as input.
// If no mistakes are found, returns the list of files that need to be encoded
// and N, which is the encoding's range.
// If mistakes are found, an error message is sent.
func ParseArgs() ([]string, int, error) {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "-H" {
		err := errors.New(usage())
		return nil, 0, err
	}

	N, err := encodingRange(args[0])
	if err != nil {
		return nil, 0, err
	}

	files := []string{}
	for i := 1; i < len(args); i++ {
		if _, err := os.Stat(args[i]); err != nil { // if the file does not exist.
			return nil, 0, err
		}
		files = append(files, args[i])
	}
	return files, N, nil
}

// usage returns a string representing the correct usage.
func usage() string {
	usage := `rotn.go [h] n=N <filename.txt> [<filename.txt>...]
	-h -> print correct usage
	n=N -> set the ROT encoding range, default is zero`
	return usage
}

// encodingRange breaks -n=N into runes and builds N as int.
func encodingRange(arg string) (int, error) {
	if arg[0] != '-' || arg[1] != 'n' || arg[2] != '=' {
		err := errors.New("Invalid command")
		return 0, err
	}

	r := []rune(arg)
	N := 0
	isPositive := 1
	if arg[3] == '-'{ // is positive or negative range.
		isPositive = -1;
	}

	for i := 3; i < len(r); i++ {
		if !('0' <= arg[i] && arg[i] <= '9') {
			err := errors.New("N is not a valid integer")
			return 0, err
		}
		N = N * 10 + (int(r[i]) - '0')
	}
	return (N * isPositive), nil
}
