package matchers

import (
	"encoding/binary"
	"unsafe"
)

var (
	TypeDoc  = registerType("DOC", "doc", buildMime("application/msword"))
	TypeDocx = registerType("DOCX", "docx", buildMime("application/vnd.openxmlformats-officedocument.wordprocessingml.document"))
	TypeOdt  = registerType("ODT", "odt", buildMime("application/vnd.oasis.opendocument.text"))
	TypePdf  = registerType("PDF", "pdf", buildMime("application/pdf"))
)

var Documents = Map{
	TypeDoc:  Doc,
	TypeDocx: Docx,
	TypeOdt:  Odt,
	TypePdf:  Pdf,
}

// this method is the exact copy of https://github.com/h2non/filetype/blob/master/matchers/document.go
func Doc(mime []byte) bool {
	if len(mime) > 513 {
		return mime[0] == 0xD0 && mime[1] == 0xCF &&
			mime[2] == 0x11 && mime[3] == 0xE0 &&
			mime[512] == 0xEC && mime[513] == 0xA5
	} else {
		return len(mime) > 3 &&
			mime[0] == 0xD0 && mime[1] == 0xCF &&
			mime[2] == 0x11 && mime[3] == 0xE0
	}
}

// i did this myself :)
func Docx(mime []byte) bool {
	const HEAD_SIZE = 0x1E
	walk := func(headerStart *int) {
		var head = []byte(mime)
		if *headerStart < 0 || *headerStart+HEAD_SIZE > len(head) {
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

	if !bytescmp(mime, zipMagic[:], 0) {
		return false
	}

	if bytescmp(mime, []byte("word/"), HEAD_SIZE) {
		return true
	}

	if !bytescmp(mime, []byte("[Content_Types].xml"), HEAD_SIZE) &&
		!bytescmp(mime, []byte("_rels/.rels"), HEAD_SIZE) &&
		!bytescmp(mime, []byte("docProps"), HEAD_SIZE) &&
		!bytescmp(mime, []byte("_rels"), HEAD_SIZE) {
		return false
	}

	headStart := 0
	walk(&headStart)
	if headStart < 0 {
		return false
	}
	walk(&headStart)
	if headStart < 0 {
		return false
	}

	if bytescmp(mime, []byte("word/"), headStart+HEAD_SIZE) {
		return true
	}

	walk(&headStart)
	if headStart < 0 {
		return false
	}
	if bytescmp(mime, []byte("word/"), headStart+HEAD_SIZE) {
		return true
	}

	return false
}

func Pdf(mime []byte) bool {
	// bytes(25 50 44 46 2D) ascii(%PDF-)
	return len(mime) >= 5 &&
		mime[0] == 0x25 && mime[1] == 0x50 &&
		mime[2] == 0x44 && mime[3] == 0x46 &&
		mime[4] == 0x2D
}

func Odt(mime []byte) bool {
	return checkOdf(mime, TypeOdt.Mime.Value())
}

// this method is ALMOST the exact copy of https://github.com/h2non/filetype/blob/master/matchers/document.go
func checkOdf(mime []byte, mimetype string) bool {
	if 38+len(mimetype) >= len(mime) {
		return false
	}

	if !bytescmp(mime, zipMagic[:], 0) {
		return false
	}

	if mime[8] != 0 || mime[9] != 0 {
		return false
	}

	if mime[26] != 8 || mime[27] != 0 {
		return false
	}

	if int(mime[18]) != len(mimetype) ||
		mime[19] != 0 || mime[20] != 0 || mime[21] != 0 ||
		int(mime[22]) != len(mimetype) ||
		mime[23] != 0 || mime[24] != 0 || mime[25] != 0 {
		return false
	}

	if mime[28] != 0 || mime[29] != 0 {
		return false
	}

	// Optimised the code a little to avoid string conversion and allocation
	return bytescmp(mime, []byte("mimetype"), 30) &&
		bytescmp(mime, unsafe.Slice(unsafe.StringData(mimetype), len(mimetype)), 38)
}
