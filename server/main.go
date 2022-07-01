package main

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"server/lichess"
)

var (
	_, b, _, _ = runtime.Caller(0)
	callerpath = filepath.Dir(b)
)

func main() {
	lichess.Main()
	return

	//yacboard.PerftTest(2)

	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".wasm", "application/wasm")

	contentFolder := filepath.Join(callerpath, "../client/html-content/")
	_, err := os.ReadFile(filepath.Join(contentFolder, "index.html"))
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir(contentFolder)))

	fmt.Println("run server: localhost:9090")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
