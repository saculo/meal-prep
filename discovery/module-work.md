# Module Work Candidates — Meal Prep

> **Architecture amendment required:** These candidates were written for
> packages in one application. `design.md` adopts independently runnable
> processes instead: Go for `recipe-finder`, Rust for `shopping`, and HTTP/JSON
> contracts with isolated data stores. Amend each affected candidate with its
> concrete contract, timeout and unavailable behaviour before implementation.
> Split M001 into the quantity-policy fixtures plus Gleam and Rust conformance
> implementations.

Ordered for building: each item depends only on items above it. Module
blueprints are in `modules.md`; archetype reasoning is in `domain-archetypes.md`.

| # | Title | Module | Readiness |
|---|---|---|---|
| M001 | Add and scale ingredient amounts | `quantity` | ready |
| M002 | Create reusable ingredients with optional calories | `ingredients` | ready |
| M003 | Suggest similar ingredients without blocking | `ingredients` | ready |
| M004 | Capture a recipe from reusable ingredients | `recipes` | ready |
| M005 | Add temporary ingredient lines to a recipe | `recipes` | ready |
| M006 | Calculate recipe calories from its lines | `recipes` | ready |
| M007 | Plan meals on the calendar | `meal-calendar` | ready |
| M008 | Reschedule a planned meal | `meal-calendar` | ready |
| M009 | Mark a planned meal as eaten | `meal-calendar` | ready |
| M010 | Record a meal eaten without a plan | `meal-calendar` | ready |
| M011 | Show recency on the calendar | `meal-calendar` | ready |
| M012 | Find recipes with filters, ordered by recency | `recipe-finder` | ready |
| M013 | Find tonight's dinner by available ingredients | `recipe-finder` | ready |
| M014 | Recover from no results | `recipe-finder` | ready |
| M015 | Generate a shopping list from planned meals | `shopping` | ready |
| M016 | Edit, check off and share a shopping list | `shopping` | ready |

## M001 — Add and scale ingredient amounts

- module: `quantity`
- source context: shared kernel
- readiness: ready
- suggested command: `/craftpath-work "Add and scale ingredient amounts"`
- suggested skills: [backend, testing]

### Problem

Recipes and shopping lists both work with ingredient amounts. Amounts must always carry a mass or volume unit, combine correctly within a unit family, and never mix grams with millilitres. This shared rule is used by every later module.

### Scenarios

Scenario: Combine volumes in different units
  Given an amount of 100 ml and an amount of 0.5 l
  When they are added
  Then the result is 0.6 l

Scenario: Combine masses in different units
  Given an amount of 200 g and an amount of 1 kg
  When they are added
  Then the result is 1.2 kg

Scenario: Keep mass and volume apart
  Given an amount of 200 g and an amount of 100 ml
  When they are added
  Then the result shows both 200 g and 100 ml, unconverted

Scenario: Scale by a planned count
  Given an amount of 150 g
  When it is scaled by 3
  Then the result is 450 g

Scenario: Reject a unit that is not mass or volume
  Given an amount entered as 1 piece
  When it is created
  Then it is rejected

### Module boundary notes

- owns: units, dimensions, conversion within a dimension, scaling, mixed amounts
- collaborates with: none
- out of scope: calories, ingredient identity, display formatting

### Evidence

- Q043 — 100 ml + 0.5 l is calculated
- Q045 — amounts scale with planned portions
- Q060 — every amount is in g/kg/ml/l
- Q062 — 1 ml is never treated as 1 g

### Open questions

- None

## M002 — Create reusable ingredients with optional calories

