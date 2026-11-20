---
description: Derive evidence-backed bounded context candidates and a context map from discovery artifacts. Separates business areas from bounded contexts and records uncertainty.
argument-hint: "<discovery-dir-or-files> [--out discovery/<slug>] [--no-questions]"
allowed-tools: Read, Write, Edit, Glob, Grep, Bash, AskUserQuestion
---

# /context-map

Create a candidate bounded-context map from discovery artifacts.

Input: `$ARGUMENTS`

## Load skills

Use these skills in order:

1. `context-mapping`
2. `guided-conversation` when boundary questions block a useful map
3. `problem-classifier` for context problem-class hints
4. `contradiction-finder` only if contradictions were not already scanned

## Goal

Produce evidence-backed bounded context candidates and relationships. Do **not** implement code. Do **not** turn every noun into a context.

## Output directory

Use the discovery directory passed as input. If no directory is clear, create `discovery/<slug>/`.

Write:

- `context-candidates.md`
- `context-map.md`
- update `questions.md` if clarification questions are asked

## Procedure

### 1. Resolve input

Read relevant discovery files:

- `corpus.md`
- `problem-brief.md`
- `work-candidates.md`
- `event-storm.md`
- `contradictions.md`
- `questions.md`

Ignore implementation folders unless explicitly named.

### 2. Extract boundary signals

Use `context-mapping`.

Look for:

- decision ownership
- business capabilities
- vocabulary shifts
- invariants and lifecycle ownership
- events, commands, policies and external systems
- hotspots and change pressure
- core/supporting/generic importance

### 3. Ask only blocking questions

If boundaries are too ambiguous, ask a small batch of questions unless `--no-questions` is present. Record answers in `questions.md` and cite them as user-supplied evidence.

### 4. Produce context candidates

Write `context-candidates.md` using the skill format. Mark confidence and open questions.

### 5. Produce context map

Write `context-map.md` using the skill format:

- bounded contexts table
- downstream → upstream relationships
- published language/shared kernel candidates
- domain events
- external systems/ACL candidates
- boundary hypotheses needing validation

## Final response

Return a short summary:

- output directory
- number of context candidates
- high-confidence contexts
- relationships found
- blocking questions
- recommended next command, usually `/module-design <discovery-dir>`
