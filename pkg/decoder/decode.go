package decoder

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func Decode(encname string, b []byte) ([]byte, error) {
	enc, err := enc(encname)
	if err != nil {
		return nil, err
	}
	r := bytes.NewBuffer(b)
	decoded, err := ioutil.ReadAll(transform.NewReader(r, enc.NewDecoder()))
	return decoded, err
}

func enc(encname string) (enc encoding.Encoding, err error) {
	switch encname {
	case "ShiftJIS":
		enc = japanese.ShiftJIS
	case "EUCJP":
		enc = japanese.EUCJP
	case "ISO2022JP":
		enc = japanese.ISO2022JP
	default:
		err = fmt.Errorf("Unknown encname %s", encname)
	}
	return
}
