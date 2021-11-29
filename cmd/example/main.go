package main

import (
	"io/ioutil"
	"log"

	"github.com/xyproto/faviconstring"
)

func main() {
	data, err := faviconstring.Image("aaaaaaaa aaaqqaaa ffqqqqff ffqqqqff aaqqqqaa aaqqqqaa pppqqppp pppppppp #08f")
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile("favicon.ico", data, 0644); err != nil {
		log.Fatalln(err)
	}
}
