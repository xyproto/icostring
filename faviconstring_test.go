package faviconstring

import (
	"io/ioutil"
	"testing"
)

func TestImage(t *testing.T) {
	data, err := Image("aaaafqqfaqqapppp:255:0:0")
	if err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}

}
