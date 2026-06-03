package main

import(
	"fmt"

	"github.com/danieledambrosi-rizzoli/documenttype/types"
	"github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

func main() {
	typePng := types.RegisterType("PNG", "png", types.BuildMIME("image/png"))
	matchPng := matchers.RegisterMatcher(
		typePng,
		func(magic []byte) bool {
			return len(magic) > 3 &&
			magic[0] == 0x89 && magic[1] == 0x50 &&
			magic[2] == 0x4E && magic[3] == 0x47
		},
	)
	var pngMagic = []byte{
    	0x89, 0x50, 0x4E, 0x47,
    	0x0D, 0x0A, 0x1A, 0x0A,
	}

	if (matchPng(pngMagic) == typePng) {
		fmt.Println("This is a PNG")
		fmt.Println(typePng.Ext, typePng.Mime.Value())
	} else {
		fmt.Println("This isn't a PNG")
	}
}