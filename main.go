package main

import (
	"fmt"
	"mime"
	"net/http"
)

func main() {
	mime.AddExtensionType(".js", "application/javascript")
	//ct := mime.TypeByExtension(".js")
	//fmt.Printf("ct: %s\n", ct)

	err := http.ListenAndServe(":9090", http.FileServer(http.Dir(".")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
