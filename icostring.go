package icostring

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"io"
	"strings"

	ico "github.com/biessek/golang-ico"
	"github.com/g4s8/hexcolor"
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
		'n': 13,
		'o': 14,
		'p': 15,
	}
	errTooShort = errors.New("16 or 64 letters representing a 4x4 grayscale image ('a'..'p') are expected")
	errHash     = errors.New("16 or 64 letters followed by (optionally): a hex color like #f00")
)

// WriteImage converts the textual representation to an .ico image, using 16 letters 'a'..'p'.
// The string can be followed by a hex color, like #f00 for red, #ff0000 for red or #00f for blue,
// which defines a custom color. The custom color can be used with 'q'. The letter 't' is transparency.
func WriteImage(w io.Writer, s string) error {
	var (
		// Create a new image
		width  = 16
		height = 16
		m      = image.NewRGBA(image.Rect(0, 0, width, height))

		// These are used in the loops below
		x, y      int
		line      string
		intensity byte
		runes     []rune

		// Default 'q' color.
		r byte = 255
		g byte
		b byte
		a byte = 255
	)

	// only 4x4 grayscale images are supported

	s = strings.ReplaceAll(s, " ", "")

	if strings.Contains(s, "#") {
		parts := strings.SplitN(s, "#", 2)
		s = parts[0]
		customColor, err := hexcolor.Parse("#" + parts[1])
		if err != nil {
			return err
		}
		r = byte(customColor.R)
		g = byte(customColor.G)
		b = byte(customColor.B)
		a = byte(customColor.A)
	}

	text := ""
	switch len(s) {
	case 64:
		// Create an intermediate representation
		for i, ru := range s {
			if i > 0 && i%8 == 0 { // 8 characters per row, before scaling up
				text += fmt.Sprintf("%s\n%s\n", line, line)
				line = ""
			}
			line += fmt.Sprintf("%c%c", ru, ru)
		}
		text += fmt.Sprintf("%s\n%s", line, line)
	case 16:
		// Create an intermediate representation
		for i, ru := range s {
			if i > 0 && i%4 == 0 { // 4 characters per row, before scaling up
				text += fmt.Sprintf("%s\n%s\n%s\n%s\n", line, line, line, line)
				line = ""
			}
			line += fmt.Sprintf("%c%c%c%c", ru, ru, ru, ru)
		}
		text += fmt.Sprintf("%s\n%s\n%s\n%s", line, line, line, line)
	default:
		return errTooShort
	}

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
			switch runes[x] {
			case 't': // transparent
				m.Set(x, y, color.RGBA{0, 0, 0, 0})
			case 'q': // color
				m.Set(x, y, color.RGBA{r, g, b, a})
			default:
				intensity = lookupRunes[runes[x]]*16 + 15 // from 0..15 to 15..255
				// Draw pixel to image
				m.Set(x, y, color.RGBA{intensity, intensity, intensity, 0xff})
			}
		}
	}

	// Encode and write the image to the supplied io.Writer
	return ico.Encode(w, m)
}

// Image converts the textual representation to an .ico image, using 16 letters 'a'..'p'.
// The string can be followed by a hex color, like #f00 for red, #ff0000 for red or #00f for blue,
// which defines a custom color. The custom color can be used with 'q'. The letter 't' is transparency.
func Image(s string) ([]byte, error) {
	var buf bytes.Buffer
	if err := WriteImage(&buf, s); err != nil {
		return []byte{}, nil
	}
	return buf.Bytes(), nil
}
