package decode

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestToUTF8ja(t *testing.T) {
	testCase := []struct {
		txt []byte
		res []byte
	}{
		{txt: []byte("Hello World"), res: []byte("Hello World")},
		{txt: []byte("こんにちは。世界の国から。"), res: []byte("こんにちは。世界の国から。")},
		{txt: []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x42, 0x90, 0xa2, 0x8a, 0x45, 0x82, 0xcc, 0x8d, 0x91, 0x82, 0xa9, 0x82, 0xe7, 0x81, 0x42}, res: []byte("こんにちは。世界の国から。")},
		{txt: []byte{0xa4, 0xb3, 0xa4, 0xf3, 0xa4, 0xcb, 0xa4, 0xc1, 0xa4, 0xcf, 0xa1, 0xa3, 0xc0, 0xa4, 0xb3, 0xa6, 0xa4, 0xce, 0xb9, 0xf1, 0xa4, 0xab, 0xa4, 0xe9, 0xa1, 0xa3}, res: []byte("こんにちは。世界の国から。")},
		{txt: []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}, res: []byte("こんにちは世界")},
	}

	for _, tst := range testCase {
		res, err := ToUTF8ja(bytes.NewReader(tst.txt))
		if err != nil {
			t.Errorf("ToUTF8ja(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if bytes.Compare(buf.Bytes(), tst.res) != 0 {
			t.Errorf("ToUTF8ja(%v)  = \"%v\", want \"%v\".", tst.txt, buf.String(), string(tst.res))
		}
	}
}

func ExampleToUTF8() {
	jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
	res, err := ToUTF8(bytes.NewReader(jisText))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf)
	// Output:
	// こんにちは世界
}

func ExampleToUTF8ja() {
	jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
	res, err := ToUTF8ja(bytes.NewReader(jisText))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf)
	// Output:
	// こんにちは世界
}
