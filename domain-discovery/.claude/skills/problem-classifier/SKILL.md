---
name: problem-classifier
description: Classify candidate requirements by modeling/implementation problem type and suggest craftpath skills. Adapted from DD-AI 007 for craftpath planning.
---

# Problem Classifier

## Classes

### CRUD / state bookkeeping

Signals:

- create/update/delete/list
- status changes
- validation around stored records
- simple lifecycle tracking

Likely craftpath skills: `backend`, maybe `frontend`.

### Transformation / presentation

Signals:

- render, format, summarize, report, export
- convert between representations
- produce command/prompt/output text

Likely skills: `backend`, `frontend`, `ux`, `testing` when golden tests matter.

### Integration

Signals:

- CLI invokes external command
- harness integration
- Git/GitHub/API/filesystem boundary
- sync between two systems

Likely skills: `backend`, `infrastructure`, `testing`.

### Resource contention / concurrency

Signals:

- locking, ordering, race, duplicate work
- multiple agents/users/processes touching same state
- id allocation, transactionality

Likely skills: `backend`, `architecture`, `testing`.

### Safety / guard / evidence

Signals:

- preventing unsafe edits
- approval gates
- auditability
- verification evidence
- refusing invalid state transitions

Likely skills: `backend`, `testing`, `architecture`.

### Workflow / orchestration

Signals:

- multi-step process
- gates
- dependencies
- resumability
- task lifecycle

Likely skills: `backend`, `testing`, maybe `ux` for command experience.

### Documentation / spec-only

Signals:

- clarify docs/templates
- add examples
- improve generated text
- no behavior change

Likely skills: `testing` only if docs are generated or checked.

## Output per candidate

```md
- class: <primary>[, <secondary>]
- why: <1 sentence>
- suggested skills: [...]
- planning risk: low|medium|high
- likely design needed: yes|no|maybe
```

## Rules

- Classification is a planning aid, not a verdict.
- If classification is unclear, ask one targeted question.
- Use the class to improve scenarios and skill suggestions.
