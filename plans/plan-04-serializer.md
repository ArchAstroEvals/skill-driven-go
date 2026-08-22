# Plan 04: Field selection

Objective: callers choose which fields come back.

Steps:
1. selectFields picks a subset, dropping nils.
2. selectList maps over record lists.

Acceptance:
- Unrequested fields are absent from output.
