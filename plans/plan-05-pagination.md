# Plan 05: Pagination

Objective: slice listings with page and per_page.

Steps:
1. paginate() slices with 1-based pages.
2. Page zero and overflow return empty lists.

Acceptance:
- Page 1 holds the first per_page records.
