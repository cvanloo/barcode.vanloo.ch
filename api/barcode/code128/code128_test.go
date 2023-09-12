package code128

import (
	"image"
	_ "image/png"
	"os"
	"testing"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		path, expected string
	}{
		{"testfiles/test_code128-1.png", "ABCD-1234-abcd"},
		{"testfiles/test_code128-2.png", "PJJ123C"},
		{"testfiles/test_code128-3.png", "hello world"},
		{"testfiles/test_code128-4.png", "hello, world!"},
		{"testfiles/test_code128-5.png", "3456abcd"},
		{"testfiles/test_code128-6.png", "667390"},
		{"testfiles/test_code128-7.png", "biz\n"},
		{"testfiles/test_code128-8.png", "ABCDEFG"},

		// Dirty images
		{"testfiles/ClearCutGray.png", "hello"},
		{"testfiles/ClearCutDither.png", "hello"},
		{"testfiles/ClearCutBlackAround.png", "hello"},
		{"testfiles/ClearCutWhiteAround.png", "hello"},
	}
	code128 := Code128{}
	for _, c := range cases {
		f, err := os.Open(c.path)
		if err != nil {
			t.Error(err)
			continue
		}
		img, _, err := image.Decode(f)
		if err != nil {
			t.Error(err)
			continue
		}
		bs, err := code128.Decode(img)
		if err != nil {
			t.Error(err)
		}
		if string(bs) != c.expected {
			t.Errorf("got: `%s', want: `%s'", string(bs), c.expected)
			continue
		}
		t.Logf("%s: %s", c.path, string(bs))
	}
}

func TestDecodeRightToLeft(t *testing.T) {
	f, err := os.Open("testfiles/test_code128-rotate.png")
	if err != nil {
		t.Fatal(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	code128 := Code128{}
	bs, err := code128.Decode(img)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bs))
}