- module: `ingredients`
- source context: Ingredient Management
- readiness: ready
- suggested command: `/craftpath-work "Create reusable ingredients with optional calories"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

While creating a recipe, a household user may need an ingredient that doesn't exist yet. They create it once and reuse it later. Calories are often unknown at that moment, so they are optional, and they may be known per 100 g, per 100 ml, or both.

### Scenarios

Scenario: Create an ingredient and use it immediately
  Given a user is creating a recipe
  And "oat milk" is not in the ingredient catalogue
  When they create "oat milk" with 45 kcal per 100 ml
  Then "oat milk" is available in the recipe being created
  And it is available for future recipes

Scenario: Create an ingredient without calories
  Given a user creates "smoked paprika" with no calorie values
  When they save it
  Then the ingredient is saved
  And it shows that its calories are missing

Scenario: Store calories for both measures
  Given a user creates "honey" with 304 kcal per 100 g and 420 kcal per 100 ml
  When they save it
  Then both values are kept

### Module boundary notes

- owns: reusable ingredient, its optional calories and macros
- collaborates with: `recipes` calls this module during recipe creation
- out of scope: similar-ingredient tips (M003), pantry stock, recalculating recipes

### Evidence

- Q024 — name, calories, macros optional
- Q030 — ingredient needed while creating a recipe
- Q031 — ingredient saved and reusable
- Q032 — saved ingredient fills into the recipe immediately
- Q052 — ingredients can be saved without calories
- Q064 — an ingredient can be measured by mass or volume

### Open questions

- None

## M003 — Suggest similar ingredients without blocking

- module: `ingredients`
- source context: Ingredient Management
- readiness: ready
- suggested command: `/craftpath-work "Suggest similar ingredients without blocking"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

Duplicate ingredients make the catalogue harder to use. When a user types a new ingredient name, the app should show similar existing ones as a tip, but never stop them creating the new one.

### Scenarios

Scenario: Show similar ingredients
  Given the catalogue contains "chicken breast"
  When a user starts creating "chicken"
  Then "chicken breast" is shown as a similar ingredient

Scenario: Create despite a similar ingredient
  Given "chicken breast" is shown as similar
  When the user still creates "chicken"
  Then "chicken" is saved

### Module boundary notes

- owns: similarity policy
- collaborates with: none
- out of scope: merging duplicates, fuzzy spelling correction

### Evidence

- Q033 — give a tip that something similar exists and show it

### Open questions

- None

## M004 — Capture a recipe from reusable ingredients

- module: `recipes`
- source context: Recipe Collection
- readiness: ready
- suggested command: `/craftpath-work "Capture a recipe from reusable ingredients"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

The household wants to save recipes they cooked recently, before they forget them. A recipe has a name, tags, meal types, and ingredient lines with amounts. The amounts are written either per portion or for a whole batch. Ingredient data is copied into the recipe when it is created, and the recipe never changes when an ingredient changes later.

### Scenarios

Scenario: Save a recipe without putting it on the calendar
  Given a user enters "Bolognese" with tags "pasta" and "family", meal type dinner
  And lines "spaghetti 100 g" and "passata 150 ml" written per portion
  When they save it
  Then "Bolognese" is in the recipe base
  And nothing is added to the calendar

Scenario: Save a recipe written for a whole batch
  Given a user enters "Chilli" with amounts for the whole pot
  When they save it with basis "per batch"
  Then the recipe records that its amounts are per batch

Scenario: Later ingredient edits do not change the recipe
  Given "Bolognese" was saved using "passata" at 30 kcal per 100 ml
  When "passata" is later changed to 35 kcal per 100 ml
  Then "Bolognese" still uses 30 kcal per 100 ml

Scenario: Reject a line without a mass or volume unit
  Given a user enters a line "tomato 1 piece"
  When they try to save the recipe
  Then the line is rejected with a request for g, kg, ml or l

### Module boundary notes

- owns: recipe, basis, lines copied from ingredients
- collaborates with: `ingredients` through local port `IngredientCatalog` (find, create)
- out of scope: temporary lines (M005), calories (M006), editing saved recipes, adding to calendar (M010)

### Evidence

- Q021 — save recently made recipes for the future
- Q022 — name, ingredients, tags
- Q025 — saving works with or without the calendar
- Q057 — ingredients are only used for creation; no recalculation
- Q059 — amounts per portion or per whole batch
- Q060 — every amount is in g/kg/ml/l

### Open questions

- None

## M005 — Add temporary ingredient lines to a recipe

- module: `recipes`
- source context: Recipe Collection
- readiness: ready
- suggested command: `/craftpath-work "Add temporary ingredient lines to a recipe"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

Capturing a recipe should not stop because an ingredient isn't in the catalogue. The user can type it as a one-off line for this recipe only. Later, if they want, they can save it as a reusable ingredient.

