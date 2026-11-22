# The meal-prep image: the web app built by Astro, served by the Go server. VERSION is the commit.
FROM node:24-slim AS web
WORKDIR /web
COPY web/package.json web/package-lock.json ./
RUN npm ci
COPY web/ ./
RUN npm run build

FROM golang:1.27 AS server
ARG VERSION=dev
WORKDIR /src
COPY go.mod ./
COPY cmd/ cmd/
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o /server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=server /server /server
COPY --from=web /web/dist /web
ENV STATIC_DIR=/web PORT=8080
EXPOSE 8080
ENTRYPOINT ["/server"]
