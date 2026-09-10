package main

import "testing"

func TestConfig(t *testing.T) {
  t.Setenv("PORT", "")
  t.Setenv("TOKEN", "")
  cfg := loadConfig()
  if cfg.Port != "8080" || cfg.Token != "tok-dev" {
    t.Fatalf("bad defaults %+v", cfg)
  }
  if cfg.ReadTimeout != 5 || cfg.WriteTimeout != 10 {
    t.Fatal("want timeout defaults")
  }
}
