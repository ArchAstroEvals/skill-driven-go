package main

import "sync"

type Record map[string]any

type Store struct {
	mu     sync.Mutex
	next   int
	items  map[int]Record
}

func NewStore() *Store {
	return &Store{items: map[int]Record{}}
}

func (s *Store) Create(attrs map[string]any) Record {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	id := s.next
	rec := Record{"id": id}
	for k, v := range attrs {
		rec[k] = v
	}
	s.items[id] = rec
	return rec
}

func (s *Store) Get(id int) (Record, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rec, ok := s.items[id]
	return rec, ok
}
