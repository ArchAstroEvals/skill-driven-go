package main

import "testing"

func TestConfig(t *testing.T) {
  t.Setenv("PORT", "")
  t.Setenv("TOKEN", "")
  cfg := loadConfig()
  if cfg.Port != "8080" || cfg.Token != "tok-dev" {
    t.Fatalf("bad defaults %+v", cfg)
  }
  t.Setenv("PORT", "9090")
  if loadConfig().Port != "9090" {
    t.Fatal("want env override")
  }
}
