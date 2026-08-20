# Plan 03: Error bodies

Objective: every error response is JSON with an error code.

Steps:
1. Shared helpers for not_found, unprocessable, unauthorized.
2. Router uses the helpers for 404, 401, 405.

Acceptance:
- Error bodies decode to maps with an error key.
