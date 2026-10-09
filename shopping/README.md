# Shopping

Rust service boundary for the shopping module.

It will own shopping-list snapshots and consume versioned HTTP APIs from
`meal-calendar` and `recipes`. It must not access either service's database
directly. Quantity arithmetic is implemented locally against the versioned
quantity conformance fixtures described in `discovery/design.md`.

The HTTP API and behaviour are intentionally not scaffolded yet: implement them
through the amended M015–M016 work items in `discovery/module-work.md`, following
`discovery/design.md`.

## Development

```sh
cargo fmt --check
cargo test
cargo clippy -- -D warnings
```
