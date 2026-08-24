package main

import "testing"

func TestTokens(t *testing.T) {
  if !validToken("Bearer tok-abc", "tok-abc") {
    t.Fatal("want valid token")
  }
  if validToken("Bearer nope", "tok-abc") {
    t.Fatal("want invalid rejected")
  }
  if validToken("", "tok-abc") {
    t.Fatal("want empty rejected")
  }
}
