---
id: T-0001
status: approved
risk: low
---

# DESIGN.md and the discovery follow the 2026-11-23 decisions

## Open decisions
None.

## Decisions
- Agent: new corpus entries are Q065 units, Q066 ingredient units, Q067 piece calories, Q068 portions, Q069 missing portions = 1 (2026-11-23, user-supplied). The gRPC-only rule is technical and lives in DESIGN.md, not the corpus.
- Agent: C004 reopened and resolved for Q043 + Q065; a new C005 records the portions clash (Q054 vs Q059/Q061 vs Q068).
- Agent: three tasks in one wave, split by file; the Q numbers are fixed here so they can run in parallel.
- Q10 Shopping list snapshot: allowed. The stored list (items with name, amount, unit at generation) is shopping's own record, not a copy kept in sync with recipes or planning.
- Q11 Checks: none needed for now. No grep checks, no verify-surface skill; the user reads the PR.

## Acceptance criteria
- **AC1** Units: the documents state that an amount is in g/kg, ml/l or pieces; g↔kg and ml↔l convert; mass, volume and pieces of one ingredient never combine. No document keeps "no pieces" as a current rule.
- **AC2** Ingredient units: a reusable ingredient has one or more units (an egg in pieces or g); a recipe line picks one. Calories are optional per 100 g, per 100 ml and per piece; a line whose unit has no calorie value shows nutrition as incomplete.
- **AC3** Portions: optional on the recipe and on the planned meal, missing = 1; the shopping list scales by meal portions ÷ recipe portions, with the brief's three examples. The per-portion/per-batch basis and whole batches are gone.
- **AC4** Services: DESIGN.md has services ask each other over gRPC only and never copy data; planning asks recipes for matching recipes and sorts by its own eaten history; shopping asks planning and recipes and stores its own generated list (Q10). No NATS, events, outbox, subjects or listeners remain, except one sentence saying events come when a feature needs them.
- **AC5** Apps: DESIGN.md names five apps (recipes, planning, shopping, gateway, web), each its own deployment and image, with an image name that covers all five.
- **AC6** Records: Q065–Q069 in the corpus; Q060, Q054, Q059, Q061 and Q004's "portions required" marked superseded, not deleted; C004 reopened and resolved; C005 added; problem brief and work candidates (R003, R010 and any other stating an old rule) cite the new entries. GLOSSARY.md has Unit and Portions and the documents use those words.

## Operations
| Operation | Class | Why |
|-----------|-------|-----|
| — | n/a | Documents only; the change adds or changes no operation. |

## Seams
None: no code (Q11).

## Behaviors
None: no code to test (Q11).

## Design
- all-event-storms/ stays the evidence and DESIGN.md the technical design citing it (brief Q6).
- Superseded entries keep their text and get a "superseded by Qxxx (2026-11-23)" note, so the trail of why stays (brief Q2).
- DESIGN.md loses the "Kept swappable" paragraph, NATS in Data, event tests in Testing and NATS in First steps; the services table's "Calls / listens" column becomes "Calls" (brief Q5, AC4).

## Tasks
### Wave 1
- **W1.1** Corpus and contradictions: add Q065–Q069; mark Q004 (portions part), Q054, Q059, Q060, Q061 superseded; reopen C004; add C005. Touches: all-event-storms/corpus.md, all-event-storms/contradictions.md. Serves: AC1, AC2, AC3, AC6.
- **W1.2** Problem brief and work candidates: update rules, risks, out-of-scope line and evidence index; update R003 (portions optional), R010 (scaling with examples, units scenario), R005/R006 if they state units or calories; cite Q065–Q069. Touches: all-event-storms/problem-brief.md, all-event-storms/work-candidates.md. Serves: AC1, AC2, AC3, AC6.
- **W1.3** DESIGN.md: services table, rules, how services talk, data, testing, deployment, first steps, per AC1–AC5; check GLOSSARY.md words. Touches: DESIGN.md, GLOSSARY.md. Serves: AC1–AC6.

## Verification plan
| Check | Claim | Surface | Kind | How (what a user does and sees) | Evidence |
|-------|-------|---------|------|---------------------------------|----------|
| V1 | AC1–AC6 | docs | manual | The user reads the PR diff against BRIEF.md (Q11: no automated checks for now). | PR review |

## Risk
low: documents only, no code or infra. Blast radius: later tasks building from a wrong or inconsistent document. Priority: the first service (recipes) is built from these documents next.

## Assumptions
- Recipe deletion stays undefined; with no events nothing cascades (brief).
- Pieces never convert to g, even for an ingredient that has both (brief).
- Temporary free-text ingredient lines stay as today: listed as written, not scaled or combined.

## Impl notes
- All: AC3's "brief's three examples" are copied into each task prompt (Chili, 400 g beans, 4 portions, planned for 2 -> 200 g; Omelette, 3 eggs, no portions, planned for 2 -> 6 eggs; Chili planned without portions -> 100 g), because the spec names them but does not state them.
- All: the Q065-Q069 topics are given to every task so W1.2 and W1.3 can cite them while W1.1 writes them, because the three run in parallel.
