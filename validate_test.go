package main

import ("testing"
  "reflect")

func TestRequired(t *testing.T) {
  if got := required(map[string]any{"name": "x"}, []string{"name"}); len(got) != 0 {
    t.Fatal("want none missing")
  }
  if got := required(map[string]any{}, []string{"name"}); !reflect.DeepEqual(got, []string{"name"}) {
    t.Fatalf("want name missing got %v", got)
  }
}
