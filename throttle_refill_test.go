package main

import ("testing"
  "time")

func TestRefill(t *testing.T) {
  b := NewBucket(1, 10*time.Millisecond)
  now := time.Now()
  b.Allow(now)
  if b.Allow(now) {
    t.Fatal("want denied before window")
  }
  if !b.Allow(now.Add(20*time.Millisecond)) {
    t.Fatal("want allowed after window")
  }
}
