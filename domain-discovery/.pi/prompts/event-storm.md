---
description: Run a lightweight Event Storming discovery interview and turn the conversation into evidence, events, commands, policies, hotspots and craftpath-ready work seeds.
argument-hint: "[domain/problem/context] [--out discovery/<slug>/event-storm.md]"
allowed-tools: Read, Write, Edit, Glob, Grep, Bash, AskUserQuestion
---

# /event-storm

Run a lightweight Event Storming interview for craftpath discovery.

Input: `$ARGUMENTS`

## Load skills

Use:

1. `lightweight-event-storming`
2. `guided-conversation`
3. `evidence-corpus`
4. `requirement-candidates` when deriving work seeds

## Purpose

Help the user discover process/domain knowledge before `/domain-discover` or `/craftpath-work`.

This is **not** full Event Storming simulation with fake stakeholders. You are the facilitator/scribe. Ask the real user questions and capture their answers.

## Procedure

1. Establish scope: process, product area, or user journey.
2. Ask for the happy path as domain events in past tense.
3. For each event, ask:
   - what command/action caused it?
   - who/what triggered it?
   - what policy/rule decided it?
   - what data was needed?
   - what can go wrong?
4. Mark hotspots, contradictions, unclear terms and external systems.
5. Ask for edge cases and failure paths.
6. Write the session artifact.
7. Derive small work seeds only when the evidence is strong enough.

## Output

Create `discovery/<slug>/event-storm.md` unless `--out` is provided.

Format:

```md
# Event Storm — <topic>

## Scope

## Timeline

| # | Event | Command/action | Actor | Policy/rule | Data | Source |
|---|---|---|---|---|---|---|

## Hotspots

## External systems

## Terms / ubiquitous language

## Open questions

## Work seeds

### R001 — <title>
Problem:
Scenario draft:
Evidence:
Readiness:
```

## Rules

- Events are past tense facts: `Task was approved`, not `Approve task`.
- Commands are intent/actions: `Approve task`.
- Policies connect events to commands: `When X happens and condition Y holds, do Z`.
- Hotspots are risks, conflicts, confusion, manual work, delays, or pain.
- Do not invent stakeholders. If the user has no answer, record an open question.
