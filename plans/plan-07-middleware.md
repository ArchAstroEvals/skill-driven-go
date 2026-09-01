# Plan 07: Middleware

Objective: cross-cutting behavior via chained middleware.

Steps:
1. Logging middleware records method, path, status.
2. Request ID middleware sets and echoes X-Request-ID.
3. CORS middleware sets the allow-origin header.
4. Throttle bucket limits bursts with refill.

Acceptance:
- Each middleware has a handler-level test.
