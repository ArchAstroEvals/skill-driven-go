package main

func required(attrs map[string]any, fields []string) []string {
	var missing []string
	for _, f := range fields {
		if attrs[f] == nil || attrs[f] == "" {
			missing = append(missing, f)
		}
	}
	return missing
}

func fieldType(v any) string {
	switch v.(type) {
	case string:
		return "string"
	case float64:
		return "number"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}
}
