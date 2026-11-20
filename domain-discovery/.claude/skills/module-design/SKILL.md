---
name: module-design
description: Convert a validated or candidate context map into module blueprints and craftpath-ready module work items. Keeps module design evidence-backed and separates domain responsibilities from implementation hypotheses.
---

# Module Design

## Purpose

Turn bounded context candidates into implementation-facing module blueprints without writing code.

The output should be ready to feed into `/craftpath-work` one module capability at a time.

## Inputs

Prefer these files when available:

- `discovery/<slug>/context-map.md`
- `discovery/<slug>/context-candidates.md`
- `discovery/<slug>/work-candidates.md`
- `discovery/<slug>/corpus.md`
- existing code/module notes, only if explicitly supplied

## Output files

- `discovery/<slug>/modules.md`
- `discovery/<slug>/module-work.md`

## Module blueprint format

```md
# Modules — <topic>

## `<module-name>` — <public name>

- source context: <bounded context>
- responsibility: <one sentence>
- readiness: ready|needs-questions|too-large|hypothesis
- evidence: Q001, Q004

### Owns

- <entities/value objects/concepts/decisions/lifecycle>

### Does not own

- <explicit exclusions>

### Public entry points / capabilities

- `<operation>(...)` — <business intent, not technical implementation>

### Publishes

- `<DomainEvent>` — <when emitted, why>

### Consumes

- `<DomainEvent>` / external input — <how it reacts>

### Queries needed from other modules

- <module> → <needed information>

### Invariants / rules

- <rule protected by this module>

### Integration boundaries

- <sync call, async event, ACL, shared kernel, open host service>

### Open questions

- <question or None>

### Implementation hypotheses, not requirements

- <optional notes; mark as hypothesis>
```

## Module work format

```md
# Module Work Candidates — <topic>

## M001 — <imperative title>

- module: `<module-name>`
- source context: <context>
- readiness: ready|needs-questions|too-large
- suggested command: `/craftpath-work "<title>"`
- suggested skills: [backend|frontend|testing|architecture|infrastructure|ux]

### Problem

<2-4 sentences focused on observable module capability>

### Scenarios

Scenario: <name>
  Given <precondition>
  When <action/event>
  Then <observable outcome>

### Module boundary notes

- owns: <what this work should keep inside the module>
- collaborates with: <other module/event/query>
- out of scope: <what not to implement>

### Evidence

- Q001 — <reason>

### Open questions

- <question or None>
```

## Design rules

- Module names should follow context language, not database/table names.
- Public entry points describe business capabilities, not framework endpoints.
- Events are past-tense facts: `MealPlanGenerated`, not `GenerateMealPlan`.
- Commands/operations are intent: `GenerateMealPlan`.
- Keep shared concepts explicit: published language, shared kernel, or duplicated local model.
- Avoid premature microservices. A module can be an internal package/component.
- If one context yields many unrelated responsibilities, split the module work or flag the context boundary as suspect.
- If multiple contexts need the same concept, decide whether it is published language, shared kernel, or duplicated with translation.

## Readiness rules

`ready` means:

- context responsibility is clear
- owned concepts and non-owned concepts are explicit
- at least one operation/event scenario is testable
- blocking relationship questions are resolved

`needs-questions` means useful but blocked by one or more boundary/API/event questions.

`too-large` means a module or work item contains multiple independent capabilities and must be split.

`hypothesis` means the module is plausible but weakly supported by evidence.
