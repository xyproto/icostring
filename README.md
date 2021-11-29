# faviconstring

Convert a string of 16 characters, 'a'..'p'  to a grayscale ICO file.

* 'a' is darkest, 'p' is brightest.
* 'q' is also allowed, to set a custom color (provided with an r, g and b byte).
* 't' is transparent.

Includes both a Go package and a command line utility.

### String format

To make the string shorter, it represents a 4x4 image that will be scaled up to 16x16 when it is converted to an ICO.

This string `aaaaaaaaaaaaaaaa` will be converted to an image where every pixel is dark.

For a completely blue image, use the custom color `q`, where the red, green and blue colors are appended at the end: `qqqqqqqqqqqqqqqq:0:0:255`.

The default custom color is red, so this is a completely red image: `qqqqqqqqqqqqqqqq`.

For a dark gray square surrounded by transparent pixels, this can be used: `ttttttaattttaatttttt`.

For a yellow square surrounded by transparent pixels, this can be used: `ttttttqqttttqqtttttt:255:255:0`.

For a tiny piece of art, try: `aaaafqqfaqqapppp:255:128:80`.

### Example use

```go
package main

import (
    "io/ioutil"
    "log"

    "github.com/xyproto/faviconstring"
)

func main() {
    data, err := faviconstring.From("aaaafqqfaqqapppp:255:128:80")
    if err != nil {
        log.Fatalln(err)
    }
    if err := ioutil.WriteFile("favicon.ico", data, 0644); err != nil {
        log.Fatalln(err)
    }
}
```

### Example images

| aaaa aaaa aaaa aaaa           | qqqq qqqq qqqq qqqq               |
| ----------------------------- | --------------------------------- |
| ![aaaa](img/aaaaaaaaaaaa.ico) | ![qqqq](img/qqqqqqqqqqqqqqqq.ico) |

### General info

* Version: 0.0.1
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
