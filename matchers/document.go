package matchers

import (
	"encoding/binary"

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

	if !bytescmp(magic, zipMagic[:], 0) {
		return false
	}

	if bytescmp(magic, []byte("word/"), HEAD_SIZE) {
		return true
	}

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
	return false
}
