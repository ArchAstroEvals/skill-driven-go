package main

import "strings"

func searchByName(records []Record, needle string) []Record {
	out := []Record{}
	for _, rec := range records {
		if name, ok := rec["name"].(string); ok && strings.Contains(name, needle) {
			out = append(out, rec)
		}
	}
	return out
}
