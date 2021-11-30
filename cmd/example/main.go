package main

import (
	"log"
	"os"

	"github.com/xyproto/icostring"
)

func WriteFile(filename, imageString string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := icostring.WriteImage(f, imageString); err != nil {
		return err
	}
	return f.Close()
}

func main() {
	if err := WriteFile("favicon.ico", "aaaafqqfaqqapppp"); err != nil {
		log.Fatalln(err)
	}
}
