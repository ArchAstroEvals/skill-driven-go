package main

import "os"

type Config struct {
	Port  string
	Token string
}

func loadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	token := os.Getenv("TOKEN")
	if token == "" {
		token = "tok-dev"
	}
	return Config{Port: port, Token: token}
}
