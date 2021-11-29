package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xyproto/faviconstring"
)

func writeImage(imageString, filename string) error {
	data, err := faviconstring.Image(imageString)
	if err != nil {
		return err
	}
	if filename == "-" {
		_, err = os.Stdout.Write(data)
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func main() {
	var err error
	argCount := len(os.Args) - 1
	switch argCount {
	case 2:
		imageString := os.Args[1]
		filename := os.Args[2]
		err = writeImage(imageString, filename)
	case 1:
		imageString := os.Args[1]
		err = writeImage(imageString, "-")
	default:
		fmt.Fprintln(os.Stderr, "Syntax: genico IMAGESTRING FILENAME")
		fmt.Fprintln(os.Stderr, "Example: genico aaaafqqfaqqapppp favicon.ico")
		fmt.Fprintln(os.Stderr, "Example: genico aaaafqqfaqqapppp#08f favicon.ico")
		fmt.Fprintln(os.Stderr, "a-p is grayscale, q is the custom suffix color and t is transparent. Spaces are ignored.")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
