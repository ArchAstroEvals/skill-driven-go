package main

import "strings"

func searchByName(records []Record, needle string) []Record {
	lowered := strings.ToLower(needle)
	out := []Record{}
	for _, rec := range records {
		if name, ok := rec["name"].(string); ok && strings.Contains(strings.ToLower(name), lowered) {
			out = append(out, rec)
		}
	}
	return out
}
