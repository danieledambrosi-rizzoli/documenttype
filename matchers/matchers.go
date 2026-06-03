package matchers

import (
	"github.com/danieledambrosi-rizzoli/documenttype/types"
)

type Matcher func([]byte)bool
type TypeMatcher func([]byte)types.Type

type Map map[types.Type]Matcher

var (
	matchers = make(map[types.Type]TypeMatcher)
	// supportedTypes []types.Type
)

func RegisterMatcher(t types.Type, fn Matcher) TypeMatcher {
	if _, exists := matchers[t]; exists {
		return matchers[t]
	}
	matcher := func(magic []byte) types.Type {
		if fn(magic) {
			return t
		}
		return types.Unknown
	}

	matchers[t] = matcher
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
	registerMap()
}