package main

func selectFields(rec Record, fields []string) Record {
	out := Record{}
	for _, f := range fields {
		if v, ok := rec[f]; ok && v != nil {
			out[f] = v
		}
	}
	return out
}

func selectList(records []Record, fields []string) []Record {
	out := make([]Record, 0, len(records))
	for _, rec := range records {
		out = append(out, selectFields(rec, fields))
	}
	return out
}
