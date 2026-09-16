package main

import (
  "sync"
  "testing")

func TestConcurrentCreates(t *testing.T) {
  s := NewStore()
  var wg sync.WaitGroup
  for i := 0; i < 50; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      s.Create(map[string]any{"name": "c"})
    }()
  }
  wg.Wait()
  if len(s.List()) != 50 {
    t.Fatalf("want 50 got %d", len(s.List()))
  }
}
