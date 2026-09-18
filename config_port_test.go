package main

import "testing"

func TestPortFallback(t *testing.T) {
  t.Setenv("PORT", "abc")
  if cfg := loadConfig(); cfg.Port != "8080" {
    t.Fatalf("want 8080 got %q", cfg.Port)
  }
  t.Setenv("PORT", "9090")
  if cfg := loadConfig(); cfg.Port != "9090" {
    t.Fatalf("want 9090 got %q", cfg.Port)
  }
}
