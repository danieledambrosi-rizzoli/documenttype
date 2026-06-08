package documenttype

import(
	"os"
	"testing"
)

func TestDocx(t *testing.T) {
	cases := []struct{
		path string
		ext  string
		match  bool
	}{
		{"samples/sample.docx", "docx", true},
		{"samples/sample2.docx", "docx", true},
		{"samples/sample.odt", "docx", false},
	}

	for _, c := range cases {
		f, err := os.Open(c.path)
		if err != nil {
			t.Fatalf("The file doesn't exist")
		}
		defer f.Close()

		buffer := make([]byte, 6192)
		f.Read(buffer)
		typ := GetType(buffer)

		res := typ.Ext == c.ext

		if res != c.match {
			t.Fatalf("Invalid match %s", c.path)
		}
	}
}