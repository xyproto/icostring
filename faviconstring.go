package faviconstring

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"strconv"
	"strings"

	ico "github.com/biessek/golang-ico"
)

var (
	// 4-bit, 16-color grayscale grading by runes
	lookupRunes = map[rune]byte{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
		'i': 8,
		'j': 9,
		'k': 10,
		'l': 11,
		'm': 12,
		'n': 15,
		'o': 13,
		'p': 14,
	}
)

// Create converts the textual representation to an .ico image, using 16 letters 'a'..'p'.
// The string can be followed by a colon spearated color, like :255:0:0 for red, or :0:0:255 for blue.
// The letter 't' is transparent.
func From(s string) ([]byte, error) {
	var (
		// Create a new image
		width  = 16
		height = 16
		m      = image.NewRGBA(image.Rect(0, 0, width, height))

		// These are used in the loops below
		x, y      int
		line      string
		intensity byte
		ru        rune
		runes     []rune

		// Default 'q' color.
		r byte = 255
		g byte = 0
		b byte = 0
	)

	if len(s) < 16 {
		// only 4x4 grayscale images are supported!
		return []byte{}, errors.New("16 letters representing a 4x4 grayscale image ('a'..'p') are expected")
	}
	if colonCount := strings.Count(s, ":"); colonCount > 0 && colonCount != 3 {
		return []byte{}, errors.New("16 letters followed by (optionally): a colon, a byte, a colon, a byte, a colon and a byte was expected")
	}

	if strings.Contains(s, ":") {
		parts := strings.SplitN(s, ":", 4)
		s = parts[0]
		rString := parts[1]
		gString := parts[2]
		bString := parts[3]
		if rInt, err := strconv.Atoi(rString); err != nil {
			return []byte{}, err
		} else {
			r = byte(rInt)
		}
		if gInt, err := strconv.Atoi(gString); err != nil {
			return []byte{}, err
		} else {
			g = byte(gInt)
		}
		if bInt, err := strconv.Atoi(bString); err != nil {
			return []byte{}, err
		} else {
			b = byte(bInt)
		}
	}

	// Create an intermediate representation

	text := ""
	for i, ru := range s {
		if i > 0 && i%4 == 0 {
			text += fmt.Sprintf("%s\n%s\n%s\n%s\n", line, line, line, line)
			line = ""
		}
		line += fmt.Sprintf("%c%c%c%c", ru, ru, ru, ru)
	}
	text += fmt.Sprintf("%s\n%s\n%s\n%s", line, line, line, line)
	line = ""

	// 	if verbose {
	// 		fmt.Println(text)
	// 	}

	// Draw the pixels
	for y, line = range strings.Split(text, "\n") {
		if y >= 16 { // max 16x16 pixels
			break
		}
		runes = []rune(line)
		if len(runes) < 16 {
			continue
		}
		for x = 0; x < 16; x++ { // max 16x16 pixels
			ru = runes[x]
			if ru == 't' { // transparent
				m.Set(x, y, color.RGBA{0, 0, 0, 0})
			} else if ru == 'q' { // color
				m.Set(x, y, color.RGBA{r, g, b, 0xff})
			} else {
				intensity = lookupRunes[ru]*16 + 15 // from 0..15 to 15..255
				// Draw pixel to image
				m.Set(x, y, color.RGBA{intensity, intensity, intensity, 0xff})
			}
		}
	}

	var buf bytes.Buffer
	if err := ico.Encode(&buf, m); err != nil {
		return []byte{}, nil
	}
	return buf.Bytes(), nil
}