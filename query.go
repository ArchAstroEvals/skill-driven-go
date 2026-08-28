package main

import "strconv"

var knownFilterFields = []string{"kind", "name"}

func filterBy(records []Record, field, value string) []Record {
	out := []Record{}
	for _, rec := range records {
		if v, ok := rec[field]; ok && v == value {
			out = append(out, rec)
		}
	}
	return out
}

func filterKnown(records []Record, params map[string]string) []Record {
	out := records
	for _, field := range knownFilterFields {
		if value, ok := params[field]; ok && value != "" {
			out = filterBy(out, field, value)
		}
	}
	return out
}

func sortBy(records []Record, field string, desc bool) []Record {
	out := append([]Record{}, records...)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0; j-- {
			a, aok := out[j-1][field].(string)
			b, bok := out[j][field].(string)
			less := (!aok && bok) || (aok && bok && a > b)
			if desc {
				less = (aok && !bok) || (aok && bok && a < b)
			}
			if !less {
				break
			}
			out[j-1], out[j] = out[j], out[j-1]
		}
	}
	return out
}

func atoiOr(s string, def int) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}
