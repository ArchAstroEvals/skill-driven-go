package main

import ("testing"
  "time")

func TestThrottle(t *testing.T) {
  b := NewBucket(1, time.Second)
  now := time.Now()
  if !b.Allow(now) {
    t.Fatal("want first allowed")
  }
  if b.Allow(now) {
    t.Fatal("want second denied")
  }
}
