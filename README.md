# faviconstring

Convert a string of 16 characters, 'a'..'p'  to a grayscale ICO file.

* 'a' is darkest, 'p' is brightest.
* 'q' is also allowed, to set a custom color (provided with an r, g and b byte).
* 't' is transparent.

Includes both a Go package and a command line utility.

### Example use

```go
package main

import (
    "io/ioutil"
    "log"

    "github.com/xyproto/faviconstring"
)

func main() {
    data, err := faviconstring.From("aaaafqqfaqqapppp:255:0:0")
    if err != nil {
        log.Fatalln(err)
    }
    if err := ioutil.WriteFile("favicon.ico", data, 0644); err != nil {
        log.Fatalln(err)
    }
}
```

### General info

* Version: 0.0.1
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
