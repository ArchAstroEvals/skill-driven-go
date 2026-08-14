# Plan 01: Records CRUD

Objective: expose records over HTTP with create, fetch, list, delete.

Steps:
1. In-memory store with mutex.
2. POST /records returns 201 with the stored record.
3. GET /records/{id} returns 200 or 404.
4. GET /records returns the list.
5. DELETE /records/{id} returns the deleted record or 404.

Acceptance:
- go test ./... passes.
- Unknown paths return 404, wrong methods return 405.
