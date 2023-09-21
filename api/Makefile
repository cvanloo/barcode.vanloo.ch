.PHONY: build install clean
build: main.wasm

main.wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

install: main.wasm
	cp main.wasm ../web/public/main.wasm

clean:
	-rm main.wasm
