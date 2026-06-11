package documenttype

import(
	"os"
	"testing"
	"errors"
)

type Cases struct {
	path string
	ext  string		
	match  bool
}

func testCases(cases []Cases) (bool, error) {
	for _, c := range cases {
		f, err := os.Open(c.path)
		if err != nil {
			return false, errors.New("The file doesn't exist")
		}
		defer f.Close()

		buffer := make([]byte, 6192)
		f.Read(buffer)
		typ := GetType(buffer)

		res := typ.Ext == c.ext

		if res != c.match {
			return false, errors.New("Invalid match " + c.path)
		}
	}

	return true, nil
}

func TestDocx(t *testing.T) {
	cases := []Cases {
		{"samples/sample.docx", "docx", true},
		{"samples/sample2.docx", "docx", true},
		{"samples/sample.odt", "docx", false},
	}

	ok, err := testCases(cases)
	if err != nil || !ok {
		t.Fatal(err.Error())
	}
}

func TestOdt(t *testing.T) {
	cases := []Cases {
		{"samples/sample.docx", "odt", false},
		{"samples/sample2.docx", "odt", false},
		{"samples/sample.odt", "odt", true},
	}

	ok, err := testCases(cases)
	if err != nil || !ok {
		t.Fatal(err.Error())
	}
}

func TestPdf(t *testing.T) {
	cases := []Cases {
		{"samples/sample.pdf", "pdf", true},
		{"samples/sample2.pdf", "pdf", true},
		{"samples/sample.odt", "pdf", false},
	}

	ok, err := testCases(cases)
	if err != nil || !ok {
		t.Fatal(err.Error())
	}
}