### Scenarios

Scenario: Save a recipe with a temporary line
  Given a user is capturing a recipe
  And "za'atar" is not in the catalogue
  When they add a temporary line "za'atar 5 g" with no calories
  Then the recipe saves with that line
  And the ingredient catalogue is unchanged

Scenario: Save a temporary line as a reusable ingredient
  Given a recipe has a temporary line "za'atar 5 g"
  When the user chooses to save it as reusable
  Then "za'atar" is created in the ingredient catalogue

### Module boundary notes

- owns: temporary lines inside a recipe
- collaborates with: `ingredients.CreateIngredient` when saving as reusable
- out of scope: replacing temporary lines in other recipes

### Evidence

- Q023 — temporary/free-text, with option to be added
- Q034 — temporary just for this recipe's use
- Q037 — reusable or temporary at creation

### Open questions

- None

## M006 — Calculate recipe calories from its lines

- module: `recipes`
- source context: Recipe Collection
- readiness: ready
- suggested command: `/craftpath-work "Calculate recipe calories from its lines"`
- suggested skills: [backend, frontend, testing]

### Problem

Each recipe line contributes calories: its amount times the ingredient's calories per 100 g (for grams) or per 100 ml (for millilitres). The recipe total is the sum. When a value is missing, the total is shown as incomplete rather than as a precise number. A per-portion figure is available when the recipe is written per portion.

### Scenarios

Scenario: Sum calories over mass and volume lines
  Given a per-portion recipe with "spaghetti 100 g" at 350 kcal per 100 g
  And "passata 150 ml" at 30 kcal per 100 ml
  When its calories are calculated
  Then the line calories are 350 and 45
  And the recipe total is 395 kcal per portion

Scenario: Mark the total incomplete when a value is missing
  Given a recipe line "olive oil 10 ml"
  And olive oil has calories per 100 g but none per 100 ml
  When the recipe's calories are calculated
  Then that line has no calories
  And the recipe total is marked incomplete

Scenario: No per-portion figure for a batch recipe
  Given a recipe written per batch
  When its calories are shown
  Then the whole-batch total is shown
  And no per-portion figure is shown

### Module boundary notes

- owns: line calories, total, incomplete marker, per-portion figure
- collaborates with: none (values were copied in M004/M005)
- out of scope: macros, calorie warnings during planning, recipe total weight

### Evidence

- Q052 — missing calories are allowed and shown as incomplete
- Q058 — per portion is optional
- Q062 — no ml↔g conversion
- Q063 — calories per ingredient line; no recipe total weight
- Q064 — an ingredient may have per-100 g and/or per-100 ml values

### Open questions

- None

## M007 — Plan meals on the calendar

- module: `meal-calendar`
- source context: Meal Calendar
- readiness: ready
- suggested command: `/craftpath-work "Plan meals on the calendar"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

The household plans meals for a few days or a week, mostly dinners. A planned meal is a recipe on a date, for a meal type, with a count. The count is in portions or whole batches, matching how the recipe is written. Repeats are allowed, with a warning.

### Scenarios

Scenario: Plan a per-portion recipe
  Given "Bolognese" is written per portion
  When a user plans it for Friday dinner, 3 portions
  Then Friday shows a planned dinner "Bolognese" for 3 portions

Scenario: Plan a per-batch recipe in whole batches
  Given "Chilli" is written per batch
  When a user plans it for Sunday lunch, 1 batch
  Then Sunday shows a planned lunch "Chilli", 1 batch

Scenario: Several meals on one day
  Given Friday already has a planned dinner
  When a user plans a lunch on Friday
  Then Friday shows both meals

Scenario: Warn but allow a repeat
  Given "Bolognese" is planned for Monday
  When a user plans "Bolognese" for Wednesday in the same view
  Then they see a repeat warning
  And the meal is still planned

Scenario: View selected dates
  Given meals are planned across two weeks
  When a user views Thursday to Sunday
  Then only meals in that range are shown

### Module boundary notes

- owns: planned meals, repeat-warning policy, range views
- collaborates with: `recipes` through local port `RecipeCatalog` (name, basis, meal types)
- out of scope: rescheduling (M008), eaten meals (M009/M010), recency (M011), calorie warnings

### Evidence

- Q001 — plan for a few days, not only a week
- Q002 — all meal types, mostly dinners
- Q003 — from week view or selected dates
- Q004 — date, meal type, portions, recipe
- Q005 — meals are placed as planned
- Q006 — several meals per day
- Q012 — warn but allow repeats
- Q059 — portions or batches, matching the recipe
- Q061 — per-batch recipes are planned in whole batches

### Open questions

- None

## M008 — Reschedule a planned meal

- module: `meal-calendar`
- source context: Meal Calendar
- readiness: ready
- suggested command: `/craftpath-work "Reschedule a planned meal"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

