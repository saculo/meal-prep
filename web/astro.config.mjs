// A static site: `astro build` writes web/dist, which the Go server serves. In `astro dev`, /api
// and /version go to the Go server on :8080.
import { defineConfig } from "astro/config"

export default defineConfig({
  output: "static",
  vite: {
    server: { proxy: { "/api": "http://localhost:8080", "/version": "http://localhost:8080" } },
  },
})
