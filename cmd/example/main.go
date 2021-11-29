package main

import (
	"io/ioutil"
	"log"

	"github.com/xyproto/faviconstring"
)

func main() {
	data, err := faviconstring.Image("aaaafqqfaqqapppp:255:0:0")
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile("favicon.ico", data, 0644); err != nil {
		log.Fatalln(err)
	}
}
