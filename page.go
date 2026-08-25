package main

func paginate(records []Record, page, perPage int) []Record {
	start := (page - 1) * perPage
	if start >= len(records) {
		return []Record{}
	}
	end := start + perPage
	if end > len(records) {
		end = len(records)
	}
	return records[start:end]
}
