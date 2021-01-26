all: static/main.wasm static/wasm_exec.js
	~/go/bin/goexec 'http.ListenAndServe(`:9999`, http.FileServer(http.Dir(`.`)))'

static/wasm_exec.js:
	cp "/usr/local/Cellar/go/1.15.6/misc/wasm/wasm_exec.js" static

static/main.wasm : main.go
	GO111MODULE=auto GOOS=js GOARCH=wasm go build -o static/main.wasm .