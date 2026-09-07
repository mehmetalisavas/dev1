# dev1

A deliberately tiny Go web app used as a Beast Deploy test app.

- Listens on `$PORT` (default 8080)
- `GET /` — status page (replica hostname, hit counter, uptime)
- `GET /healthz` — health gate endpoint
- Logs every request to stdout

Deploy: pick this repo in Beast Deploy (Create → Service — GitHub repo),
port `8080`, health path `/healthz`. Every push to `main` deploys.
