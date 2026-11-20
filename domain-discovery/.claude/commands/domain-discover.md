---
description: Turn discussion notes, transcripts, issues or rough ideas into craftpath-ready small work candidates. Builds evidence, finds contradictions, asks clarification questions, writes a solution-free brief and candidate requirements/tasks.
argument-hint: "<path-or-inline-description> [--out discovery/<slug>] [--no-questions]"
allowed-tools: Read, Write, Edit, Glob, Grep, Bash, AskUserQuestion
---

# /domain-discover

Discover small craftpath-ready work items from messy input.

Input: `$ARGUMENTS`

## Load skills

Use these skills in order:

1. `evidence-corpus`
2. `contradiction-finder`
3. `guided-conversation`
4. `problem-brief`
5. `requirement-candidates`
6. `problem-classifier`
7. optionally `implementation-scout`

## Goal

Produce small, independent work candidates suitable for later execution with craftpath's `/craftpath-work` command. Do **not** implement anything. Do **not** modify the target implementation project.

## Output directory

Create `discovery/<slug>/` unless `--out` is provided.

Write:

- `corpus.md` — evidence quote bank
- `contradictions.md` — conflicts and clarification questions
- `questions.md` — open questions and answers gathered this run
- `problem-brief.md` — solution-free brief
- `work-candidates.md` — craftpath-ready candidate requirements/tasks

## Procedure

### 1. Resolve input

Input may be:

- a file path
- a directory of notes/docs
- inline text
- a mixture of paths and text

Read only relevant files. Ignore build artifacts, dependencies, `.git`, `node_modules`, and existing generated discovery folders unless explicitly named.

### 2. Build evidence corpus

Use `evidence-corpus`.

Extract verbatim quotes with source references. Preserve source language. Do not paraphrase inside quotes.

### 3. Find contradictions

Use `contradiction-finder`.

Write conflicts to `contradictions.md`. For each real conflict include:

- clashing quotes
- why they cannot both be true
- clarification question

### 4. Ask clarification questions

Use `guided-conversation` unless `--no-questions` is present.

Ask only questions that block producing useful small work candidates. Batch them. Record answers in `questions.md` and add answer quotes to the corpus as `user-supplied` evidence.

### 5. Write problem brief

Use `problem-brief`.

Strip solution guesses. Keep:

- current pain
- actors/users
- desired observable outcomes
- constraints
- risks
- non-goals
- evidence references

### 6. Generate candidate work items

Use `requirement-candidates`.

Each candidate must be small enough for one craftpath work item. Each candidate needs:

- stable id: `R001`, `R002`, ...
- short imperative or outcome title
- problem statement
- Given/When/Then scenarios
- out of scope
- assumptions
- unresolved questions
- evidence quotes
- readiness: `ready`, `needs-questions`, or `too-large`
- suggested `/craftpath-work "..."` command

### 7. Classify candidates

Use `problem-classifier`.

Classify each candidate as one or more of:

- CRUD/state bookkeeping
- transformation/presentation
- integration
- resource contention/concurrency
- safety/guard/evidence
- workflow/orchestration
- documentation/spec-only

Suggest craftpath skills: `backend`, `frontend`, `testing`, `architecture`, `infrastructure`, `ux`.

### 8. Optional implementation scout

Use `implementation-scout` only if there is enough implementation context and the user asked for implementation hints.

Scout for:

- likely algorithmic hotspot
- likely boundary/coupling hotspot
- whether a design task may be needed

Keep this separate from the requirement. Mark as hypothesis.

## Final response

Return a short summary:

- output directory
- number of candidates
- ready candidates
- questions still blocking work
- recommended next `/craftpath-work` command

Do not paste all files unless the user asks.
