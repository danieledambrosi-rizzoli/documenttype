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