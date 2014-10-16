// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"log"

	"github.com/chai2010/gopkg/image"
	_ "github.com/chai2010/gopkg/image/png"
	"github.com/chai2010/gopkg/image/qrcode/internal/libqrencode"
)

func main() {
	text := "http://godoc.org/github.com/chai2010/gopkg/image/qrcode/libqrencode"
	code, err := libqrencode.Encode(1, libqrencode.Q, libqrencode.EightBit, text)
	if err != nil {
		log.Fatal(err)
	}
	err = image.Save("qrgen.png", code, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}
