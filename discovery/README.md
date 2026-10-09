# Meal Prep Discovery

## Current implementation basis

| Read this for | File |
|---|---|
| Module ownership and boundaries | [modules.md](modules.md) |
| Multi-process decision and Go/Rust split | [design.md](design.md) |
| Build groups and delivery order | [work-groups.md](work-groups.md) |
| Candidate scenarios M001–M016 | [module-work.md](module-work.md) |
| Bounded-context relationships | [context-map.md](context-map.md) |
| Archetype reasoning | [domain-archetypes.md](domain-archetypes.md) |

## Evidence and earlier discovery

- `all-event-storms/` — combined evidence corpus, contradictions, questions,
  problem brief, and early work candidates.
- `ingredient-management/`, `recipe-capture-to-calendar/`,
  `shopping-list-from-planned-meals/`, `weekly-meal-planning/`, and
  `what-should-we-eat/` — focused Event Storming evidence.
- `meal-diary-domain-discovery/` — earlier consolidated analysis retained for
  traceability.

The current plan is `modules.md`, `design.md`, `work-groups.md`, and
`module-work.md`; the other files are supporting evidence, not competing plans.
