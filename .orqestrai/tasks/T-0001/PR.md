<!-- orqestrai T-0001. Written by ship from the task files in .orqestrai/tasks/T-0001/. -->

## What and why
DESIGN.md and the discovery contradicted each other and the decisions of 2026-11-23:
- The discovery forbade piece amounts (Q060/C004), while DESIGN.md still listed `pcs`.
- DESIGN.md used NATS events, and planning kept an event-fed copy of recipe data.
- Portions were described three different ways.

Now:
- Units are g/kg, ml/l and pieces, and mass, volume and pieces never combine.
- Portions are optional on both the recipe and the planned meal, and missing counts as 1.
- Services only ask each other over gRPC and never copy each other's data.
- There are five apps, each with its own image.

This is for anyone who builds the services from these documents next.

## Acceptance criteria
Documents only. All six ACs are covered by V1, which is manual: your read of this diff (see below). No AC was verified automatically.
- [ ] **AC1** Units: g/kg, ml/l or pieces. g↔kg and ml↔l convert. Mass, volume and pieces never combine. No document keeps "no pieces" as a current rule.
- [ ] **AC2** An ingredient has one or more units, and a recipe line picks one. Calories are optional per 100 g, per 100 ml and per piece. A line whose unit has no calorie value shows nutrition as incomplete.
- [ ] **AC3** Portions are optional on the recipe and the planned meal, missing = 1. The shopping list scales by meal ÷ recipe portions:
  - Chili 400 g for 4 portions, planned for 2 → 200 g.
  - Omelette, 3 eggs, no portions, planned for 2 → 6 eggs.
  - Chili planned without portions → 100 g.
  - The per-portion/per-batch basis is gone.
- [ ] **AC4** gRPC only, and no service copies another's data:
  - Planning asks recipes for the matching recipes and sorts them by its own eaten history.
  - Shopping stores the list it generates.
  - The only mention of events is one sentence: "Events come when a feature needs them; none today."
- [ ] **AC5** Five apps, each with its own deployment and image (`ghcr.io/saculo/meal-prep-<app>:<commit>`).
- [ ] **AC6** Records:
  - Q065–Q069 are added to the corpus.
  - Q004 (its portions part), Q054, Q059, Q060 and Q061 are marked superseded, not deleted.
  - C004 is reopened and resolved, and C005 is added.
  - The problem brief and R003, R006 and R010 cite the new entries.
  - GLOSSARY.md defines Unit and Portions.

## For you to check
- [ ] V1 (AC1–AC6): read the diff against `.orqestrai/tasks/T-0001/BRIEF.md`. There is no app to run.

## Verify
- Rounds: 1.
- Quality gate: not configured. No `AGENTS.md` with `## Checks` exists yet, and you chose no checks for this task (Q11).
- Review: no findings at confidence ≥ 75. No blocking findings, so nothing needed fixing.

## Open findings (not blocking)
- The problem brief's evidence index reads "Q001–Q053 and Q065–Q069". It leaves out entries that are still current (Q055–Q058, Q062–Q064).
- The brief's shopping-list line dropped its Q045 citation, though Q045 was only clarified, not superseded.

## Learn
- **L1** ADR 0001: services ask over gRPC and never copy data, and events come only when a feature needs them. It also records what this costs. `docs/adr/0001-services-ask-over-grpc-and-never-copy.md`
- **L2** A rule for changing a discovery rule: a new corpus entry, the old one marked superseded, the contradiction reopened, and the evidence index kept complete. `AGENTS.md`
- Dropped: a glossary entry "App" vs "service". It would call gateway "not a service" while DESIGN.md lists gateway in its Services table.

## Decisions
- New corpus entries Q065–Q069, dated 2026-11-23. The gRPC-only rule is technical, so it lives in DESIGN.md (and ADR 0001), not in the corpus.
- Shopping's stored list is shopping's own record, not a copy kept in sync (Q10).
- Superseded entries keep their text and get a "Superseded by Qxxx (2026-11-23)" note.

## Risk
Low: documents only. The risk is that later tasks build from an inconsistent document. Recipes, the first service, is built from these next.

🤖 Generated with [Claude Code](https://claude.com/claude-code)
