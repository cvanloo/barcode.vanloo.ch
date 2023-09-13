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
		//{"testfiles/test_code128-1.png", "ABCD-1234-abcd"},
		//{"testfiles/test_code128-2.png", "PJJ123C"},
		//{"testfiles/test_code128-3.png", "hello world"},
		//{"testfiles/test_code128-4.png", "hello, world!"},
		//{"testfiles/test_code128-5.png", "3456abcd"},
		{"testfiles/test_code128-6.png", "667390"},
		//{"testfiles/test_code128-7.png", "biz\n"},
		//{"testfiles/test_code128-8.png", "ABCDEFG"},

		// Rotated
		//{"testfiles/test_code128-rotate.png", "ABCD-1234-abcd"},

		// Dirty images
		//{"testfiles/ClearCutGray.png", "hello"},
		//{"testfiles/ClearCutDither.png", "hello"},
		//{"testfiles/ClearCutBlackAround.png", "hello"},
		//{"testfiles/ClearCutWhiteAround.png", "hello"},

		// Data after stop
		//{"testfiles/test_code128-data-after-stop.png", "Hello, World!"}, FIXME: failing
	}
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
		bs, err := Decode(img)
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

func TestEncode(t *testing.T) {
	cases := []string{
		"Hello, World!",
		"11223467",
		"\026\025",
	}

	for _, c := range cases {
		img, err := Encode(c)
		if err != nil {
			t.Error(err)
			continue
		}
		bs, err := Decode(img)
		if err != nil {
			t.Error(err)
			continue
		}
		if string(bs) != c {
			t.Errorf("got: `%s', want: `%s'", string(bs), c)
		}
	}

}
