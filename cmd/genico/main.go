package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xyproto/faviconstring"
)

func writeImage(imageString, filename string) error {
	data, err := faviconstring.From(imageString)
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
	case 5:
		imageString := os.Args[1] + ":" + os.Args[2] + ":" + os.Args[3] + ":" + os.Args[4]
		filename := os.Args[5]
		err = writeImage(imageString, filename)
	case 4:
		imageString := os.Args[1] + ":" + os.Args[2] + ":" + os.Args[3] + ":" + os.Args[4]
		err = writeImage(imageString, "-")
	case 2:
		imageString := os.Args[1]
		filename := os.Args[2]
		err = writeImage(imageString, filename)
	case 1:
		imageString := os.Args[1]
		err = writeImage(imageString, "-")
	default:
		fmt.Fprintln(os.Stderr, "syntax: genico IMAGESTRING RED GREEN BLUE FILENAME")
		fmt.Fprintln(os.Stderr, "example: genico aaaafqqfaqqapppp 255 0 0 test.ico")
		fmt.Fprintln(os.Stderr, "a-p is grayscale, q is the color and t is transparent")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}