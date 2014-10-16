// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package qrcode

/*
#cgo CFLAGS : -I./internal/libdecodeqr/libdecodeqr -I./internal/libdecodeqr/opencv110/cxcore/include -I./internal/libdecodeqr/opencv110/cv/include -I./internal/libdecodeqr/opencv110/cvaux/include  -I./internal/libdecodeqr/opencv110/ml/include -I./internal/libdecodeqr/opencv110/highgui/include -DCV_NO_BACKWARD_COMPATIBILITY
#cgo linux CXXFLAGS: -Wunused -DHAVE_CAMV4L2 -DHAVE_CAMV4L
#cgo windows LDFLAGS: -L. -ldecodeqr
#cgo linux LDFLAGS: -L. -ldecodeqr -lm -ldl -lstdc++ -lgstapp-0.10

#cgo linux pkg-config: gtk+-2.0 gstreamer-0.10 libxine libavdevice libavformat libavfilter libavcodec libswscale libavutil

#include "decodeqr.h"
*/
import "C"

// ----------------------------------------------------------------------------
// cv.h
// ----------------------------------------------------------------------------

type (
	_CvPoint  C.CvPoint
	_CvBox2D  C.CvBox2D
	_IplImage C.IplImage
)

// ----------------------------------------------------------------------------
// qrerror.h
// ----------------------------------------------------------------------------

const (
	_QR_IMAGEREADER_WORKING = int(C.QR_IMAGEREADER_WORKING) // 0x1000
	_QR_IMAGEREADER_DECODED = int(C.QR_IMAGEREADER_DECODED) // 0x2000

	_QR_VERSIONINFO_ERROR         = int(C.QR_VERSIONINFO_ERROR)         // 0x0f00
	_QR_VERSIONINFO_INVALID       = int(C.QR_VERSIONINFO_INVALID)       // 0x0100
	_QR_VERSIONINFO_MISMATCH      = int(C.QR_VERSIONINFO_MISMATCH)      // 0x0200
	_QR_VERSIONINFO_UNRECOVERABLE = int(C.QR_VERSIONINFO_UNRECOVERABLE) // 0x0800

	_QR_FORMATINFO_ERROR         = int(C.QR_FORMATINFO_ERROR)         // 0x00f0
	_QR_FORMATINFO_INVALID_LEVEL = int(C.QR_FORMATINFO_INVALID_LEVEL) // 0x0010
	_QR_FORMATINFO_UNRECOVERABLE = int(C.QR_FORMATINFO_UNRECOVERABLE) // 0x0080

	_QR_CODEDATA_ERROR           = int(C.QR_CODEDATA_ERROR)           // 0x000f
	_QR_CODEDATA_NOT_SUPPORT_ECI = int(C.QR_CODEDATA_NOT_SUPPORT_ECI) // 0x0001
	_QR_CODEDATA_LENGTH_MISMATCH = int(C.QR_CODEDATA_LENGTH_MISMATCH) // 0x0002
	_QR_CODEDATA_UNRECOVERABLE   = int(C.QR_CODEDATA_UNRECOVERABLE)   // 0x0008

	_QR_IMAGEREADER_ERROR                      = int(C.QR_IMAGEREADER_ERROR)                      // 0x4000
	_QR_IMAGEREADER_NOT_INVALID_SRC_IMAGE      = int(C.QR_IMAGEREADER_NOT_INVALID_SRC_IMAGE)      // 0x0100
	_QR_IMAGEREADER_NOT_FOUND_FINDER_PATTERN   = int(C.QR_IMAGEREADER_NOT_FOUND_FINDER_PATTERN)   // 0x0200
	_QR_IMAGEREADER_NOT_FOUND_CODE_AREA        = int(C.QR_IMAGEREADER_NOT_FOUND_CODE_AREA)        // 0x0400
	_QR_IMAGEREADER_NOT_DETERMINABLE_CODE_AREA = int(C.QR_IMAGEREADER_NOT_DETERMINABLE_CODE_AREA) // 0x0800
)

// ----------------------------------------------------------------------------
// qrtypes.h
// ----------------------------------------------------------------------------

const (
	_QR_LEVEL_M = int(C.QR_LEVEL_M) // 0
	_QR_LEVEL_L = int(C.QR_LEVEL_L) // 1
	_QR_LEVEL_H = int(C.QR_LEVEL_H) // 2
	_QR_LEVEL_Q = int(C.QR_LEVEL_Q) // 3
)

const (
	_QR_MODE_NUMBER   = int(C.QR_MODE_NUMBER)   // 1
	_QR_MODE_ALPHABET = int(C.QR_MODE_ALPHABET) // 2
	_QR_MODE_JOINT    = int(C.QR_MODE_JOINT)    // 3
	_QR_MODE_8BIT     = int(C.QR_MODE_8BIT)     // 4
	_QR_MODE_FNC1_1   = int(C.QR_MODE_FNC1_1)   // 5
	_QR_MODE_ECI      = int(C.QR_MODE_ECI)      // 7
	_QR_MODE_KANJI    = int(C.QR_MODE_KANJI)    // 8
	_QR_MODE_FNC1_2   = int(C.QR_MODE_FNC1_2)   // 9
)

