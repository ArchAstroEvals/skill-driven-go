package main

import (
	"os"
	"strconv"
)

type Config struct {
	Port         string
	Token        string
	ReadTimeout  int
	WriteTimeout int
}

func loadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if _, err := strconv.Atoi(port); err != nil {
		port = "8080"
	}
	token := os.Getenv("TOKEN")
	if token == "" {
		token = "tok-dev"
	}
	return Config{Port: port, Token: token, ReadTimeout: 5, WriteTimeout: 10}
}
