---
name: context-mapping
description: Derive evidence-backed bounded context candidates and a context map from discovery artifacts. Separates business areas from bounded contexts, marks weak boundaries as hypotheses, and captures relationships/published language.
---

# Context Mapping

## Purpose

Turn discovery outputs into candidate bounded contexts and their relationships.

This skill does **not** design code modules yet. It names modeling boundaries from evidence and records uncertainty explicitly.

## Inputs

Prefer these files when available:

- `discovery/<slug>/corpus.md`
- `discovery/<slug>/problem-brief.md`
- `discovery/<slug>/work-candidates.md`
- `discovery/<slug>/event-storm.md`
- `discovery/<slug>/contradictions.md`
- `discovery/<slug>/questions.md`

## Output files

- `discovery/<slug>/context-candidates.md`
- `discovery/<slug>/context-map.md`

## Boundary signals

Use evidence for:

- different decision owners or actors
- different business capabilities
- distinct vocabulary / terms with different meanings
- invariants/rules that must be protected together
- lifecycle ownership of important domain objects
- events emitted/consumed between areas
- data needed from another area but not owned by it
- external systems and anti-corruption needs
- hotspots, contradictions, or frequent change pressure
- core/supporting/generic business importance

## Do not confuse

- A topic, screen, table, CRUD entity, or team name is not automatically a bounded context.
- A business area is only a candidate input to context design.
- A shared word is not a shared model if different groups mean different things.
- An implementation package is not proof of a domain boundary.

## Context candidate format

```md
# Context Candidates — <topic>

## C001 — <Context name>

- boundary confidence: high|medium|low
- domain classification: core|supporting|generic|unknown
- problem class: CRUD/state bookkeeping|transformation/presentation|integration|resource contention/concurrency|safety/guard/evidence|workflow/orchestration|mixed|unknown
- responsibility: <one sentence>
- owns: <concepts / decisions / lifecycle>
- does not own: <important exclusions>
- upstream/downstream signals: <events, queries, commands, policies>
- evidence: Q001, Q004
- open questions: <questions or None>
- boundary risks: <risks or None>
```

## Context map format

```md
# Context Map — <topic>

## Summary

<short paragraph>

## Bounded contexts

| Context | Classification | Problem class | Responsibility | Confidence | Evidence |
|---|---|---|---|---|---|

## Relationships

Direction is **downstream → upstream**: the dependent context is named first.

| # | Dependency | What downstream needs | Mechanism | Pattern | Evidence | Open questions |
|---|---|---|---|---|---|---|

## Published language / shared kernel candidates

| Concept | Kind | Meaning | Used by | Evidence | Risk |
|---|---|---|---|---|---|

## Domain events

| Event | Published by | Consumed by | Payload language | Evidence | Confidence |
|---|---|---|---|---|---|

## External systems / ACL candidates

| System | Context touching it | Translation needed | Evidence | Risk |
|---|---|---|---|---|

## Boundary hypotheses needing validation

- <hypothesis, why uncertain, question to ask>
```

## Relationship patterns

Use these labels when helpful:

- Published events / Published Language
- Customer-Supplier
- Open-Host Service
- Anti-Corruption Layer
- Shared Kernel
- Separate Ways
- Partnership
- Conformist
- Unknown

## Rules

- Every non-trivial context or relationship claim cites evidence.
- Keep contradictions visible; do not resolve them silently.
- Mark low-evidence boundaries as hypotheses.
- Prefer fewer, clearer contexts over many noun-based contexts.
- Keep implementation suggestions out unless explicitly marked as hypotheses.
