# Plan 09: Docs and CI

Objective: document the API and run tests in CI.

Steps:
1. README covers run, test, endpoints, env vars.
2. CHANGELOG tracks releases.
3. CircleCI runs go test on every push.

Acceptance:
- A new checkout passes go test ./... from the README steps.
