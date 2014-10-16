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
// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// QR codes encoder.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/qrcode"
)

var (
	flagText   = flag.String("text", "chaishushan@gmail.com", "Set text")
	flagLevel  = flag.String("n", "Q", "Set QR encode level(L|M|Q|H), default is 'Q'.")
	flagOutput = flag.String("o", "qrcode.png", "Set output filename (PNG only).")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] text\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

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

	code, err := qrcode.Encode(*flagText, level)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, code.PNG(), 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Save as:", filename)

	// Load file data
	data, err := ioutil.ReadFile("testdata/01-1.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// Decode image
	m, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	// Decode QR Code
	text, err := qrcode.Decode(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("testdata/01-1.jpg:", text)
}
```


BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
