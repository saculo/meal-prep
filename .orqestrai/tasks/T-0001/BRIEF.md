---
id: T-0001
status: approved
---

# Bring DESIGN.md and the discovery in line with the decisions of 2026-11-23

**Category:** chore (documents only; no code exists yet)

**Triage check:** Not built: the documents say the opposite today (see RESEARCH.md). Q060/C004
forbid pieces while DESIGN.md still lists `pcs` and R010 still has "grams plus pieces";
DESIGN.md uses NATS events and an event-fed copy of recipe data in planning. Out-of-scope
match: none (no out-of-scope files, no archive).

**Current behavior:**
- Discovery: Q060 requires mass or volume for every amount; C004 resolved for Q060 (no pieces).
  R010 and DESIGN.md still mention pieces, so the documents contradict each other.
- Discovery: Q054 says a recipe has no portions; Q059/Q061 have the recipe declare a basis
  (per portion or per whole batch, batches planned whole); Q004 requires portions on a planned
  meal. DESIGN.md gives the recipe "portions" and the ingredient a single "unit".
- DESIGN.md: services talk over gRPC and NATS JetStream; planning keeps a copy of recipe names
  and ingredients fed by `recipes.*` events; outbox, durable consumers, swappable broker; NATS
  in Data, Testing and First steps.
- DESIGN.md: five apps each with its own deployment and image, but the image rule names only
  services and only recipes has a Dockerfile in the layout.

**Desired behavior:**
- **Units.** An amount is in g or kg (mass), ml or l (volume), or pieces. g↔kg and ml↔l
  convert; mass, volume and pieces of one ingredient never combine (200 g tomato + 1 tomato
  stays two amounts on the shopping list).
- **Ingredient units.** A reusable ingredient has one or more units it can be used in (an egg
  in pieces or in g/kg); a recipe line picks one of them. Calories are optional per 100 g, per
  100 ml and per piece; a line whose unit has no calorie value shows nutrition as incomplete.
- **Portions.** A recipe may state how many portions its amounts make; a planned meal may state
  its portions. Missing portions count as 1. The shopping list scales each line by meal
  portions ÷ recipe portions. Examples: Chili, 400 g beans, 4 portions, planned for 2 →
  200 g. Omelette, 3 eggs, no portions, planned for 2 → 6 eggs. Chili planned without
  portions → 100 g. The per-portion/per-batch basis and whole batches are gone.
- **Services talk only by asking.** No NATS, events or outbox until a feature needs them
  (DESIGN.md says so in one sentence). Services ask each other over gRPC and never copy
  another service's data. Planning asks recipes for the recipes that match (filters, selected
  ingredients) and sorts them by its own eaten history; shopping asks planning for planned
  meals and recipes for their lines.
- **Five apps.** recipes, planning, shopping, gateway and web, each its own deployment and
  its own image.
- **Records.** The 2026-11-23 answers are new user-supplied corpus entries; Q060, Q054,
  Q059/Q061 and the "portions required" part of Q004 are marked superseded, not deleted;
  C004 is reopened and resolved for Q043 + the new entries. all-event-storms/ stays the
  evidence; DESIGN.md is the technical design citing it. No document keeps a reference to
  the old rules as current.

**Who sees it:** whoever reads DESIGN.md and the discovery next: the user and every later task
(plan steps, implementers) that builds the services from these documents.

**Scope:**
- DESIGN.md: services table, rules, how services talk, data, testing, deployment, first steps.
- all-event-storms/: corpus.md, contradictions.md, problem-brief.md, work-candidates.md
  (at least R003, R010; any other that states the old rules).
- GLOSSARY.md (new): Unit, Portions.

**Out of scope:**
- Any code, service folder, Dockerfile, CI or cluster change.
- The domain-discovery/ toolkit.
- Designing events for later.

**Open decisions:** none.

**Decisions:**
- Units are exactly g, kg, ml, l and pieces (from the request; closes Q060's "etc.").
- Shopping asks planning for planned meals and recipes for their lines, over gRPC (follows from "only ask, never copy").
- Q1 Scope: DESIGN.md and all-event-storms/ both (see Q6).
- Q2 Units per ingredient: an ingredient can be used in different units (an egg as pieces or as g/kg). Recording (decided by the agent, reversible): new corpus entries, old ones marked superseded, C004 reopened.
- Q3 Piece calories: (a) optional calories per piece; a piece line without it shows nutrition as incomplete.
- Q4 Ingredient units and portions: an ingredient has one or more units; a recipe has portions; a meal has portions (as a whole meal).
- Q5 Events later: one sentence in DESIGN.md, no swappable-broker detail.
- Q6 Merge: (a) update both so they agree; all-event-storms/ the evidence, DESIGN.md the technical design citing it.
- Q7 Portions: recipe portions optional; a planned meal's portions scale the recipe (400 g for 4, planned 2 -> 200 g). Q059/Q061 basis and whole batches go; Q054 superseded.
- Q8 Missing portions: count as 1, on the recipe and on the planned meal.

**Assumptions:**
- Recipe deletion is not in the discovery; with no events there is nothing to cascade, and this task adds no deletion rule.
- Pieces never convert to g even when an ingredient has both (2 eggs + 100 g egg stays two amounts).
