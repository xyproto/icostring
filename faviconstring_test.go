package faviconstring

import (
	"io/ioutil"
	"testing"
)

func TestImage1(t *testing.T) {
	data, err := Image("aaaa fqqf aqqa pppp #ff0000")
	if err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}
}

func TestImage2(t *testing.T) {
	data, err := Image("aaaafqqfaqqapppp")
	if err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}
}

func TestImage3(t *testing.T) {
	data, err := Image("aaaaaaaa aaaaaaaa ffqqqqff ffqqqqff aaqqqqaa aaqqqqaa pppppppp pppppppp #f00")
	if err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}
}

func TestImage4(t *testing.T) {
	data, err := Image("aaaaaaaaaaaaaaaaffqqqqffffqqqqffaaqqqqaaaaqqqqaapppppppppppppppp")
	if err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}
}
