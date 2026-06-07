package matchers

import (
	"encoding/binary"
	"unsafe"

	"github.com/danieledambrosi-rizzoli/documenttype/types"
)

var registerType = types.RegisterType
var buildMime = types.BuildMIME

var (
	TypeDoc  = registerType("DOC", "doc", buildMime("application/msword"))
	TypeDocx = registerType("DOCX", "docx", buildMime("application/vnd.openxmlformats-officedocument.wordprocessingml.document"))
	TypeOdt  = registerType("ODT", "odt", buildMime("application/vnd.oasis.opendocument.text"))
)

var Documents = Map{
	TypeDoc:  Doc,
	TypeDocx: Docx,
	TypeOdt:  Odt,
}

// this method is the exact copy of https://github.com/h2non/filetype/blob/master/matchers/document.go
func Doc(magic []byte) bool {
	if len(magic) > 513 {
		return magic[0] == 0xD0 && magic[1] == 0xCF &&
			magic[2] == 0x11 && magic[3] == 0xE0 &&
			magic[512] == 0xEC && magic[513] == 0xA5
	} else {
		return len(magic) > 3 &&
			magic[0] == 0xD0 && magic[1] == 0xCF &&
			magic[2] == 0x11 && magic[3] == 0xE0
	}
}

// i did this myself :)
func Docx(magic []byte) bool {
	const HEAD_SIZE = 0x1E
	walk := func(headerStart *int) {
		var head = []byte(magic)
		if *headerStart < 0 || *headerStart + HEAD_SIZE > len(head) {
			*headerStart = -1
			return
		}

		if !bytescmp(head, zipMagic[:], *headerStart) {
			*headerStart = -1
			return
		}

		*headerStart += HEAD_SIZE +
			int(binary.LittleEndian.Uint16(head[*headerStart+26:*headerStart+28])) +
			int(binary.LittleEndian.Uint16(head[*headerStart+28:*headerStart+30])) +
			int(binary.LittleEndian.Uint32(head[*headerStart+18:*headerStart+22]))
	}

	if !bytescmp(magic, zipMagic[:], 0) { return false }

	if bytescmp(magic, []byte("word/"), HEAD_SIZE) { return true }

	if !bytescmp(magic, []byte("[Content_Types].xml"), HEAD_SIZE) &&
		!bytescmp(magic, []byte("_rels/.rels"), HEAD_SIZE) &&
		!bytescmp(magic, []byte("docProps"), HEAD_SIZE) &&
		!bytescmp(magic, []byte("_rels"), HEAD_SIZE) {
		return false
	}

	headStart := 0
	walk(&headStart)
	if headStart < 0 { return false }
	walk(&headStart)
	if headStart < 0 { return false }

	if bytescmp(magic, []byte("word/"), headStart+HEAD_SIZE) { return true }

	walk(&headStart)
	if headStart < 0 { return false }
	if bytescmp(magic, []byte("word/"), headStart+HEAD_SIZE) { return true }

	return false
}

func Odt(magic []byte) bool {
	return checkOdf(magic, TypeOdt.Mime.Value())
}

// this method is ALMOST the exact copy of https://github.com/h2non/filetype/blob/master/matchers/document.go
func checkOdf(magic []byte, mimetype string) bool {
	if 38+len(mimetype) >= len(magic) {
		return false
	}

	if !bytescmp(magic, zipMagic[:], 0) { return false }

	if magic[8] != 0 || magic[9] != 0 { return false }

	if magic[26] != 8 || magic[27] != 0 { return false }

	if int(magic[18]) != len(mimetype) ||
		magic[19] != 0 || magic[20] != 0 || magic[21] != 0 ||
		int(magic[22]) != len(mimetype) ||
		magic[23] != 0 || magic[24] != 0 || magic[25] != 0 {
		return false
	}

	if magic[28] != 0 || magic[29] != 0 { return false }

	// Optimised the code a little to avoid string conversion and allocation
	return bytescmp(magic, []byte("mimetype"), 30) &&
		bytescmp(magic, unsafe.Slice(unsafe.StringData(mimetype), len(mimetype)), 38)
}
