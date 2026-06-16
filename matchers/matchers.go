package matchers

import (
	"github.com/danieledambrosi-rizzoli/documenttype/types"
)

type Matcher func([]byte)bool
type TypeMatcher func([]byte)types.Type

type Map map[types.Type]Matcher

var (
	Matchers = make(map[types.Type]TypeMatcher)
	SupportedTypes []TypeMatcher
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
	SupportedTypes = append([]TypeMatcher{matcher}, SupportedTypes...)
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
	// we put PlainTextFiles to the bottom of the queue to optimize
	registerMap(PlainTextFiles, Documents)
}