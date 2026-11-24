# 0001 Services ask each other over gRPC and never copy data

**Status:** accepted (2026-11-23)

## Context
The first DESIGN.md had services talk over gRPC and NATS JetStream: an outbox per schema,
durable consumers, and in planning a copy of recipe names and ingredients fed by `recipes.*`
events, used for suggestions. That copy broke "none reads another's data", and nothing in
the discovery needs events yet.

## Decision
No NATS, events or outbox until a feature needs them. A service asks another over gRPC
and never stores the answer in its own schema. Planning asks recipes for the recipes that
match the filters and selected ingredients, then sorts them by its own eaten history.
Shopping asks planning for planned meals and recipes for their lines. Shopping's stored
list (name, amount, unit at generation) is its own record, not a synced copy.

## Consequences
- Suggestions and shopping lists need recipes (and planning) running at request time; no
  local fallback.
- Recipes must offer filtering by tags and selected ingredients in its contract, because
  planning cannot filter a local copy.
- Recipe deletion is not defined yet; with no events, nothing cascades.
- When a feature needs events, that is a new decision (broker, outbox, which data may be
  copied); this ADR is then superseded, not edited.
