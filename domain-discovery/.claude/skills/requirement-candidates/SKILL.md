---
name: requirement-candidates
description: Generate small craftpath-ready work candidates from a problem brief and evidence corpus. Each candidate maps to a possible `/craftpath-work` run and includes scenarios, scope, assumptions, evidence and readiness.
---

# Requirement Candidates

## Goal

Produce small independent candidates that can become craftpath work items.

A candidate is not a task file yet. It is a requirement seed for `/craftpath-work`.

## Output file

`discovery/<slug>/work-candidates.md`

## Candidate format

```md
# Work Candidates

## R001 — <short title>

- readiness: ready|needs-questions|too-large
- class: <classification>
- suggested skills: [backend|frontend|testing|architecture|infrastructure|ux]
- suggested command: `/craftpath-work "<title>"`

### Problem

<2-4 sentences, no solution design>

### Scenarios

Scenario: <name>
  Given <precondition>
  When <action>
  Then <observable outcome>

### Out of scope

- <excluded thing>

### Assumptions

- <assumption or `None`>

### Open questions

- <question or `None`>

### Evidence

- Q001 — <short reason>
- Q004 — <short reason>

### Implementation hints, not requirements

- <optional hypothesis; keep separate>
```

## Sizing rules

A candidate should be small enough that craftpath can decompose it into a few focused implementation tasks.

Mark `too-large` when it contains multiple independent outcomes. Split it into smaller R-items.

## Readiness rules

`ready` means:

- problem is clear
- at least one scenario is testable
- out of scope is known enough
- no blocking contradiction remains

`needs-questions` means:

- useful but one or more answers are required before `/craftpath-work`

`too-large` means:

- valuable but must be split before implementation

## Craftpath alignment

Scenarios should be compatible with craftpath's `requirement.md` template:

```text
Scenario: <name>
  Given <precondition>
  When <action>
  Then <observable outcome>
```

Avoid implementation-specific acceptance criteria at discovery time unless the requirement is itself developer tooling.
