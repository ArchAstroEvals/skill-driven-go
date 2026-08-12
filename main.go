package main

import (
	"net/http"
	"os"
)

func main() {
	srv := NewServer("tok-dev")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		os.Exit(1)
	}
}
