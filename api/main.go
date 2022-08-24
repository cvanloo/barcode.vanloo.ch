// Package main provides a simple prototype/test-implementation of the barcode
// API.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	enableCORS := func(w *http.ResponseWriter) {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

	mux.HandleFunc("/supported_types", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		supported := []struct {
			Value, Name string
		}{{
			Value: "code-128",
			Name:  "Code-128",
		}, {
			Value: "gs1-128",
			Name:  "GS1-128",
		}}

		bs, err := json.Marshal(supported)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(bs)
	})

	mux.HandleFunc("/create_barcode", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		r.ParseForm()
		bc := r.Form.Get("type")
		text := r.Form.Get("text")
		fmt.Printf("type: %s, text: %s", bc, text)

		b, err := os.ReadFile("barcode.png")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(b)
	})

	fmt.Println("Listening and serving on :8080")
	http.ListenAndServe(":8080", mux)
}
