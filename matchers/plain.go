/*

	HEURISTIC PLAINTEXT DETECTION
	inspired by https://github.com/profullstack/text-type-detection/blob/master/src/index.js

	WHAT SHOULD I PROVIDE
	- isAscii detection
	- isUtf8 detection
	|_check the BOM

	// maybe later
	- HTML
	- MD
	- TXT
	- LATEX

*/

package matchers

import (
	// "github.com/danieledambrosi-rizzoli/documenttype/types"
	// "github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

type bomEncoding int

const (
	UNKNOWN bomEncoding = iota
	UTF8
	UTF16_LE
	UTF16_BE
	UTF32_LE
	UTF32_BE
)

var (
	TypeLatex = registerType("LATEX", "latex", buildMime("application/x-latex"))
	TypeHtml  = registerType("HTML", "html", buildMime("text/html"))
	TypeXhtml = registerType("XHTML", "xhtml", buildMime("application/xhtml+xml"))
	TypeTxt   = registerType("TXT", "txt", buildMime("text/plain"))
	TypeMd    = registerType("MD", "md", buildMime("text/markdown"))
)

var PlainTextFiles = Map{
	TypeLatex: Latex,
	TypeHtml:  Html,
	TypeXhtml: Xhtml,
	TypeTxt:   Text,
	TypeMd:    Md,
}

func bytesFromBOM(bom bomEncoding) int {
	switch bom {
	case UTF8:
		return 3
	case UTF16_BE, UTF16_LE:
		return 2
	case UTF32_BE, UTF32_LE:
		return 4
	case UNKNOWN:
		return 0
	default:
		return 0
	}
}

// https://github.com/simdutf/simdutf/blob/master/src/encoding_types.cpp
func checkBOM(mime []byte) bomEncoding {
	length := len(mime)
	if length < 2 {
		return UNKNOWN
	}
	if length >= 2 && mime[0] == 0xff && mime[1] == 0xfe {
		if length >= 4 && mime[2] == 0x00 && mime[3] == 0x0 {
			return UTF32_LE
		} else {
			return UTF16_LE
		}
	} else if length >= 2 && mime[0] == 0xfe && mime[1] == 0xff {
		return UTF16_BE
	} else if length >= 4 && mime[0] == 0x00 && mime[1] == 0x00 &&
		mime[2] == 0xfe && mime[3] == 0xff {
		return UTF32_BE
	} else if length >= 3 && mime[0] == 0xef && mime[1] == 0xbb &&
		mime[2] == 0xbf {
		return UTF8
	}
	return UNKNOWN
}

func IsAscii(buf []byte) bool {
	var idx uint
	for idx+8 <= uint(len(buf)) {
		data := leBytes(buf[idx : idx+8 : idx+8])
		if data&m80 != 0 {
			return false
		}
		idx += 8
	}
	var data uint64
	if len(buf) >= 8 {
		shft := 64 - uint(len(buf))%8*8
		data = leBytes(buf[len(buf)-8:len(buf):len(buf)]) >> shft
	} else {
		var shft uint
		for idx < uint(len(buf)) {
			data |= uint64(buf[idx]) << shft
			shft += 8
			idx++
		}
	}
	return data&m80==0
}

// implementation from https://github.com/sugawarayuuta/charcoal
func IsUtf8(buf []byte) bool {
	s64 := state64{xe0: m80, xed: m80, xf0: m80}
	var idx uint
	for idx+8 <= uint(len(buf)) {
		data := leBytes(buf[idx : idx+8 : idx+8])
		if data&m80 != 0 {
			break
		}
		idx += 8
	}
	for idx+8 <= uint(len(buf)) {
		data := leBytes(buf[idx : idx+8 : idx+8])
		if !s64.add(data) {
			return false
		}
		// in theory if data is only ascii, we could go back to fast checking...
		idx += 8
	}
	var data uint64
	if len(buf) >= 8 {
		shft := 64 - uint(len(buf))%8*8
		data = leBytes(buf[len(buf)-8:len(buf):len(buf)]) >> shft
	} else {
		var shft uint
		for idx < uint(len(buf)) {
			data |= uint64(buf[idx]) << shft
			shft += 8
			idx++
		}
	}
	return (data&m80 == 0 || s64.add(data)) && s64.top&m80 == 0
}

func Latex(mime []byte) bool {
	return false
}

func Text(mime []byte) bool {
	return checkBOM(mime) != UNKNOWN || IsUtf8(mime)
}

func Html(mime []byte) bool {
	return false
}

func Xhtml(mime []byte) bool {
	return false
}

func Md(mime []byte) bool {
	return false
}