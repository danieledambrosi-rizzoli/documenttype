package documenttype

import(
	"github.com/danieledambrosi-rizzoli/documenttype/types"
	"github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

var SupportedTypes = matchers.SupportedTypes
var Matchers = matchers.Matchers

func GetType(buf []byte) types.Type {
	for _, matcher := range SupportedTypes {
		var t = matcher(buf)
		if  t != types.Unknown {
			return t
		}
	}
	return types.Unknown
}

func IsType(buf []byte, t types.Type) bool {
	return Matchers[t](buf) != types.Unknown
}