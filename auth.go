package main

func validToken(header, expected string) bool {
	return header == "Bearer "+expected
}
