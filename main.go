package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	cfg := loadConfig()
	srv := NewServer(cfg.Token)
	httpd := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      srv,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}
	if err := httpd.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
