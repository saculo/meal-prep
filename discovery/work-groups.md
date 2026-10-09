# Work Groups — Meal Prep

These groups organise the 16 ready module work candidates into coherent delivery
increments. Complete a group before starting the next one; items within a group
are ordered where one relies on another.

## G1 — Shared quantity foundation

**Outcome:** A single, trusted representation of mass and volume amounts exists
for every later module.

- M001 — Add and scale ingredient amounts (`quantity`)

**Why first:** Recipe capture and shopping-list generation both depend on its
unit validation, conversion, mixed mass/volume handling, and scaling rules.

## G2 — Reusable ingredient catalogue

**Outcome:** The household can create reusable ingredients, including ones with
unknown nutrition, and gets non-blocking duplicate hints.

1. M002 — Create reusable ingredients with optional calories (`ingredients`)
2. M003 — Suggest similar ingredients without blocking (`ingredients`)

**Depends on:** G1 for quantities used later in recipe lines. M003 extends the
catalogue created in M002.

## G3 — Recipe base

**Outcome:** The household can preserve recipes with reusable or one-off lines
and obtain trustworthy, explicitly incomplete calorie totals.

1. M004 — Capture a recipe from reusable ingredients (`recipes`)
2. M005 — Add temporary ingredient lines to a recipe (`recipes`)
3. M006 — Calculate recipe calories from its lines (`recipes`)

**Depends on:** G1 and G2. M004 establishes the saved recipe and copied
ingredient data; M005 adds the alternative line type; M006 calculates from
those stored snapshots.

## G4 — Plan and record meals

**Outcome:** Meals can be planned, moved, recorded as eaten, and shown with
recency information.

1. M007 — Plan meals on the calendar (`meal-calendar`)
2. M008 — Reschedule a planned meal (`meal-calendar`)
3. M009 — Mark a planned meal as eaten (`meal-calendar`)
4. M010 — Record a meal eaten without a plan (`meal-calendar`)
5. M011 — Show recency on the calendar (`meal-calendar`)

**Depends on:** G3, since calendar records reference recipe summaries and copy
recipe calorie values. M008 extends planned meals. M009 and M010 establish
eaten facts; M011 derives recency from those facts.

## G5 — Choose a recipe (Go)

**Outcome:** The Go `recipe-finder` service lets the household browse or find a
suitable recipe, prioritising variety and recovering when filters find nothing.

1. M012 — Find recipes with filters, ordered by recency (`recipe-finder`)
2. M013 — Find tonight's dinner by available ingredients (`recipe-finder`)
3. M014 — Recover from no results (`recipe-finder`)

**Depends on:** G3 for recipe summaries and G4 for recency. M013 adds an
ingredient criterion to M012; M014 handles empty searches from either path.

## G6 — Prepare and share shopping (Rust)

**Outcome:** The Rust `shopping` service lets the household make a stable
shopping-list snapshot from open planned meals, adjust it, check items off, and
share the remaining items.

1. M015 — Generate a shopping list from planned meals (`shopping`)
2. M016 — Edit, check off and share a shopping list (`shopping`)

**Depends on:** G1 for quantity arithmetic, G3 for recipe lines, and G4 for
open planned meals. M016 operates on the saved list created in M015.

## Delivery boundary

- A group is a planning and review boundary, not a single implementation task.
- Each M-item remains independently planned, implemented, tested, and committed.
- Do not start a downstream group by assuming unfinished upstream behaviour.
