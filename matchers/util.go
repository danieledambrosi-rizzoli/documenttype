package matchers

import (
	"bytes"
	"github.com/danieledambrosi-rizzoli/documenttype/types"
)

var zipMagic = [...]byte{'P', 'K', 0x03, 0x04}
var registerType = types.RegisterType
var buildMime = types.BuildMIME

func bytescmp(haystack, needle []byte, offset int) bool {
	if offset < 0 { return false }
	if haystack == nil || needle == nil { return false }
	if len(needle) + offset > len(haystack) { return false }

	return bytes.Equal(needle, haystack[offset:offset+len(needle)])
}

// implementation from https://github.com/sugawarayuuta/charcoal
func leBytes(buf []byte) uint64 {
	_ = buf[7] // bounds check hint to compiler; see golang.org/issue/14808
	return uint64(buf[0]) | uint64(buf[1])<<8 |
		uint64(buf[2])<<16 | uint64(buf[3])<<24 |
		uint64(buf[4])<<32 | uint64(buf[5])<<40 |
		uint64(buf[6])<<48 | uint64(buf[7])<<56
}