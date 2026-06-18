package documenttype

import(
	"github.com/danieledambrosi-rizzoli/documenttype/types"
	"github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

var supportedTypes = matchers.SupportedTypes
var Matchers = matchers.Matchers

var TypeDoc  = matchers.TypeDoc
var TypeDocx = matchers.TypeDocx
var TypeOdt  = matchers.TypeOdt
var TypePdf  = matchers.TypePdf
var TypeTxt  = matchers.TypeTxt

func GetType(buf []byte) types.Type {
	for _, matcher := range supportedTypes {
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