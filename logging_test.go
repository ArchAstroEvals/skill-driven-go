package main

import ("bytes"
  "log"
  "os"
  "strings"
  "testing")

func TestLogging(t *testing.T) {
  var buf bytes.Buffer
  log.SetOutput(&buf)
  defer log.SetOutput(os.Stderr)
  srv := NewServer("tok-dev")
  doRequest(t, srv, "GET", "/health", "", "")
  if !strings.Contains(buf.String(), "GET /health 200") {
    t.Fatalf("want log line got %q", buf.String())
  }
}