Plans change. A planned meal can be moved to another date or meal type without re-entering it.

### Scenarios

Scenario: Move a planned meal to another day
  Given "Bolognese" is planned for Friday dinner
  When a user moves it to Saturday dinner
  Then it shows on Saturday and not on Friday

Scenario: Rescheduling does not touch an existing shopping list
  Given a shopping list was generated including Friday's "Bolognese"
  When "Bolognese" is moved to the following week
  Then the existing shopping list is unchanged

### Module boundary notes

- owns: moving planned meals
- collaborates with: none
- out of scope: moving eaten meals

### Evidence

- Q007 — planned meals can be moved
- Q056 — a generated shopping list is a snapshot

### Open questions

- None

## M009 — Mark a planned meal as eaten

- module: `meal-calendar`
- source context: Meal Calendar
- readiness: ready
- suggested command: `/craftpath-work "Mark a planned meal as eaten"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

When a planned meal is eaten, the household records it. The eaten record starts from the plan, and the date or count can be changed if things went differently. It keeps the recipe name and calories it was recorded with. After this the plan counts as fulfilled and no longer appears as planned.

### Scenarios

Scenario: Mark eaten as planned
  Given "Bolognese" is planned for Friday dinner, 3 portions
  When a user marks it eaten
  Then an eaten record exists for Friday, 3 portions, linked to the plan
  And the plan no longer counts as planned

Scenario: Adjust values when marking eaten
  Given "Bolognese" is planned for Friday, 3 portions
  When a user marks it eaten on Saturday with 2 portions
  Then the eaten record shows Saturday and 2 portions

Scenario: Eaten record keeps its values
  Given "Bolognese" was eaten on Friday at 395 kcal per portion
  When the recipe's details change later
  Then Friday's eaten record still shows 395 kcal per portion

### Module boundary notes

- owns: eaten meals, the link to the plan, copied name and calories
- collaborates with: `recipes.GetRecipeCalories` when recording
- out of scope: correcting or deleting eaten records, daily calorie totals

### Evidence

- Q036 — eaten history is never recalculated
- Q051 — a meal is planned or eaten
- Q055 — values can be modified when marking eaten
- Q057 — no recalculation from ingredients

### Open questions

- None

## M010 — Record a meal eaten without a plan

- module: `meal-calendar`
- source context: Meal Calendar
- readiness: ready
- suggested command: `/craftpath-work "Record a meal eaten without a plan"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

Often the household cooks something unplanned, then saves the recipe and wants it on the calendar as eaten. This works either straight after capturing a recipe or from an existing recipe.

### Scenarios

Scenario: Capture a recipe and record it as eaten
  Given a user is capturing "Shakshuka"
  When they save it and choose to add it to the calendar for yesterday, 2 portions
  Then "Shakshuka" is in the recipe base
  And yesterday shows an eaten "Shakshuka" for 2 portions with no linked plan

Scenario: Record an existing recipe as eaten
  Given "Bolognese" is in the recipe base and not planned
  When a user records it eaten today, 3 portions
  Then today shows an eaten "Bolognese" for 3 portions

### Module boundary notes

- owns: unplanned eaten meals
- collaborates with: `recipes` (the capture flow is application-level: capture, then record)
- out of scope: changing recipe capture itself

### Evidence

- Q021 — save recently made recipes
- Q025 — add to base and calendar together
- Q026 — calendar add needs date and portions
- Q027 — adding from capture means eaten

### Open questions

- None

## M011 — Show recency on the calendar

