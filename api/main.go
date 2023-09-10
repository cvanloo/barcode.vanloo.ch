// Package main provides a simple prototype/test-implementation of the barcode
// API.
package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"framagit.org/miya/barcode.vanloo.ch/api/code128"
)

func must[T any](t T, e error) T {
	if e != nil {
		panic(fmt.Sprintf("must: %v", e))
	}
	return t
}

func main() {
	//f := must(os.Open("code128/test_code128.png"))
	//f := must(os.Open("code128/test_code128-rotate.png"))
	//f := must(os.Open("code128/test_code128-1.png"))
	//f := must(os.Open("code128/test_code128-2.png"))
	//f := must(os.Open("code128/test_code128-3.png"))
	//f := must(os.Open("code128/test_code128-4.png"))
	//f := must(os.Open("code128/test_code128-5.png"))
	//f := must(os.Open("code128/test_code128-6.png"))
	//f := must(os.Open("code128/test_code128-7.png"))
	f := must(os.Open("code128/test_code128-8.png"))
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("decoding image: %v", err)
	}
	fmt.Println(must(code128.Decode(img)))
}

/*
func main2() {
	mux := http.NewServeMux()

	enableCORS := func(w *http.ResponseWriter) {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

	mux.HandleFunc("/api/supported_types", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		supported := []struct {
			Value, Name string
		}{{
			Value: "code-128",
			Name:  "Code-128",
		}, /*{
			Value: "gs1-128",
			Name:  "GS1-128",
		}*//*}

		bs, err := json.Marshal(supported)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(bs)
	})

	mux.HandleFunc("/api/create_barcode", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		r.ParseForm()
		bc := r.Form.Get("type")
		text := r.Form.Get("text")
		fmt.Printf("type: %s, text: %s\n", bc, text)

		b, err := code128.Encode(text)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bs, err := barcode.Scale(b, 312, 80)

		err = png.Encode(w, bs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
	})

	fmt.Println("Listening and serving on :8080")
	http.ListenAndServe(":8080", mux)
}
*/
