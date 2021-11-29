# faviconstring

Convert a string of 16 characters (`'a'..'p'`) to a grayscale ICO file, with one optional color.

Includes both a Go package and a command line utility.

### Example images

| aaaaaaaaaaaaaaaa                  | qqqqqqqqqqqqqqqq                  | tttttaattaattttt                             | aaaafqqfaqqapppp:80:128:255                 |
| --------------------------------- | --------------------------------- | -------------------------------------------- | ------------------------------------------- |
| ![dark](img/aaaaaaaaaaaaaaaa.ico) | ![red](img/qqqqqqqqqqqqqqqq.ico)  | ![transparent](img/tttttaattaattttt.ico)     | ![art](img/aaaafqqfaqqapppp:80:128:255.ico) |

### Quick installation

For Go 1.17 or later:

    go install github.com/xyproto/faviconstring/cmd/genico@latest

### Example use

```go
package main

import (
    "io/ioutil"
    "log"

    "github.com/xyproto/faviconstring"
)

func main() {
    data, err := faviconstring.Image("aaaafqqfaqqapppp:80:128:255")
    if err != nil {
        log.Fatalln(err)
    }
    if err := ioutil.WriteFile("favicon.ico", data, 0644); err != nil {
        log.Fatalln(err)
    }
}
```

### String format

To make the string shorter, it represents a 4x4 image that will be scaled up to 16x16 when it is converted to an ICO.

* The first four letters is the top row, the four next letters is the second row etc.
* `a` is the darkest grayscale color, `b` is a bit lighter, `c` is a bit lighter than that etc.
* `p` is the lightest grayscale color.
* `q` is a custom color that is either red, or defined at the end of the string with three bytes separated by `:`, like this: `:255:255:255`.
* `t` is transparent.

#### Example image strings

* A dark image: `aaaaaaaaaaaaaaaa` (`a` is the darkest grayscale color)
* A blue image: `qqqqqqqqqqqqqqqq:0:0:255`. (`q` is the custom color that is optionally defined at the end of the string).
* A red image: `qqqqqqqqqqqqqqqq` (the default custom color is red, `:255:0:0`)
* A dark gray square surrounded by transparent pixels: `ttttttaattttaatttttt`.
* A yellow square surrounded by transparent pixels: `ttttttqqttttqqtttttt:255:255:0`.
* A tiny piece of art: `aaaafqqfaqqapppp:80:128:255`.

### General info

* Version: 0.0.1
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
