# Questions

Clarification answers gathered in this run are recorded below and added to `corpus.md` as user-supplied evidence.

## Q001 — Calendar state model

- status: answered
- why: Decides whether calendar meals can be planned separately from eaten history, which affects planning, recipe-to-calendar, and shopping-list inputs.
- answer: "Yes, meal should be planned or eaten."
- affects: [R003, R008, R010]

Question: For MVP, should calendar meals support at least separate `planned` and `eaten` states, or should the recipe-capture flow be treated as a special eaten-history entry separate from planning calendar entries?

## Q002 — Ingredient calories requirement

- status: answered
- why: Decides whether ingredient and recipe capture should block missing calories or allow incomplete nutrition.
- answer: "Yes, they can be"
- affects: [R005, R006]

Question: Can reusable and temporary ingredients be saved without calories, and if so should recipes/shopping lists show them as incomplete rather than blocking save?

## Q003 — Random fallback scope

- status: answered
- why: Decides acceptance criteria for random recipe fallback in planning and dinner decision flows.
- answer: "From base of receipts"
- affects: [R009]

Question: Should random fallback choose from selected meal type only, active filters, partial matches, or all recipes?
