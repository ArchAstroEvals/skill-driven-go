# Plan 08: Bulk and CSV

Objective: insert many records at once and export listings.

Steps:
1. BulkCreate validates every item before storing any.
2. toCSV renders header plus rows with encoding/csv.
3. Commas in values are quoted.

Acceptance:
- A single invalid item rejects the whole batch.
- CSV output parses back to the same fields.
