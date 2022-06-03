package main

import (
	"github.com/MaxKlaxxMiner/yacgui-web/modules/yacboard"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	callerpath = filepath.Dir(b)
)

func main() {
	yacboard.PerftTest(2)

	//_ = mime.AddExtensionType(".js", "application/javascript")
	//_ = mime.AddExtensionType(".wasm", "application/wasm")
	//
	//contentFolder := filepath.Join(callerpath, "../client/html-content/")
	//_, err := os.ReadFile(filepath.Join(contentFolder, "index.html"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//http.Handle("/", http.FileServer(http.Dir(contentFolder)))
	//
	//fmt.Println("run server: localhost:9090")
	//err = http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	fmt.Println("Failed to start server", err)
	//	return
	//}
}