type (
	_QrDecoderHandle C.QrDecoderHandle
	_QrCodeHeader    C.QrCodeHeader
)

// same as C.QrCodeHeader
type qrCodeHeader struct {
	Model         int32
	Version       int32
	Level         int32
	CharactorSize int32
	ByteSize      int32
}

// ----------------------------------------------------------------------------
// decodeqr.h
// ----------------------------------------------------------------------------

const (
	_DEFAULT_ADAPTIVE_TH_SIZE  = int(C.DEFAULT_ADAPTIVE_TH_SIZE)  // 25
	_DEFAULT_ADAPTIVE_TH_DELTA = int(C.DEFAULT_ADAPTIVE_TH_DELTA) // 10
)

/////////////////////////////////////////////////////////////////////////
//
// initializer
//
// ARGS: none
// RETURN:
//   QrDecoderHandle handle
//
func qr_decoder_open() _QrDecoderHandle {
	p := C.qr_decoder_open()
	return _QrDecoderHandle(p)
}

/////////////////////////////////////////////////////////////////////////
//
// initializer with source image size
//
// ARGS:
//   int width:  pixel width of source image
//   int height: pixel height of source image
//   int depth:  image depth (bit par pixel; use OpenCV IPL_DEPTH_*)
//   int channel: number of image channel
//
// RETURN:
//   QrDecoderHandle handle
//
// NOTE:
//   24-bit full color image has IPL_DEPTH_8U depth and 3 channels.
//
func qr_decoder_open_with_image_size(width, height, depth, channel int) _QrDecoderHandle {
	p := C.qr_decoder_open_with_image_size(
		C.int(width), C.int(height), C.int(depth), C.int(channel),
	)
	return _QrDecoderHandle(p)
}

/////////////////////////////////////////////////////////////////////////
//
// finalizer
//
// ARGS:
//   QrDecoderHandle decoder: handler
//
// RETURN: none
//
func qr_decoder_close(p _QrDecoderHandle) {
	C.qr_decoder_close(C.QrDecoderHandle(p))
}

/////////////////////////////////////////////////////////////////////////
//
// get status
//
// ARGS:
//   QrDecoderHandle decoder: handler
//
// RETURN: status code
//
func qr_decoder_get_status(p _QrDecoderHandle) int {
	v := C.qr_decoder_get_status(C.QrDecoderHandle(p))
	return int(v)
}

/////////////////////////////////////////////////////////////////////////
//
// get working status
//
// ARGS:
//   QrDecoderHandle decoder: handler
//
// RETURN: status code
//
func qr_decoder_is_busy(p _QrDecoderHandle) int {
	v := C.qr_decoder_is_busy(C.QrDecoderHandle(p))
	return int(v)
}

/////////////////////////////////////////////////////////////////////////
//
// set source image size
//
// ARGS:
//   QrDecoderHandle decoder: handler
//   int width:  pixel width of source image
//   int height: pixel height of source image
//   int depth:  image depth (bit par pixel; use OpenCV IPL_DEPTH_*)
//   int channel: number of image channel
//
// RETURN:
//   QrDecoderHandle handle
//
// NOTE:
//   This method provide same function as qr_decoder_open_with_image_size().
//
func qr_decoder_set_image_size(p _QrDecoderHandle, width, height, depth, channel int) _QrDecoderHandle {
	v := C.qr_decoder_set_image_size(C.QrDecoderHandle(p),
		C.int(width), C.int(height), C.int(depth), C.int(channel),
	)
	return _QrDecoderHandle(v)
}

/////////////////////////////////////////////////////////////////////////
//
// preset gaven image as source image
//
// ARGS:
//   QrDecoderHandle decoder: handler
//   IplImage *src: source image
//
// RETURN:
//   QrDecoderHandle handle
//
func qr_decoder_set_image_buffer(p _QrDecoderHandle, src *_IplImage) _QrDecoderHandle {
	v := C.qr_decoder_set_image_buffer(C.QrDecoderHandle(p), (*C.IplImage)(src))
	return _QrDecoderHandle(v)
}

/////////////////////////////////////////////////////////////////////////
//
// get source image buffer
//
// ARGS:
//   QrDecoderHandle decoder: handler
//
// RETURN:
//   IplImage *: pointer to buffer source image|NULL
//
// NOTE:
//   See OpenCV reference manual to access to IplImage *
//
func qr_decoder_get_image_buffer(p _QrDecoderHandle) *_IplImage {
	v := C.qr_decoder_get_image_buffer(C.QrDecoderHandle(p))
	return (*_IplImage)(v)
}

func qr_decoder_get_transformed_image_buffer(p _QrDecoderHandle) *_IplImage {
	v := C.qr_decoder_get_transformed_image_buffer(C.QrDecoderHandle(p))
	return (*_IplImage)(v)
}
func qr_decoder_get_binarized_image_buffer(p _QrDecoderHandle) *_IplImage {
	v := C.qr_decoder_get_binarized_image_buffer(C.QrDecoderHandle(p))
	return (*_IplImage)(v)
}
func qr_decoder_get_tmp_image_buffer(p _QrDecoderHandle) *_IplImage {
	v := C.qr_decoder_get_tmp_image_buffer(C.QrDecoderHandle(p))
	return (*_IplImage)(v)
}

