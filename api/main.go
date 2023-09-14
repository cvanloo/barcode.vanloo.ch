// Package main provides a simple prototype/test-implementation of the barcode
// API.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"image/png"
	"syscall/js"

	"github.com/cvanloo/barcode/code128"
)

func SupportedTypes() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve := args[0]
			reject := args[1]

			go func() {
				resp, err := supportedTypes()
				if err != nil {
					jsErr := js.Global().Get("Error")
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
					jsErr := js.Global().Get("Error")
					reject.Invoke(jsErr.New(err.Error()))
				} else {
					js.CopyBytesToJS(buf, resp)
					resolve.Invoke(true)
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

	bc, err := code128.Encode(text)
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

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("supportedTypes", SupportedTypes())
	js.Global().Set("createBarcode", CreateBarcode())

	<-c
	//select {}
}
