package main

func filterBy(records []Record, field, value string) []Record {
	out := []Record{}
	for _, rec := range records {
		if v, ok := rec[field]; ok && v == value {
			out = append(out, rec)
		}
	}
	return out
}
