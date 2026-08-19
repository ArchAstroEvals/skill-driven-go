# Plan 02: Validation

Objective: reject records missing required fields or with wrong types.

Steps:
1. required() collects missing fields.
2. fieldType() classifies string, number, boolean, unknown.
3. Blank strings count as missing.

Acceptance:
- Unit tests cover missing, blank, and typed values.
