package main

import (
	"gordle-service/internal/handlers"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.NewRouter())

	if err != nil {
		// NOTE:
		// 1. We can log and exit using os.Exit(1)
		// but defer calls might or might not be executed
		// in this scenario.
		panic(err)
	}
}
