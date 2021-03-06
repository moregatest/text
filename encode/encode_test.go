package encode

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
)

func TestFromUTF8To(t *testing.T) {
	testCase := []struct {
		e   detect.CharEncoding
		txt []byte
		res []byte
	}{
		{e: detect.ISO8859L1, res: []byte("Hello World"), txt: []byte("Hello World")},
		{e: detect.UTF8, res: []byte("こんにちは。世界の国から。"), txt: []byte("こんにちは。世界の国から。")},
		{e: detect.ShiftJIS, res: []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x42, 0x90, 0xa2, 0x8a, 0x45, 0x82, 0xcc, 0x8d, 0x91, 0x82, 0xa9, 0x82, 0xe7, 0x81, 0x42}, txt: []byte("こんにちは。世界の国から。")},
		{e: detect.EUCJP, res: []byte{0xa4, 0xb3, 0xa4, 0xf3, 0xa4, 0xcb, 0xa4, 0xc1, 0xa4, 0xcf, 0xa1, 0xa3, 0xc0, 0xa4, 0xb3, 0xa6, 0xa4, 0xce, 0xb9, 0xf1, 0xa4, 0xab, 0xa4, 0xe9, 0xa1, 0xa3}, txt: []byte("こんにちは。世界の国から。")},
		{e: detect.ISO2022JP, res: []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42, 0x0a}, txt: []byte("こんにちは世界\n")},
	}

	for _, tst := range testCase {
		res, err := FromUTF8To(tst.e, bytes.NewReader(tst.txt))
		if err != nil {
			t.Errorf("ToUTF8ja(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if bytes.Compare(buf.Bytes(), tst.res) != 0 {
			t.Errorf("ToUTF8ja(%v)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}

func ExampleFromUTF8To() {
	utf8Text := "こんにちは，世界\n"
	res, err := FromUTF8To(detect.ISO2022JP, bytes.NewBufferString(utf8Text))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf.Bytes())
	// Output:
	// [27 36 66 36 51 36 115 36 75 36 65 36 79 33 36 64 36 51 38 27 40 66 10]
}
