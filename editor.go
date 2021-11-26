package faviconstring

const (
	// Mode "enum"
	modeBlank = iota
	modeGray4 // for 4-bit grayscale images
	modeRGB   // for 8+8+8 bit RGB images
	modeRGBA  // for 8+8+8+8 bit RGBA images
)

// Mode is a per-filetype mode, like for Markdown
type Mode int
