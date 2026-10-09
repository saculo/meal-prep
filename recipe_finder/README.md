# Recipe Finder

Go service boundary for the read-only recipe-finder module.

It will consume versioned recipe-summary and recency HTTP APIs from `recipes`
and `meal-calendar`. It owns no database and must not access either service's
database directly.

The HTTP API and behaviour are intentionally not scaffolded yet: implement them
through the amended M012–M014 work items in `discovery/module-work.md`, following
`discovery/design.md`.

## Development

```sh
go test ./...
go vet ./...
gofmt -w .
```