- module: `meal-calendar`
- source context: Meal Calendar (history)
- readiness: ready
- suggested command: `/craftpath-work "Show recency on the calendar"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

While planning, the household wants variety. The calendar shows each meal's recipe name, a category cue, and when that recipe was last eaten, or that it hasn't been eaten recently. Recency comes only from eaten records.

### Scenarios

Scenario: Show last eaten date
  Given "Bolognese" was last eaten on 1 October
  When a user views a week with "Bolognese" planned
  Then the meal shows last eaten 1 October

Scenario: Show never eaten
  Given "Shakshuka" has no eaten records
  When it appears on the calendar
  Then it shows as never eaten

Scenario: Planned meals do not count as eaten
  Given "Chilli" is planned but never eaten
  When its recency is shown
  Then it shows as never eaten

Scenario: Show a category cue
  Given "Bolognese" is tagged "pasta"
  When it appears on the calendar
  Then it shows the "pasta" category cue

### Module boundary notes

- owns: recency projection, calendar view composition
- collaborates with: `recipes` for names and categories
- out of scope: suggestion ordering (M012)

### Evidence

- Q010 — never eaten first, then oldest last eaten
- Q017 — sort by oldest last eaten
- Q018 — never eaten shown first
- Q028 — names and category cues
- Q029 — last eaten or not eaten recently signal

### Open questions

- None

## M012 — Find recipes with filters, ordered by recency

- module: `recipe-finder`
- source context: Recipe Suggestions
- readiness: ready
- suggested command: `/craftpath-work "Find recipes with filters, ordered by recency"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

When planning, the household browses recipes with filters such as meal type and tags. Results put never-eaten recipes first, then the ones eaten longest ago. Each card shows the name and last-eaten date.

### Scenarios

Scenario: Order by recency
  Given "Shakshuka" was never eaten
  And "Chilli" was last eaten in August and "Bolognese" in September
  When a user searches dinners
  Then the order is "Shakshuka", "Chilli", "Bolognese"

Scenario: Filter by tag
  Given recipes tagged "pasta" and "soup"
  When a user filters by "pasta"
  Then only "pasta" recipes are shown

Scenario: Card shows decision data
  Given "Chilli" was last eaten on 12 August
  When it appears in results
  Then its card shows "Chilli" and 12 August

### Module boundary notes

- owns: composed criteria, ordering, result cards
- collaborates with: `recipes` (local port `RecipeSource`), `meal-calendar` (local port `RecencySource`)
- out of scope: ingredient matching (M013), empty results (M014), placing on the calendar

### Evidence

- Q008 — browse/search recipes with suggestions
- Q009 — filters
- Q010 — never eaten first, then oldest
- Q017 — sort by oldest last eaten
- Q018 — never eaten first
- Q019 — name, last eaten date, matching ingredients

### Open questions

- None

## M013 — Find tonight's dinner by available ingredients

- module: `recipe-finder`
- source context: Recipe Suggestions
- readiness: ready
- suggested command: `/craftpath-work "Find tonight's dinner by available ingredients"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

To decide tonight's dinner, a user picks the ingredients they have. Any recipe containing at least one of them is shown, with the matching ingredients on its card and the same recency ordering.

### Scenarios

Scenario: Match any selected ingredient
  Given "Bolognese" uses spaghetti and passata
  And "Shakshuka" uses eggs and passata
  When a user selects "eggs"
  Then "Shakshuka" is shown and "Bolognese" is not

Scenario: Show which ingredients matched
  Given a user selects "eggs" and "passata"
  When "Shakshuka" is shown
  Then its card lists eggs and passata as matching

### Module boundary notes

- owns: any-selected match rule
- collaborates with: same ports as M012
- out of scope: pantry stock, all-selected matching

### Evidence

- Q014 — find a dinner for today
- Q015 — ingredients are selected manually
- Q016 — any selected
- Q019 — matching ingredients on the card

### Open questions

- None

## M014 — Recover from no results

- module: `recipe-finder`
- source context: Recipe Suggestions
- readiness: ready
- suggested command: `/craftpath-work "Recover from no results"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

A search can come back empty. The user is then offered two ways forward: relax the filters, or get a random recipe from the whole base.

### Scenarios

