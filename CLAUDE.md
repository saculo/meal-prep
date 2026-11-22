# meal-prep

Meal planning for a two-person household. The discovery is in `all-event-storms/`
(start with `problem-brief.md`); the toolkit that made it is in `domain-discovery/`.

- `cmd/server/` — the Go server (standard library only, so far): the API under `/api`,
  `/version` (the commit it was built from), and the built web app from `STATIC_DIR`.
- `web/` — the web app, Astro, built to static files in `web/dist`.
- `Dockerfile` — one image with both; CI pushes `ghcr.io/saculo/meal-prep:<commit>`
  for every pull request and every push to main.

Run it: `cd web && npm ci && npm run build && cd .. && go run ./cmd/server` (port 8080),
or `npm run dev` in `web/` with the Go server running (it proxies `/api` and `/version`).

Checks (CI and orq's gate, `.orqestrai/config.toml`): `gofmt -l .` empty, `go vet ./...`,
`go test ./...`; in `web/`: `npm run check`, `npm run build`.
