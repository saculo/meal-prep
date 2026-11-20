---
description: Convert a context map into module blueprints and craftpath-ready module work candidates. Does not write implementation code.
argument-hint: "<discovery-dir-or-context-map> [--out discovery/<slug>] [--no-questions]"
allowed-tools: Read, Write, Edit, Glob, Grep, Bash, AskUserQuestion
---

# /module-design

Design module blueprints from a context map and discovery evidence.

Input: `$ARGUMENTS`

## Load skills

Use these skills in order:

1. `module-design`
2. `guided-conversation` when module boundary/API questions block useful work items
3. `requirement-candidates` for craftpath-sized scenario shaping
4. `problem-classifier` for suggested craftpath skills

## Goal

Produce module blueprints and module-level work candidates suitable for later `/craftpath-work` execution. Do **not** implement code.

## Output directory

Use the discovery directory passed as input. If no directory is clear, create `discovery/<slug>/`.

Write:

- `modules.md`
- `module-work.md`
- update `questions.md` if clarification questions are asked

## Procedure

### 1. Resolve input

Read relevant files:

- `context-map.md`
- `context-candidates.md`
- `work-candidates.md`
- `corpus.md`
- `problem-brief.md`
- `questions.md`

Read implementation/code notes only if explicitly supplied.

### 2. Build module blueprints

Use `module-design`.

For each context or context slice, define:

- module responsibility
- owned concepts and decisions
- explicit non-ownership
- public capabilities / entry points
- events published and consumed
- queries needed from other modules
- invariants/rules
- integration boundaries
- readiness and evidence

### 3. Ask only blocking questions

If public capabilities or module boundaries are too ambiguous, ask a small batch unless `--no-questions` is present. Record answers in `questions.md`.

### 4. Generate module work candidates

Write `module-work.md`. Each candidate should be small enough for one `/craftpath-work` run and include scenarios.

### 5. Keep implementation hypotheses separate

Do not prescribe frameworks, databases, endpoints or package structure unless the user supplied existing implementation constraints. Mark such notes as hypotheses.

## Final response

Return a short summary:

- output directory
- modules designed
- ready module work candidates
- candidates needing questions
- recommended next `/craftpath-work "..."` command
