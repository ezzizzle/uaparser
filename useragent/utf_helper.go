package useragent

// http://play.golang.org/p/fVf7duRtdH

import (
	"bytes"
	"fmt"
	"unicode/utf16"
	"unicode/utf8"
)

// DecodeUTF16 decoes utf-16 bytes to a string
func DecodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)

	// Check endianess
	var littleEndian bool
	if b[0] == 0xff {
		littleEndian = true
	} else {
		littleEndian = false
	}

	for i := 2; i < lb; i += 2 {
		if littleEndian {
			u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		} else {
			u16s[0] = uint16(b[i+1]) + (uint16(b[i]) << 8)
		}

		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}
