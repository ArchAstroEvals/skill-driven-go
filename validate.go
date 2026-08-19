package main

func required(attrs map[string]any, fields []string) []string {
	var missing []string
	for _, f := range fields {
		if attrs[f] == nil {
			missing = append(missing, f)
		}
	}
	return missing
}
