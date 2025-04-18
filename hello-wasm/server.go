package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed index.html
//go:embed main.wasm
//go:embed wasm_exec.js
var assets embed.FS

// NOTE:
// 1. Same go package can have multiple main function entry-points
// and can be called via go build server.go (init function is
// another exception in this rule) - go run server.go .

func main() {

	fs := http.FileServer(http.FS(assets))
	http.Handle("/", fs)
	// NOTE:
	// 1. No API endpoints, therefore no need to supply any handlers
	log.Print("Listening on 127.0.0.1:30001...")
	err := http.ListenAndServe("127.0.0.1:30001", nil)

	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
