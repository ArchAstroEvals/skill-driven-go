package main

func notFound(resource string) map[string]any {
	return map[string]any{"error": "not_found", "resource": resource}
}

func unprocessable(fields []string) map[string]any {
	return map[string]any{"error": "unprocessable", "fields": fields}
}

func unauthorized() map[string]any {
	return map[string]any{"error": "unauthorized"}
}

func methodNotAllowed() map[string]any {
	return map[string]any{"error": "method_not_allowed"}
}
