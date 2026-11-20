# Work Candidates

## R001 — Start meal planning for a flexible range

- readiness: ready
- class: workflow/orchestration, CRUD/state bookkeeping
- suggested skills: [frontend, backend, testing, ux]
- suggested command: `/craftpath-work "Start meal planning for a flexible range"`

### Problem

The household needs to plan meals for a week or only a few selected dates. Planning should begin from either week view or selected dates and establish the meal type before recipe suggestions.

### Scenarios

Scenario: Choose planning range and meal type
  Given a household user wants to plan meals
  When they open a week view or select specific dates
  Then they can choose a planning range and meal type for suggestions

### Out of scope

- Shopping-list generation
- Marking meals eaten or skipped
- Calorie warning rules

### Assumptions

- Dinner is the primary meal type, but the model should not exclude other meal types.

### Open questions

- None

### Evidence

- Q001 — Flexible planning range
- Q002 — Meal types include mostly dinner but not only dinner
- Q003 — Week view or selected dates entry points

### Classification notes

- why: The candidate starts a multi-step planning workflow and records user-selected state.
- planning risk: low
- likely design needed: no

## R002 — Suggest recipes using filters and meal recency

- readiness: ready
- class: transformation/presentation
- suggested skills: [backend, frontend, testing, ux]
- suggested command: `/craftpath-work "Suggest recipes using filters and meal recency"`

### Problem

Meal planning should support variety and avoiding repeats. Suggestions should use selected filters and order recipes so never-eaten recipes appear before recipes sorted by oldest last eaten.

### Scenarios

Scenario: Order planning suggestions by recency
  Given a household user selected a meal type and filters
  When recipe suggestions are shown
  Then never-eaten recipes appear first
  And previously eaten recipes are ordered by oldest last eaten

### Out of scope

- Random fallback when no recipes match
- Calorie warning thresholds
- Visual card design beyond required data

### Assumptions

- Missing meal history means never eaten.

### Open questions

- None

### Evidence

- Q008 — Browse/search and suggestions are needed
- Q009 — Filters and suggestions are part of planning
- Q010 — Never-eaten first and oldest-last-eaten ordering confirmed

### Classification notes

- why: The main behavior transforms recipes and history into an ordered presentation list.
- planning risk: medium
- likely design needed: maybe

## R003 — Place and reschedule planned meals

- readiness: ready
- class: CRUD/state bookkeeping, workflow/orchestration
- suggested skills: [backend, frontend, testing]
- suggested command: `/craftpath-work "Place and reschedule planned meals"`

### Problem

A selected recipe must be placed on the calendar as a planned meal with date, meal type, portions, and recipe. Users should be able to place multiple meals on one day and move/reschedule planned meals later.

### Scenarios

Scenario: Place a planned meal
  Given a household user selected a recipe for planning
  When they place it on the calendar with date, meal type, portions, and recipe
  Then the meal is recorded as planned

Scenario: Reschedule a planned meal
  Given a meal is recorded as planned
  When the household user moves it to another date or meal slot
  Then the planned meal appears in the new place

### Out of scope

- Marking a meal eaten or skipped
- Shopping-list generation
- Recipe capture

### Assumptions

- Calendar entries can represent planned meals.

### Open questions

- None

### Evidence

- Q004 — Planned meal data fields
- Q005 — Meals are placed as planned
- Q006 — Multiple planned meals per day allowed
- Q007 — Moving/rescheduling allowed
- Q027 — Earlier eaten-only calendar statement
- Q051 — Calendar meals should be planned or eaten

### Classification notes

- why: The candidate stores and changes planned meal lifecycle state.
- planning risk: medium
- likely design needed: maybe

## R004 — Find dinner recipes by manually selected available ingredients

- readiness: ready
- class: transformation/presentation
- suggested skills: [backend, frontend, testing, ux]
- suggested command: `/craftpath-work "Find dinner recipes by manually selected available ingredients"`

### Problem

The household needs a quick way to find dinner options using ingredients they currently have without requiring pantry tracking. The first matching rule is simple: show recipes containing any selected ingredient.

### Scenarios

Scenario: Search dinner by available ingredients
  Given a household user needs dinner for today
  And they manually select available ingredients
  When they search for dinner recipes
  Then recipes containing any selected ingredient are shown

### Out of scope

- Automatic pantry inventory
- Any/all/mostly matching modes
- Adding the chosen recipe to calendar

### Assumptions

- Selected ingredients are matched against recipe ingredients.

### Open questions

- None

### Evidence

- Q014 — Dinner for today need
- Q015 — Ingredients are selected manually
- Q016 — Any selected ingredient matching rule

### Classification notes

- why: The behavior filters and presents recipe options from selected criteria.
- planning risk: low
- likely design needed: no

## R005 — Capture a recently cooked recipe into the recipe collection

