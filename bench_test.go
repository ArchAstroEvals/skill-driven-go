package main

import "testing"

func BenchmarkStoreCreate(b *testing.B) {
  s := NewStore()
  for i := 0; i < b.N; i++ {
    s.Create(map[string]any{"name": "bench"})
  }
}
