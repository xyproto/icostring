package main

import (
	"io/ioutil"
	"log"

	"github.com/xyproto/icostring"
)

func main() {
	data, err := icostring.Image("aaaaaaaa aaaqqaaa ffqqqqff ffqqqqff aaqqqqaa aaqqqqaa pppqqppp pppppppp #08f")
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile("favicon.ico", data, 0644); err != nil {
		log.Fatalln(err)
	}
}