- readiness: ready
- class: CRUD/state bookkeeping, workflow/orchestration
- suggested skills: [backend, frontend, testing, ux]
- suggested command: `/craftpath-work "Capture a recently cooked recipe into the recipe collection"`

### Problem

The household wants to save interesting recently cooked meals before they are forgotten. Capture should save the recipe with enough information to be useful later while keeping calendar addition optional.

### Scenarios

Scenario: Save a recipe without calendar placement
  Given a household user remembers a recently cooked meal
  When they enter recipe name, ingredients, and tags
  Then the recipe is saved in the shared recipe collection

Scenario: Keep calendar addition optional
  Given a recipe was captured
  When the user does not choose calendar addition
  Then the recipe remains available for future planning

### Out of scope

- Meal eaten-state modeling
- Full ingredient nutrition validation
- Photo/link/source import

### Assumptions

- Instructions and macros can be optional for initial capture.

### Open questions

- None for saving the recipe itself.

### Evidence

- Q021 — Household wants to save a recently made recipe
- Q022 — Name, ingredients, tags
- Q025 — Save with or without calendar

### Classification notes

- why: The candidate creates reusable recipe records through a capture workflow.
- planning risk: low
- likely design needed: no

## R006 — Support temporary/free-text ingredients during recipe capture

- readiness: ready
- class: CRUD/state bookkeeping, safety/guard/evidence
- suggested skills: [backend, frontend, testing, ux]
- suggested command: `/craftpath-work "Support temporary/free-text ingredients during recipe capture"`

### Problem

Recipe capture should not stop when an ingredient is not yet in the reusable ingredient database. Users need to add temporary/free-text ingredients and optionally convert them later. Ingredients can be saved without calories, so incomplete nutrition must remain visible rather than blocking recipe capture.

### Scenarios

Scenario: Add a temporary ingredient to a recipe
  Given a user is capturing a recipe
  And the ingredient is not in the reusable ingredient collection
  When they enter it as temporary/free-text
  Then the recipe can still be saved with that ingredient

Scenario: Flag incomplete nutrition
  Given a temporary/free-text ingredient has incomplete nutrition data
  When recipe nutrition is shown
  Then the missing data is observable rather than silently treated as complete

### Out of scope

- Automatic ingredient import
- Fuzzy ingredient matching
- Full conversion workflow to reusable ingredient

### Assumptions

- Temporary/free-text ingredients are scoped to the recipe unless later converted.

### Open questions

- None

### Evidence

- Q023 — Temporary/free-text ingredient with option to add later
- Q024 — Name/calories/macros/type evidence
- Q034 — Temporary for given recipe usage
- Q038 — Earlier calories-required contradiction recorded
- Q052 — Ingredients can be saved without calories

### Classification notes

- why: The candidate creates incomplete state and must guard against misleading nutrition evidence.
- planning risk: high
- likely design needed: maybe

## R007 — Create reusable ingredients during recipe creation

- readiness: ready
- class: CRUD/state bookkeeping, integration
- suggested skills: [backend, frontend, testing, ux]
- suggested command: `/craftpath-work "Create reusable ingredients during recipe creation"`

### Problem

While creating a recipe, a user may discover a needed ingredient is missing. They should be able to create a reusable ingredient from the recipe flow and immediately use it in the recipe.

### Scenarios

Scenario: Create and use a reusable ingredient from recipe creation
  Given a household user is creating a recipe
  And the needed ingredient does not exist
  When they create and save a reusable ingredient
  Then the ingredient becomes available in the current recipe and future recipes

Scenario: Show similar ingredient suggestions
  Given a user enters a new ingredient name
  When similar ingredients already exist
  Then the system shows those similar ingredients as suggestions without blocking creation

### Out of scope

- Ingredient source/import integration
- Pantry stock tracking
- Recalculating historical eaten meals

### Assumptions

- Similar suggestions can start with simple name matching.

### Open questions

- None

### Evidence

- Q030 — Ingredient needed during recipe creation
- Q031 — Ingredient saved and reusable
- Q032 — Immediately fills into recipe
- Q033 — Similar ingredient tip/list
- Q038 — Earlier calories-required contradiction recorded
- Q052 — Ingredients can be saved without calories

### Classification notes

- why: This creates reusable ingredient state and connects the ingredient flow back into recipe creation.
- planning risk: medium
- likely design needed: maybe

## R008 — Show week-view planning signals

- readiness: ready
- class: transformation/presentation
- suggested skills: [frontend, backend, ux, testing]
- suggested command: `/craftpath-work "Show week-view planning signals"`

### Problem

The week view should support planning decisions, not only list calendar items. It should show meal names, category/color cues, and last-eaten/not-eaten-recently signals.

### Scenarios

Scenario: Display planning signals in week view
  Given recipes have been added to calendar days
  When a household user views the week
  Then the view shows meal names, category cues, and last-eaten or not-eaten-recently information

### Out of scope

