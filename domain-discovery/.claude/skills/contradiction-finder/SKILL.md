---
name: contradiction-finder
description: Find conflicting or mutually incompatible statements in a discovery evidence corpus. Report conflicts with quotes and clarification questions. Never resolve conflicts silently.
---

# Contradiction Finder

## What counts

A contradiction exists when two or more evidence statements cannot all be true:

- the same rule stated incompatible ways
- conflicting numbers, thresholds or states
- mutually exclusive requirements
- a decision that violates a stated constraint
- a later statement silently reverses an earlier one
- different meanings for the same domain term

## What does not count

Do not report as contradiction:

- different levels of detail
- examples vs general rules unless incompatible
- missing information
- opinions that can coexist
- implementation alternatives not yet decided

## Output file

`discovery/<slug>/contradictions.md`

## Format

```md
# Contradictions

## C001 — <short title>

### Clashing evidence

> <quote A>

- source: <source-ref>

> <quote B>

- source: <source-ref>

### Why this conflicts

<one short explanation>

### Clarification question

<question to ask the user>
```

If no contradictions are found, write:

```md
# Contradictions

No blocking contradictions found.
```

## Rules

- Quote evidence verbatim.
- Do not pick a winner.
- Do not rewrite the corpus.
- If resolving the conflict affects candidate tasks, mark the affected candidates as `needs-questions`.
