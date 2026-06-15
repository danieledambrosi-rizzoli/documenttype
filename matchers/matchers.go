package matchers

import (
	"github.com/danieledambrosi-rizzoli/documenttype/types"
)

type Matcher func([]byte)bool
type TypeMatcher func([]byte)types.Type

type Map map[types.Type]Matcher

var (
	Matchers = make(map[types.Type]TypeMatcher)
	// supportedTypes []types.Type
)

func RegisterMatcher(t types.Type, fn Matcher) TypeMatcher {
	if _, exists := Matchers[t]; exists {
		return Matchers[t]
	}
	matcher := func(mime []byte) types.Type {
		if fn(mime) {
			return t
		}
		return types.Unknown
	}

	Matchers[t] = matcher
	// supportedTypes = append([]types.Type{t}, supportedTypes...)
	return matcher
}

func registerMap(matchers ...Map) {
	for _, m := range matchers {
		for t, fn := range m {
			RegisterMatcher(t, fn)
		}
	}
}

func init() {
	registerMap(Documents, PlainTextFiles)
}