package types

import (
	"strings"
)

type MIME struct {
	Type 	string
	Subtype string
}

func BuildMIME(mime string) MIME {
	parts := strings.Split(mime, "/")
	if len(parts) != 2 || len(parts[0]) <= 0 || len(parts[1]) <= 0 {
		return MIME{}
	}
	return MIME{parts[0], parts[1]}
}

func (mime MIME) Value() string {
	return mime.Type + "/" + mime.Subtype
}