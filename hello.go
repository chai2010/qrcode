// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// QR codes encoder.
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
