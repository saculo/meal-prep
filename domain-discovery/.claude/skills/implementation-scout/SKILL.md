---
name: implementation-scout
description: Optional post-discovery scouting skill inspired by DD-AI 004 and 005. Looks for likely algorithmic or coupling hotspots before craftpath planning, but keeps findings as hypotheses separate from requirements.
---

# Implementation Scout

## Purpose

Use after candidate requirements exist, especially before `/craftpath-work`, to flag implementation risks without turning them into requirements.

## Algorithm hotspot scan

Look for semantic signals:

- ordering with dependencies -> topological sort / cycle detection
- shortest/best route -> path finding
- matching pairs/resources -> matching / assignment
- scheduling scarce resources -> constraint solving / interval overlap
- ranking/searching/filtering -> indexing / scoring
- deduplication/idempotency -> hashing / set membership
- state transitions -> finite state machine

Report:

```md
### Algorithm hypothesis
- candidate: R00X
- pattern: <name>
- why: <domain signals>
- acceptance implication: <edge case worth testing>
```

## Coupling hotspot scan

Inspired by connascence, look for cross-boundary agreements:

- shared names/types
- shared magic values or meanings
- positional argument agreements
- ordering/execution dependencies
- timing assumptions
- identity/value consistency across modules

Report:

```md
### Boundary hypothesis
- candidate: R00X
- seam: <module/context A> <-> <module/context B>
- risk: <agreement or coupling>
- planning implication: <maybe design task / split / test level>
```

## Rules

- Do not require source-code analysis unless paths are provided.
- Mark every finding as hypothesis unless backed by concrete source refs.
- Never put implementation scout content into the requirement scenarios directly.
- Use findings to suggest craftpath skills or design tasks.
