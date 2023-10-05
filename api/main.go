// Package main provides a simple prototype/test-implementation of the barcode
// API.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	_ "image/jpeg"
	"syscall/js"

	"github.com/cvanloo/barcode/code128"
)

var jsErr = js.Global().Get("Error")

func SupportedTypes() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go func() {
				resp, err := supportedTypes()
				if err != nil {
					reject.Invoke(jsErr.New(err.Error()))
				} else {
					resolve.Invoke(resp)
				}
			}()

			return nil
		})

		promise := js.Global().Get("Promise")
		return promise.New(handler)
	})
}

func supportedTypes() (string, error) {
	supported := []struct {
		Value, Name string
	}{{
		Value: "code-128",
		Name:  "Code-128",
	}, /*{
		Value: "gs1-128",
		Name:  "GS1-128",
	}*/}

	bs, err := json.Marshal(supported)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func CreateBarcode() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		typ := args[0].String()
		text := args[1].String()
		buf := args[2]

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go func() {
				resp, err := createBarcode(typ, text)
				if err != nil {
					reject.Invoke(jsErr.New(err.Error()))
				} else {
					js.CopyBytesToJS(buf, resp)
					resolve.Invoke(len(resp))
				}
			}()

			return nil
		})

		promise := js.Global().Get("Promise")
		return promise.New(handler)
	})
}

func createBarcode(typ, text string) ([]byte, error) {
	if typ != "code-128" {
		return nil, errors.New("barcode type not supported")
	}

	cstr, err := code128.NewASCII(text)
	if err != nil {
		return nil, err
	}

	bc, err := code128.Encode(cstr)
	if err != nil {
		return nil, err
	}

	img, err := bc.Scale(312, 80)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeBarcode() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		typ := args[0].String()
		jsBuf := args[1]
		size := args[2].Int()

		buf := make([]byte, 5*1024*1024)
		if n := js.CopyBytesToGo(buf, jsBuf); n != size {
			return jsErr.New(errors.New("failed to copy memory from JS to Go"))
		}

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go func() {
				img, _, err := image.Decode(bytes.NewBuffer(buf[:size]))
				if err != nil {
					reject.Invoke(jsErr.New(err.Error()))
					return
				}
				text, err := decodeBarcode(typ, img)
				if err != nil {
					reject.Invoke(jsErr.New(err.Error()))
					return
				}
				resolve.Invoke(text)
			}()

			return nil
		})

		promise := js.Global().Get("Promise")
		return promise.New(handler)
	})
}

func decodeBarcode(typ string, img image.Image) (text string, err error) {
	if typ != "code-128" {
		return "", errors.New("barcode type not supported")
	}

	bs, _, err := code128.Decode(img)
	return string(bs), err
}

func main() {
	js.Global().Set("supportedTypes", SupportedTypes())
	js.Global().Set("createBarcode", CreateBarcode())
	js.Global().Set("decodeBarcode", DecodeBarcode())

	select{}
}
