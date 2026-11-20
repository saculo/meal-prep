---
name: problem-brief
description: Distill messy discovery material into a short, solution-free problem brief for craftpath. Keeps business/product intent and strips implementation guesses.
---

# Problem Brief

## Purpose

Create the bridge between raw discovery and craftpath `requirement.md`.

## Output file

`discovery/<slug>/problem-brief.md`

## Format

```md
# Problem Brief — <topic>

## Current pain

## Actors / users

## Desired observable outcomes

## Rules and constraints

## Risks / hotspots

## Metrics or success signals

## Out of scope / non-goals

## Evidence index
```

## Strip list

Remove or demote to hypotheses:

- specific class/function names unless the input is explicitly about developer tooling
- database/schema guesses
- UI component guesses
- framework/library choices
- architecture choices
- task decomposition

Keep implementation details only if they are constraints from existing reality.

## Rules

- Focus on problem and observable outcome.
- Prefer short bullet points.
- Each important statement should reference corpus quote ids.
- If evidence is weak, say so.
