package faviconstring

import (
	"io/ioutil"
	"testing"
)

func TestFrom(t *testing.T) {
	data, err := From("aaaafqqfaqqapppp:255:0:0")
	if err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile("test.ico", data, 0644); err != nil {
		t.Error(err)
	}

}
