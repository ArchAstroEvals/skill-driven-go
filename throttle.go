package main

import "time"

type Bucket struct {
	limit  int
	window time.Duration
	used   int
	reset  time.Time
	empty  bool
}

func NewBucket(limit int, window time.Duration) *Bucket {
	return &Bucket{limit: limit, window: window, empty: true}
}

func (b *Bucket) Allow(now time.Time) bool {
	if b.empty {
		b.empty = false
		b.reset = now.Add(b.window)
	}
	if b.used >= b.limit {
		return false
	}
	b.used++
	return true
}
