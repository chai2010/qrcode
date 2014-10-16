// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qrcode

import (
	"image"
)

func Decode(m image.Image) (text string, err error) {
	p := qr_decoder_open()
	defer qr_decoder_close(p)

	rgb := newRGBIplImage(m)
	defer rgb.Release()

	qr_decoder_decode_image(p, rgb, _DEFAULT_ADAPTIVE_TH_SIZE, _DEFAULT_ADAPTIVE_TH_DELTA)
	text = cgo_qr_decoder_get_text(p)
	return
}