/////////////////////////////////////////////////////////////////////////
//
// decode preset source image
//
// ARGS:
//   QrDecoderHandle decoder: handler
//   int adaptive_th_size: value of AdaptiveThreshold size
//   int adaptive_th_delta: value of AdaptiveThreshold delta
//
// RETURN:
//   short: status code of decoder
//
// NOTE:
//   On succeeded, status code has 0x2000.
//   See qrtypes.h for details of status code.
//
//   In case of adaptive_th_size=0, binarizing methods will be
//   used cvThreshlod() instead of cvAdaptiveThreshold()
//
func qr_decoder_decode(p _QrDecoderHandle, adaptive_th_size, adaptive_th_delta int) int {
	v := C.qr_decoder_decode(C.QrDecoderHandle(p), C.int(adaptive_th_size), C.int(adaptive_th_delta))
	return int(v)
}

/////////////////////////////////////////////////////////////////////////
//
// decode gaven image
//
// ARGS:
//   QrDecoderHandle decoder: handler
//   IplImage *src: image to decode
//   int adaptive_th_size: value of AdaptiveThreshold size
//   int adaptive_th_delta: value of AdaptiveThreshold delta
//
// RETURN:
//   short: status code of decoder
//
func qr_decoder_decode_image(p _QrDecoderHandle, src *_IplImage, adaptive_th_size, adaptive_th_delta int) int {
	v := C.qr_decoder_decode_image(C.QrDecoderHandle(p), (*C.IplImage)(src), C.int(adaptive_th_size), C.int(adaptive_th_delta))
	return int(v)
}

/////////////////////////////////////////////////////////////////////////
//
// get abstruction of decoded data
//
// ARGS:
//   QrDecoderHandle decoder: handler
//   QrCodeHeader *header: pointer to buffer of header
//
// RETURN:
//   1 (on success)||0 (on error)
//
func qr_decoder_get_header(p _QrDecoderHandle, header *_QrCodeHeader) int {
	v := C.qr_decoder_get_header(C.QrDecoderHandle(p), (*C.QrCodeHeader)(header))
	return int(v)
}

/////////////////////////////////////////////////////////////////////////
//
// get decoded text data
//
// ARGS:
//   QrDecoderHandle decoder: handler
//   unsigned char *buf: pointer to buffer of header
//   int buf_size: buffer size
//
// RETURN:
//   copied data size||0 (on error)
//
// NOTE:
//   The data DOES NOT TERMINATE with null.
//   To get actual buffer size, use QrCodeHeader's .byte_size element.
//
func qr_decoder_get_body(p _QrDecoderHandle, buf []byte) int {
	cBuf := cgoSafePtr(buf)
	defer cgoFreePtr(cBuf)
	v := C.qr_decoder_get_body(C.QrDecoderHandle(p), (*C.uchar)(cBuf), C.int(len(buf)))
	return int(v)
}

/////////////////////////////////////////////////////////////////////////
//
// get vertexes of decoded code region
//
// ARGS:
//   QrDecoderHandle decoder: handler
//
// RETURN:
//   Pointer to CvPoint[4] which consist vertexes of code region
//
func qr_decoder_get_coderegion_vertexes(p _QrDecoderHandle) *_CvPoint {
	v := C.qr_decoder_get_coderegion_vertexes(C.QrDecoderHandle(p))
	return (*_CvPoint)(v)
}

/////////////////////////////////////////////////////////////////////////
//
// get Box array of decoded finder patterns
//
// ARGS:
//   QrDecoderHandle decoder: handler
//
// RETURN:
//   Pointer to CvBox2D[3] which consist boxes of finder pattern
//
func qr_decoder_get_finderpattern_boxes(p _QrDecoderHandle) *_CvBox2D {
	v := C.qr_decoder_get_finderpattern_boxes(C.QrDecoderHandle(p))
	return (*_CvBox2D)(v)
}

/////////////////////////////////////////////////////////////////////////
//
// version information
//
func qr_decoder_version() string             { return C.GoString(C.qr_decoder_version()) }
func qr_decoder_version_description() string { return C.GoString(C.qr_decoder_version_description()) }
func qr_decoder_version_product() string     { return C.GoString(C.qr_decoder_version_product()) }
func qr_decoder_version_major() int          { return int(C.qr_decoder_version_major()) }
func qr_decoder_version_minor() int          { return int(C.qr_decoder_version_minor()) }
func qr_decoder_version_teeny() int          { return int(C.qr_decoder_version_teeny()) }
func qr_decoder_version_suffix() string      { return C.GoString(C.qr_decoder_version_suffix()) }
func qr_decoder_version_revision() string    { return C.GoString(C.qr_decoder_version_revision()) }

// ----------------------------------------------------------------------------
// END
// ----------------------------------------------------------------------------
