package main

import(
	"fmt"
	"os"

	"github.com/danieledambrosi-rizzoli/documenttype/types"
	"github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

func checktype(filename string, t types.Type) (bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	buffer := make([]byte, 8192)
	file.Read(buffer)


	return matchers.Documents[t](buffer), nil
}

func main() {
	typeDocx := matchers.TypeDocx
	typeOdt := matchers.TypeOdt

	if ok, _ := checktype("samples/sample.docx", typeDocx); ok {
		fmt.Println("sample.docx is a Docx")
		fmt.Println(typeDocx.Ext, typeDocx.Mime.Value())
	} else {
		fmt.Println("sample.docx isn't a Docx")
	}

	if ok, _ := checktype("samples/sample2.docx", typeDocx); ok {
		fmt.Println("sample2.docx is a Docx")
		fmt.Println(typeDocx.Ext, typeDocx.Mime.Value())
	} else {
		fmt.Println("sample2.docx isn't a Docx")
	}

	if ok, _ := checktype("samples/sample.odt", typeOdt); ok {
		fmt.Println("sample.odt is a Odt")
		fmt.Println(typeDocx.Ext, typeDocx.Mime.Value())
	} else {
		fmt.Println("sample.odt isn't a Odt")
	}

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
