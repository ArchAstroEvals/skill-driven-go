package main

import (
	"net/http"
	"os"
)

func main() {
	cfg := loadConfig()
	srv := NewServer(cfg.Token)
	if err := http.ListenAndServe(":"+cfg.Port, srv); err != nil {
		os.Exit(1)
	}
}
