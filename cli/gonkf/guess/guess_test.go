package guess

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
)

func TestRun(t *testing.T) {
	testCase := []struct {
		e   detect.CharEncoding
		txt []byte
	}{
		{e: detect.UTF8, txt: []byte("Hello World")},
		{e: detect.UTF8, txt: []byte("こんにちは。世界の国から。")},
		{e: detect.ShiftJIS, txt: []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x42, 0x90, 0xa2, 0x8a, 0x45, 0x82, 0xcc, 0x8d, 0x91, 0x82, 0xa9, 0x82, 0xe7, 0x81, 0x42}},
		{e: detect.EUCJP, txt: []byte{0xa4, 0xb3, 0xa4, 0xf3, 0xa4, 0xcb, 0xa4, 0xc1, 0xa4, 0xcf, 0xa1, 0xa3, 0xc0, 0xa4, 0xb3, 0xa6, 0xa4, 0xce, 0xb9, 0xf1, 0xa4, 0xab, 0xa4, 0xe9, 0xa1, 0xa3}},
		{e: detect.ISO2022JP, txt: []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}},
	}

	for _, tst := range testCase {
		e := Run(bytes.NewReader(tst.txt))
		if e != tst.e {
			t.Errorf("Encoding(%v)  = \"%v\", want \"%v\".", tst.txt, e, tst.e)
		}
	}
}
