// dev1 — a deliberately tiny web app for testing Beast Deploy.
// Platform conventions: listens on $PORT, logs to stdout, /healthz for gates.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var (
	hits    atomic.Int64
	started = time.Now()
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host, _ := os.Hostname()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		n := hits.Add(1)
		log.Printf("GET / from %s (hit #%d)", r.RemoteAddr, n)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<!doctype html><meta charset="utf-8"><title>dev1</title>
<body style="margin:0;min-height:100vh;display:flex;align-items:center;justify-content:center;background:#0a0a0a;color:#ededed;font-family:ui-monospace,monospace">
<div style="text-align:center;line-height:2">
  <div style="font-size:34px;font-weight:700;letter-spacing:.14em">dev1</div>
  <div style="color:#3dd68c">deployed on Beast Deploy</div>
  <div style="color:#a1a1a1;font-size:13px">
    replica <b style="color:#ededed">%s</b><br>
    hits <b style="color:#ededed">%d</b> · up <b style="color:#ededed">%s</b><br>
    %s
  </div>
</div></body>`, host, n, time.Since(started).Round(time.Second), time.Now().Format(time.RFC1123))
	})

	log.Printf("dev1 listening on :%s (replica %s)", port, host)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
