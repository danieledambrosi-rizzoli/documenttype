package documenttype

import(
	"testing"
	"errors"
	"fmt"

	"bufio"
	"sync"
	"os"
	"unicode/utf8"

	"github.com/danieledambrosi-rizzoli/documenttype/matchers"
)

type Tests struct {
	buf []byte
	res bool
}

func testUtf8(tests []Tests) (bool, error) {
	for _, t := range tests {
		if r := matchers.IsUtf8(t.buf); r != t.res {
			errstr := fmt.Sprintf("%s got evaluated as %t, expected %t", t.buf, r, t.res)
			return false, errors.New(errstr)
			/* EXPECTED FORMAT */
			// --- FAIL: TestIsUtf8 (0.00s)
			// plain_test.go:51: ☺☻☹ got evaluated as true, expected false
			// FAIL
		}
	}

	return true, nil
}

func TestIsUtf8(t *testing.T) {
	tests := []Tests{
		{[]byte(""), true},
		{[]byte("a"), true},
		{[]byte("abc"), true},
		{[]byte("Ж"), true},
		{[]byte("ЖЖ"), true},
		{[]byte("брэд-ЛГТМ"), true},
		{[]byte("☺☻☹"), true},
		{[]byte("aa\xe2"), false},
		{[]byte{66, 250}, false},
		{[]byte{66, 250, 67}, false},
		{[]byte("a\uFFFDb"), true},
		{[]byte("\xF4\x8F\xBF\xBF"), true},      // U+10FFFF
		{[]byte("\xF4\x90\x80\x80"), false},     // U+10FFFF+1; out of range
		{[]byte("\xF7\xBF\xBF\xBF"), false},     // 0x1FFFFF; out of range
		{[]byte("\xFB\xBF\xBF\xBF\xBF"), false}, // 0x3FFFFFF; out of range
		{[]byte("\xc0\x80"), false},             // U+0000 encoded in two bytes: incorrect
		{[]byte("\xed\xa0\x80"), false},         // U+D800 high surrogate (sic)
		{[]byte("\xed\xbf\xbf"), false},         // U+DFFF low surrogate (sic)
	}

	ok, err := testUtf8(tests)
	if err != nil || !ok {
		t.Fatal(err.Error())
	}
}

func testCasesBuffered(cases []Cases) (bool, error) {
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 6192)
		},
	}

	buf := bufferPool.Get().([]byte)[:0]
	for _, c := range cases {
		f, err := os.Open(c.path)
		if err != nil {
			return false, errors.New("The file doesn't exist")
		}
		defer f.Close()


		br := bufio.NewReader(f)
		n, err := br.Read(buf[:cap(buf)])
		if err != nil {
			return false, errors.New("Could not read the file")
		}

		typ := GetType(buf[:n])
		res := typ.Ext == c.ext
		if res != c.match {
			return false, errors.New("Invalid match " + c.path)
		}
	}

	return true, nil
}

func TestText(t *testing.T) {
	cases := []Cases{
		{"samples/darwin.txt", "txt", true},
		{"samples/divina_commedia.txt", "txt", true},
		{"samples/frankestein.txt", "txt", true},
		{"samples/symposium.txt", "txt", true},
	}

	ok, err := testCasesBuffered(cases)
	if err != nil || !ok {
		t.Fatal(err.Error())
	}
}

func BenchmarkTextRecognition(b *testing.B) {
	cases := []Cases{
		{"samples/darwin.txt", "txt", true},
		{"samples/divina_commedia.txt", "txt", true},
		{"samples/frankestein.txt", "txt", true},
		{"samples/symposium.txt", "txt", true},
	}

	b.Run("mine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ok, err := testCasesBuffered(cases)
			if err != nil || !ok {
				b.Fatal(err.Error())
			}
		}
	})
	b.Run("std", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range cases {
				buffer, err := os.ReadFile(c.path)
				if err != nil {
					b.Fatal(err.Error())
				}

				res := utf8.Valid(buffer)

				if res != c.match {
					b.Fatal("Invalid match " + c.path)
				}
			}
		}
	})
}