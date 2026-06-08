package documenttype

import(
	"github.com/danieledambrosi-rizzoli/documenttype/types"
	"github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

var Matchers = matchers.Matchers

func GetType(magic []byte) types.Type {
	for _, matcher := range Matchers {
		var t = matcher(magic)
		if  t != types.Unknown {
			return t
		}
	}
	return types.Unknown
}

func IsType(magic []byte, t types.Type) bool {
	return Matchers[t](magic) != types.Unknown
}