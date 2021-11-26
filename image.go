package faviconstring

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"strings"
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

// WriteFavicon converts the textual representation to an .ico image
// If asOther is true, .png images are written as .ico and the other way around
// TODO: Generate a string instead
func WriteFavicon(mode Mode, text, filename string, asOther bool) error {
	if mode != modeGray4 {
		return errors.New("saving .ico files is only implemented for 4-bit grayscale images")
	}

	var (
		// Create a new image
		width  = 16
		height = 16
		m      = image.NewRGBA(image.Rect(0, 0, width, height))

		// These are used in the loops below
		x, y      int
		line      string
		intensity byte
		r         rune
		runes     []rune
	)

	// Draw the pixels
	for y, line = range strings.Split(text, "\n") {
		if y >= 16 { // max 16x16 pixels
			break
		}
		runes = []rune(line)
		for x = 0; x < 16; x++ { // max 16x16 pixels
			if (x * 2) < len(runes) {
				r = runes[x*2]
				if r == 'T' { // transparent
					// Draw a black transparent pixel
					m.Set(x, y, color.RGBA{0, 0, 0, 0})
				} else {
					intensity = lookupRunes[r]*16 + 15 // from 0..15 to 15..255
					// Draw pixel to image
					m.Set(x, y, color.RGBA{intensity, intensity, intensity, 0xff})
				}
			} else {
				// Draw a white transparent pixel
				m.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 0})
			}
		}
	}

	if asOther && strings.HasSuffix(filename, ".ico") {
		filename = strings.Replace(filename, ".ico", ".png", 1)
		// Create a new file
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		// Encode the image as a .png image
		return png.Encode(f, m)
	} else if !asOther && strings.HasSuffix(filename, ".png") {
		// Create a new file
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		return png.Encode(f, m)
	} else if asOther && strings.HasSuffix(filename, ".png") {
		filename = strings.Replace(filename, ".png", ".ico", 1)
	}

	// Create a new file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Encode the image as an .ico image
	//return ico.Encode(f, m)
	return EncodeGrayscale4bit(f, m) // Sadly, this does not seem to support transparency
}

// This is from github.com/biessek/golang-ico, only to be able to use private structs
type head struct {
	Zero   uint16
	Type   uint16
	Number uint16
}

// This is from github.com/biessek/golang-ico, only to be able to use private structs
type direntry struct {
	Width   byte
	Height  byte
	Palette byte
	_       byte
	Plane   uint16
	Bits    uint16
	Size    uint32
	Offset  uint32
}

// EncodeGrayscale4bit is a modified version of the function from github.com/biessek/golang-ico, only to be able to save 4-bit .ico images
func EncodeGrayscale4bit(w io.Writer, im image.Image) error {
	b := im.Bounds()
	m := image.NewGray(b)
	draw.Draw(m, b, im, b.Min, draw.Src)
	header := head{
		0,
		1,
		1,
	}
	entry := direntry{
		Plane:  1,
		Bits:   4, // was: 32
		Offset: 22,
	}
	pngbuffer := new(bytes.Buffer)
	pngwriter := bufio.NewWriter(pngbuffer)
	err := png.Encode(pngwriter, m)
	if err != nil {
		return err
	}
	err = pngwriter.Flush()
	if err != nil {
		return err
	}
	entry.Size = uint32(len(pngbuffer.Bytes()))
	bounds := m.Bounds()
	entry.Width = uint8(bounds.Dx())
	entry.Height = uint8(bounds.Dy())
	bb := new(bytes.Buffer)
	var e error
	if e = binary.Write(bb, binary.LittleEndian, header); e != nil {
		return e
	}
	if e = binary.Write(bb, binary.LittleEndian, entry); e != nil {
		return e
	}
	if _, e = w.Write(bb.Bytes()); e != nil {
		return e
	}
	_, e = w.Write(pngbuffer.Bytes())
	return e
}
