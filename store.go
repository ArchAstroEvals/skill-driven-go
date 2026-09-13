package main

import (
	"sync"
	"time"
)

type Record map[string]any

type Store struct {
	mu    sync.Mutex
	next  int
	items map[int]Record
}

func NewStore() *Store {
	return &Store{items: map[int]Record{}}
}

func (s *Store) Create(attrs map[string]any) Record {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.next++
	id := s.next
	rec := Record{"id": id, "created_at": time.Now().Unix()}
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

func (s *Store) List() []Record {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]Record, 0, len(s.items))
	for _, rec := range s.items {
		out = append(out, rec)
	}
	return out
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return false
	}
	delete(s.items, id)
	return true
}

func (s *Store) BulkCreate(items []map[string]any) ([]Record, []string) {
	for _, attrs := range items {
		if missing := required(attrs, []string{"name"}); len(missing) > 0 {
			return nil, missing
		}
	}
	out := make([]Record, 0, len(items))
	for _, attrs := range items {
		out = append(out, s.Create(attrs))
	}
	return out, nil
}

func (s *Store) Update(id int, attrs map[string]any) (Record, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rec, ok := s.items[id]
	if !ok {
		return nil, false
	}
	for k, v := range attrs {
		if k != "id" {
			rec[k] = v
		}
	}
	return rec, true
}
