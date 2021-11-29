package faviconstring

import (
	"io/ioutil"
	"testing"
)

func TestImage1(t *testing.T) {
	data, err := Image("aaaaaaaaffqqqqfffaaqqqqaappppppp#f00")
	if err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}
}

func TestImage2(t *testing.T) {
	data, err := Image("aaaaaaaa ffqqqqfff aaqqqqaa ppppppp #ff0000")
	if err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}
}
