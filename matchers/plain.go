/*

	HEURISTIC PLAINTEXT DETECTION
	inspired by https://github.com/profullstack/text-type-detection/blob/master/src/index.js
	
	WHAT SHOULD I PROVIDE
	- isAscii detection
	- isUtf8 detection
	|_check the BOM

	- HTML
	- MD
	- TXT
	- LATEX

*/

package matchers

import(
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
	TypeLatex  = registerType("LATEX", "latex", buildMime("application/x-latex"))
	TypeHtml = registerType("HTML", "html", buildMime("text/html"))
	TypeXhtml = registerType("XHTML", "xhtml", buildMime("application/xhtml+xml"))
	TypeTxt = registerType("TXT", "txt", buildMime("text/plain"))
	TypeMd = registerType("MD", "md", buildMime("text/markdown"))
)

var Plains = Map {
	TypeLatex: Latex,
	TypeHtml:  Html,
	TypeXhtml: Xhtml,
	TypeTxt:   Text,
	TypeMd:    Md,
}

func bytesFromBOM(bom bomEncoding) int {
	switch bom {
	case UTF8: return 3
	case UTF16_BE, UTF16_LE: return 2
	case UTF32_BE, UTF32_LE: return 4
	case UNKNOWN: return 0
	default: return 0
	}
}

// https://github.com/simdutf/simdutf/blob/master/src/encoding_types.cpp
func checkBOM(magic []byte) bomEncoding {
	length := len(magic)
	if length < 3 { return UNKNOWN }
	if length >= 2 && magic[0] == 0xff && magic[1] == 0xfe {
		if length >= 4 && magic[2] == 0x00 && magic[3] == 0x0 {
			return UTF32_LE
		} else {
			return UTF16_LE
		}
	} else if length >= 2 && magic[0] == 0xfe && magic[1] == 0xff {
		return UTF16_BE
	} else if length >= 4 && magic[0] == 0x00 && magic[1] == 0x00 &&
				magic[2] == 0xfe && magic[3] == 0xff {
		return UTF32_BE
	} else if length >= 3 && magic[0] == 0xef && magic[1] == 0xbb &&
				magic[2] == 0xbf {
		return UTF8
	}
	return UNKNOWN
}

// if it isnt a valid utf8 we dont bother parsing the file
// if utf8 => ascii
// if ascii ? utf8
func IsUtf8(magic []byte) bool {
	
	return false
}

func Latex(magic []byte) bool {
	return false
}

func Text(magic []byte) bool {
	return false
}

func Html(magic []byte) bool {
	return false
}

func Xhtml(magic []byte) bool {
	return false
}

func Md(magic []byte) bool {
	return false
}