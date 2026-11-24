# Contradictions

Previously blocking contradictions were clarified by user-supplied answers in this run.

## C001 — Calendar entries as planned vs eaten

### Clashing evidence

> Meals are placed as planned

- source: Q005 — discovery/weekly-meal-planning/event-storm.md

> Planned meals only

- source: Q041 — discovery/shopping-list-from-planned-meals/event-storm.md

> at the moment lets consider that adding to the calendar means eaten always

- source: Q027 — discovery/recipe-capture-to-calendar/event-storm.md

### Why this conflicted

Weekly planning and shopping-list generation depend on a `planned` calendar state, but recipe-capture-to-calendar said adding to calendar means `eaten` always.

### Clarification answer

> Yes, meal should be planned or eaten.

- source: Q051 — user-supplied answer

### Status

Resolved for candidate planning: calendar meals should support planned/eaten state.

## C002 — Ingredient calories required vs optional/incomplete

### Clashing evidence

> Name + calories

- source: Q035 — discovery/ingredient-management/event-storm.md

> Name, calories, macros optional, type optional

- source: Q024 — discovery/recipe-capture-to-calendar/event-storm.md

> calories were initially described as required, but later ingredient can be saved without calories

- source: Q038 — discovery/ingredient-management/event-storm.md

### Why this conflicted

Temporary/free-text and new ingredients were described as requiring calories in some places, while the ingredient-management event storm recorded that ingredients may be saved without calories.

### Clarification answer

> Yes, they can be

- source: Q052 — user-supplied answer

### Status

Resolved for candidate planning: reusable and temporary ingredients can be saved without calories.

## C003 — Recipe recalculation after ingredient edits

### Clashing evidence

> Yes and recipe cna be recalculated (meal in history of eaten nope)

- source: Q036 — discovery/ingredient-management/event-storm.md

> lets not recalculate them at the moment, act like ingredients are not saved with receipt, they are only used for creation

- source: Q057 — user-supplied answer

### Why this conflicted

Q036 allowed recipes to be recalculated from edited ingredients; Q057 defers that and treats ingredients as creation-time input only.

### Status

Resolved for now in favour of Q057: no recalculation. Q036's rule that eaten history never changes still holds and is satisfied trivially.

## C004 — Piece amounts in recipes

### Clashing evidence

> If it is 200g tomato + 1 tomato, just combine it with both amounts, if 100ml + 0.5l, just calculate that

- source: Q043 — discovery/shopping-list-from-planned-meals/event-storm.md

> just assume that every ingredient have to have ml/g/kg/l etc. - make it as a requirement

- source: Q060 — user-supplied answer (superseded by Q065, 2026-11-23)

### Why this conflicted

Q043 assumed piece amounts ("1 tomato") exist; Q060 required every amount to be mass or volume.

### Clarification answer

> Amounts can be in g/kg, ml/l or pieces. g and kg convert, ml and l convert. Mass, volume and pieces of one ingredient never combine.

- source: Q065 — user-supplied answer (2026-11-23)

### Status

Reopened 2026-11-23 and resolved for Q043 + Q065: pieces are a unit again. g with kg and ml with l combine; mass, volume and pieces of one ingredient never combine. Q060 is superseded. (Earlier resolution, now void: Q060, no piece amounts.)

## C005 — Portions: required, per portion or batch, or optional

### Clashing evidence

> But recipe doesn't need portions, rather list needs how many "portions" are needed rn to be bought

- source: Q054 — user-supplied answer (superseded by Q068)

> Can be per portion and whole batch

- source: Q059 — user-supplied answer (superseded by Q068); confirmed by Q061

> Date, meal type, portions, recipe

- source: Q004 — discovery/weekly-meal-planning/event-storm.md (portions required, superseded by Q068/Q069)

### Why this conflicted

Q054 put portions only on the shopping need, Q059/Q061 gave the recipe a per-portion or per-batch basis planned in whole batches, and Q004 required portions on the planned meal.

### Clarification answer

> Portions are optional on the recipe and on the planned meal. The shopping list scales by meal portions / recipe portions.

- source: Q068 — user-supplied answer (2026-11-23)

> Missing portions count as 1.

- source: Q069 — user-supplied answer (2026-11-23)

### Status

Resolved for Q068/Q069: portions are optional on both, missing = 1; scale = meal portions ÷ recipe portions. The per-portion/per-batch basis and whole batches are gone.
