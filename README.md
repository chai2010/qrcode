QR code encoder/decoder
=======================

PkgDoc: [http://godoc.org/github.com/chai2010/qrcode](http://godoc.org/github.com/chai2010/qrcode)


Install
=======

Install `GCC` or `MinGW` ([download here](http://tdm-gcc.tdragon.net/download)) at first,
and then run these commands:

1. `go get -d github.com/chai2010/qrcode`
2. `go generate` and `go install`
3. `go run hello.go`

**Notes**

If use `TDM-GCC`, and build error like this:

	c:\tdm-gcc-64\x86_64-w64-mingw32\include\aviriff.h:3:25: error: 'refer' does not
	 name a type
	 * No warranty is given; refer to the file DISCLAIMER within this package.
	 ...

You need fix the `C:\TDM-GCC-64\x86_64-w64-mingw32\include\aviriff.h` header first:

	** // fixit: ** -> /**
	* This file is part of the mingw-w64 runtime package.
	* No warranty is given; refer to the file DISCLAIMER within this package.
	*/

Example
=======

```Go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/qrcode"
)

var (
	flagLevel  = flag.String("n", "Q", "Set QR encode level(L|M|Q|H), default is 'Q'.")
	flagOutput = flag.String("o", "qr.png", "Set output filename (PNG only).")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] text\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(-1)
	}

	filename := filepath.Clean(*flagOutput)
	if filename == "" {
		filename = "qr.png"
	}
	if !strings.HasSuffix(strings.ToLower(filename), ".png") {
		filename += ".png"
	}

	level := qrcode.Q
	switch strings.ToUpper(*flagLevel) {
	case "L":
		level = qrcode.L
	case "M":
		level = qrcode.M
	case "Q":
		level = qrcode.Q
	case "H":
		level = qrcode.H
	}

	code, err := qrcode.Encode(flag.Arg(0), level)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, code.PNG(), 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
```


BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