Scenario: Offer to relax filters
  Given no recipe matches tag "soup" and ingredient "eggs"
  When the search returns nothing
  Then the user is offered searches with those filters removed

Scenario: Random recipe from the whole base
  Given a search returned nothing
  When the user asks for a random recipe
  Then one recipe from the whole base is shown, regardless of filters

### Module boundary notes

- owns: relaxing criteria, random pick
- collaborates with: `recipes` (local port `RecipeSource`)
- out of scope: weighting the random pick

### Evidence

- Q011 — no results: relax filters or random recipe
- Q020 — randomise a meal
- Q053 — random from the recipe base

### Open questions

- None

## M015 — Generate a shopping list from planned meals

- module: `shopping`
- source context: Shopping Preparation
- readiness: ready
- suggested command: `/craftpath-work "Generate a shopping list from planned meals"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

For a chosen range, the household wants one list of what to buy for the meals still planned. Each recipe line is multiplied by the planned count. The same ingredient from different recipes is combined, converting units within mass or within volume. Once generated, the list is stored and is not changed by later calendar edits.

### Scenarios

Scenario: Scale by planned count and combine
  Given "Bolognese" (per portion, spaghetti 100 g) is planned for 3 portions on Friday
  And "Carbonara" (per portion, spaghetti 120 g) is planned for 2 portions on Saturday
  When a user generates a list for Friday to Saturday
  Then the list shows spaghetti 540 g

Scenario: Per-batch recipe
  Given "Chilli" (per batch, beans 800 g) is planned for 1 batch
  When a list is generated for that day
  Then the list shows beans 800 g

Scenario: Convert within volume
  Given planned recipes need passata 150 ml and passata 0.5 l
  When a list is generated
  Then the list shows passata 0.65 l

Scenario: Keep mass and volume side by side
  Given planned recipes need milk 200 g and milk 100 ml
  When a list is generated
  Then the list shows milk 200 g + 100 ml

Scenario: Only meals still planned
  Given Friday's "Bolognese" was already marked eaten
  When a list is generated for that week
  Then "Bolognese" ingredients are not included

Scenario: Include temporary and incomplete ingredients
  Given a planned recipe has a temporary line "za'atar 5 g" with no calories
  When a list is generated
  Then za'atar 5 g is on the list

### Module boundary notes

- owns: generation rule, combining, the stored list
- collaborates with: `meal-calendar` (local port `PlannedMealSource`), `recipes` (local port `RecipeLineSource`), `quantity`
- out of scope: editing and sharing (M016), pantry stock, regenerating on calendar change

### Evidence

- Q039 — generate from planned meals
- Q040 — user-selected range
- Q041 — planned meals only
- Q042 — with amounts
- Q043 — combine units
- Q044 — include temporary/incomplete ingredients
- Q045 — scale by portions
- Q056 — the list is a snapshot
- Q059 — per portion or per batch
- Q062 — no ml↔g conversion
- Q064 — the same ingredient may appear by mass and by volume

### Open questions

- None

## M016 — Edit, check off and share a shopping list

- module: `shopping`
- source context: Shopping Preparation
- readiness: ready
- suggested command: `/craftpath-work "Edit, check off and share a shopping list"`
- suggested skills: [backend, frontend, testing, ux]

### Problem

A generated list needs a human pass. Users edit items, tick off what is already at home, and share the rest as plain text through the phone's normal copy/share.

### Scenarios

Scenario: Edit an item
  Given the list shows spaghetti 540 g
  When a user changes it to 500 g
  Then the list shows spaghetti 500 g

Scenario: Check off an item at home
  Given the list shows passata 0.65 l
  When a user checks it off
  Then it is marked as already at home

Scenario: Share without checked-off items
  Given passata is checked off and spaghetti is not
  When the user shares the list
  Then the shared text contains spaghetti
  And it does not contain passata

### Module boundary notes

- owns: item edits, check-off state, share text
- collaborates with: OS copy/share target
- out of scope: integrations with specific apps, pantry stock

### Evidence

- Q046 — name and amount
- Q047 — manual check-off
- Q048 — the list can be edited
- Q049 — checked-off items are removed when shared
- Q050 — generic copy/share

### Open questions

- None
