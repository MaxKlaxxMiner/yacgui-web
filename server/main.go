package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"mime"
	"net/http"
)

func main() {
	yacboard.PerftTest(2)

	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".wasm", "application/wasm")

	http.Handle("/", http.FileServer(http.Dir("../client/html-content/")))

	fmt.Println("run server: localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
