package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

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

func toCSV(records []Record, fields []string) string {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.Write(fields)
	for _, rec := range records {
		row := make([]string, 0, len(fields))
		for _, f := range fields {
			row = append(row, csvCell(rec[f]))
		}
		w.Write(row)
	}
	w.Flush()
	return buf.String()
}

func csvCell(v any) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