- Detailed visual design
- Marking meals eaten/skipped
- Shopping-list generation

### Assumptions

- Last-eaten data is available from meal history.

### Open questions

- None

### Evidence

- Q028 — Names and category/color cues
- Q029 — Last-eaten/not-eaten-recently signal confirmed
- Q005 — Planned meals in planning flow
- Q027 — Earlier eaten-only calendar statement
- Q051 — Calendar meals should be planned or eaten

### Classification notes

- why: The candidate presents existing calendar/recipe/history data for decisions.
- planning risk: medium
- likely design needed: maybe

## R009 — Handle no-results and random fallback for recipe choice

- readiness: ready
- class: workflow/orchestration, transformation/presentation
- suggested skills: [backend, frontend, testing, ux]
- suggested command: `/craftpath-work "Handle no-results and random fallback for recipe choice"`

### Problem

Planning and dinner choice should not dead-end when no recipes match. Users should be offered a way to relax filters or get a random recipe from the recipe base.

### Scenarios

Scenario: Offer recovery when no planning recipes match
  Given a household user is planning meals with selected criteria
  When no recipes match
  Then the app shows no results
  And offers to relax filters or get a random recipe

Scenario: Offer random fallback for dinner
  Given no dinner recipes match the selected criteria
  When the user asks for a fallback
  Then the application offers a random meal option

### Out of scope

- Defining random weighting/scoring
- Calorie warning thresholds

### Assumptions

- Random fallback should at least return a recipe the user can open.

### Open questions

- None

### Evidence

- Q011 — No results plus relax filters/random recipe
- Q020 — Randomise a meal
- Q053 — Random fallback from recipe/receipt base

### Classification notes

- why: This candidate orchestrates fallback paths and transforms an empty result into useful options.
- planning risk: medium
- likely design needed: maybe

## R010 — Generate a shopping list from planned meals

- readiness: ready
- class: transformation/presentation, workflow/orchestration
- suggested skills: [backend, frontend, testing]
- suggested command: `/craftpath-work "Generate a shopping list from planned meals"`

### Problem

Before shopping, a household user needs an ingredient list based on planned meals in a selected date range. The list should scale quantities by planned portions, combine same ingredients when possible, convert compatible units, keep incompatible amounts visible, and include temporary/incomplete ingredients.

### Scenarios

Scenario: Generate shopping list for selected range
  Given a household user has planned meals in a selected range
  When they generate a shopping list
  Then the list includes ingredients from planned meals only
  And ingredient quantities are scaled by planned meal portions
  And same ingredients are combined by amount where possible
  And temporary or incomplete ingredients are still shown

Scenario: Represent compatible and incompatible units
  Given planned recipes include 100ml and 0.5l of the same ingredient
  When the shopping list is generated
  Then those compatible amounts are calculated together
  And incompatible amounts such as grams plus pieces remain visible together

### Out of scope

- Pantry automation
- Fuzzy/synonym ingredient matching
- Specific third-party sharing integrations

### Assumptions

- Planned meals have portion counts and recipe ingredient quantities.

### Open questions

- None

### Evidence

- Q039 — Shopping list from planned meals
- Q040 — User-selected range
- Q041 — Planned meals only
- Q043 — Unit conversion and mixed amounts
- Q044 — Include incomplete ingredients
- Q045 — Scale by portions
- Q051 — Calendar meals should be planned or eaten

### Classification notes

- why: The candidate transforms planned meals and recipe quantities into an aggregated list.
- planning risk: high
- likely design needed: yes

## R011 — Review, edit, and generically share a shopping list

- readiness: ready
- class: CRUD/state bookkeeping, transformation/presentation
- suggested skills: [frontend, backend, testing, ux]
- suggested command: `/craftpath-work "Review, edit, and generically share a shopping list"`

### Problem

After generation, the shopping list must be usable in real shopping. Users need to see names and amounts, manually check off items already at home, edit list items, and copy/share the list without checked-off items.

### Scenarios

Scenario: Review and edit generated list
  Given a shopping list was generated
  When the household user reviews it
  Then each item shows name and amount
  And the user can manually check off already-at-home items
  And the user can edit the list

Scenario: Share list without checked-off items
  Given a shopping list has checked-off items
  When the household user copies or shares it
  Then checked-off items are omitted from the shared output

### Out of scope

- App-specific communicator/list integrations
- Automatic pantry inventory
- Regenerating the list after recipe changes

### Assumptions

- Generic copy/share can be represented as plain text initially.

### Open questions

- Exact text format for mixed or missing amounts can be refined later; it does not block basic copy/share.

### Evidence

- Q046 — Items show name and amount
- Q047 — Already-at-home manually checked off
- Q048 — List can be edited
- Q049 — Checked-off removed when shared
- Q050 — Generic copy/share initially

### Classification notes

- why: The candidate maintains editable shopping-list item state and presents/share transforms it.
- planning risk: low
- likely design needed: no
