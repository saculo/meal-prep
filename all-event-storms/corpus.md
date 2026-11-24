# Evidence Corpus

## discovery/weekly-meal-planning/event-storm.md

### Q001
> We want to plan dinners (but not for week only), it can be for few days also.

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R001 evidence U001
- tags: [goal, actor, event]
- notes: Meal planning range is flexible, not fixed to a calendar week.

### Q002
> All meals which we want, but mostly dinners

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R001 evidence U003
- tags: [rule, term]
- notes: Meal type selection must support more than dinner, while dinner is the common case.

### Q003
> Can do from week view, from checking specific dates

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R001 evidence U006
- tags: [goal, event]
- notes: Planning can begin from a week view or selected dates.

### Q004
> Date, meal type, portions, recipe

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R003 evidence U010
- tags: [rule, term]
- notes: These fields are needed to place a planned meal. Portions "required" superseded by Q068/Q069 (2026-11-23): portions are optional on the planned meal, missing = 1.

### Q005
> Meals are placed as planned

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R003 evidence U002
- tags: [event, term]
- notes: Weekly planning creates planned meals.

### Q006
> Yeah

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R003 evidence U009
- tags: [rule]
- notes: Answer confirming multiple planned meals can exist on the same day.

### Q007
> Sure

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R003 evidence U015
- tags: [rule]
- notes: Answer confirming planned meals can be moved/rescheduled.

### Q008
> Browse/search recipes and give suggestions

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R002 evidence U007
- tags: [goal]
- notes: Planning includes recipe browsing/search and suggestions.

### Q009
> Filters, suggestions

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R002 evidence U008
- tags: [goal]
- notes: Filters are part of recipe selection.

### Q010
> Yeah

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R002 evidence U012
- tags: [rule]
- notes: Answer confirming never-eaten first, then oldest-last-eaten ordering.

### Q011
> No results + offer to relax filters or offer to give random recipe

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R004 evidence U014
- tags: [rule, risk]
- notes: No-result planning path should offer recovery actions.

### Q012
> Warn, but allow

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R004 evidence U016
- tags: [rule]
- notes: Repeating a recipe should warn but not block.

### Q013
> At the moment warn or dont do anything

- source: discovery/weekly-meal-planning/event-storm.md
- loc: Work seeds / R004 evidence U017
- tags: [rule, constraint]
- notes: Calorie warning behavior is non-blocking and may be deferred.

## discovery/what-should-we-eat/event-storm.md

### Q014
> We need to find a dinner for today

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R001 evidence U001
- tags: [pain, goal, event]
- notes: Today dinner decision flow.

### Q015
> Manually

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R001 evidence U002
- tags: [constraint, rule]
- notes: Available ingredients are manually selected, not pantry-derived.

### Q016
> Any selected

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R001 evidence U005
- tags: [rule]
- notes: Ingredient matching starts with any selected ingredient.

### Q017
> Sort by oldest last eaten

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R002 evidence U003
- tags: [rule]
- notes: Meal-history recency controls sorting.

### Q018
> Show first

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R002 evidence U006
- tags: [rule]
- notes: Answer confirming never-eaten recipes are shown first.

### Q019
> Name, last eaten date, matching ingredients

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R003 evidence U007
- tags: [goal, term]
- notes: Minimum visible data for decision-focused result cards.

### Q020
> Randomise a meal

- source: discovery/what-should-we-eat/event-storm.md
- loc: Work seeds / R004 evidence U003
- tags: [goal, risk]
- notes: Random fallback exists, but exact scope is unclear.

## discovery/recipe-capture-to-calendar/event-storm.md

### Q021
> With my wife we want to add a recipe which we made recently to keep it for future

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R001 evidence U001
- tags: [actor, goal, pain]
- notes: Household wants to keep recently cooked recipes.

### Q022
> Name, ingredients, some tags

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R001 evidence U003
- tags: [rule, term]
- notes: Minimal recipe capture data.

### Q023
> I think it can be temporary/free-text if I want, but with option to be added

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R002 evidence U004
- tags: [rule, goal]
- notes: Recipe capture may use temporary/free-text ingredients with later conversion.

### Q024
> Name, calories, macros optional, type optional

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R002 evidence U004
- tags: [rule, constraint]
- notes: New/free-text ingredient data during recipe capture.

### Q025
> It should be possible in both ways, sometimes adding without calendar, something adding to calendar + base

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R003 evidence U002
- tags: [rule, goal]
- notes: Saving a recipe and adding it to calendar are separable.

### Q026
> date, portions

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R003 evidence U010
- tags: [rule]
- notes: Calendar add needs date and portions.

### Q027
> at the moment lets consider that adding to the calendar means eaten always

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Timeline #10 policy/rule
- tags: [rule, risk]
- notes: One scoped flow treats calendar entries as eaten.

### Q028
> categories maybe (as some colours for example) + names

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R004 evidence U011
- tags: [goal]
- notes: Week view should show names and category/color cues.

### Q029
> yeah

- source: discovery/recipe-capture-to-calendar/event-storm.md
- loc: Work seeds / R004 evidence U011
- tags: [goal]
- notes: Answer confirming last-eaten/not-eaten-recently signal in planning view.

## discovery/ingredient-management/event-storm.md

### Q030
> We need an ingredient while creating a recipe

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R001 evidence U001
- tags: [pain, event]
- notes: Ingredient creation can be triggered from recipe creation.

### Q031
> Ingredient saved and reusable

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R001 evidence U002
- tags: [goal, event]
- notes: Ingredient becomes reusable.

### Q032
> Yes

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R001 evidence U009
- tags: [rule]
- notes: Answer confirming saved ingredient immediately fills into the recipe being created.

### Q033
> Give a tip that something exists similar and show that

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R002 evidence U005
- tags: [goal, rule]
- notes: Similar ingredients should be suggested, not blocking.

### Q034
> Can be temporary just for given recipe usage

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R003 evidence U007
- tags: [rule, term]
- notes: Temporary ingredient can be scoped to a recipe usage.

### Q035
> Name + calories

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R003 evidence U007
- tags: [rule, constraint]
- notes: Temporary ingredient was initially described with name plus calories.

### Q036
> Yes and recipe cna be recalculated (meal in history of eaten nope)

- source: discovery/ingredient-management/event-storm.md
- loc: Work seeds / R004 evidence U010
- tags: [rule, constraint]
- notes: Ingredient edits can recalculate recipes, but not eaten meal history.

### Q037
> Ingredient can be created as reusable or temporary for current recipe usage

- source: discovery/ingredient-management/event-storm.md
- loc: Timeline #2 policy/rule
- tags: [rule]
- notes: Ingredient creation mode supports reusable and temporary.

### Q038
> calories were initially described as required, but later ingredient can be saved without calories

- source: discovery/ingredient-management/event-storm.md
- loc: Timeline #6 policy/rule
- tags: [risk, rule]
- notes: Event storm records unresolved calorie requirement contradiction.

## discovery/shopping-list-from-planned-meals/event-storm.md

### Q039
> When I have some planned meat for a week or a day and want to generate shopping list from that

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U002
- tags: [pain, goal, event]
- notes: Shopping list starts from planned meals for a chosen range.

### Q040
> User selected range

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U006
- tags: [rule]
- notes: User chooses date range.

### Q041
> Planned meals only

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U010
- tags: [rule, constraint]
- notes: Shopping-list generation excludes eaten/other meal states.

### Q042
> Yeah, amount of them

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U007
- tags: [goal]
- notes: Shopping list should include ingredient amounts.

### Q043
> If it is 200g tomato + 1 tomato, just combine it with both amounts, if 100ml + 0.5l, just calculate that

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U011
- tags: [rule]
- notes: Compatible units are converted; incompatible amount expressions remain visible.

### Q044
> Add it

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U014
- tags: [rule]
- notes: Temporary/incomplete ingredients should still appear in the shopping list.

### Q045
> Yeah, scale it

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R001 evidence U017
- tags: [rule]
- notes: Planned portions scale ingredient quantities.

### Q046
> Name, amount

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R002 evidence U008
- tags: [goal]
- notes: Shopping list item display needs name and amount.

### Q047
> Manually checked off atm

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R002 evidence U012
- tags: [constraint, rule]
- notes: Already-at-home exclusion is manual.

### Q048
> Yeah

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R002 evidence U013
- tags: [goal]
- notes: Answer confirming generated shopping list can be edited.

### Q049
> Removed when shared

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R002/R003 evidence U016
- tags: [rule]
- notes: Checked-off items should not be included in shared output.

### Q050
> Atm generic copy/share

- source: discovery/shopping-list-from-planned-meals/event-storm.md
- loc: Work seeds / R003 evidence U009
- tags: [constraint]
- notes: Specific communicator/list-app integrations are not required initially.

## user-supplied

### Q051
> Yes, meal should be planned or eaten.

- source: user-supplied answer, current run
- loc: n/a
- tags: [rule]
- notes: Calendar meal state supports planned/eaten distinction.

### Q052
> Yes, they can be

- source: user-supplied answer, current run
- loc: n/a
- tags: [rule]
- notes: Answer to whether reusable and temporary ingredients can be saved without calories.

### Q053
> From base of receipts

- source: user-supplied answer, current run
- loc: n/a
- tags: [rule]
- notes: Random fallback should choose from the recipe/receipt base.

### Q054
> But recipe doesn't need portions, rather list needs how many "portions" are needed rn to be bought

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, term]
- notes: Recipe has no base yield. Portions belong to the shopping need (planned meal portions), not to the recipe definition. How recipe amounts relate to one portion is still open. Superseded by Q068 (2026-11-23): portions are optional on the recipe and on the planned meal; the shopping list scales by meal portions ÷ recipe portions.

### Q055
> Yes, during making planned meal eaten you can modify values if something has changed

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, event]
- notes: Marking a planned meal eaten creates the eaten record from the plan, and its values (e.g. portions, date) can be adjusted at that moment.

### Q056
> exactly

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule]
- notes: Confirms the shopping list is a snapshot at generation; rescheduling meals afterwards does not change an already generated list.

### Q057
> lets not recalculate them at the moment, act like ingredients are not saved with receipt, they are only used for creation

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, constraint]
- notes: Ingredients are an input when a recipe is created; the recipe keeps its own computed values and is not recalculated when an ingredient changes later. Supersedes the "recipe can be recalculated" part of Q036 for now. Eaten history therefore needs no separate snapshot policy.

### Q058
> It should have calories per 100g and optionally per portion

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, goal]
- notes: Corrected by Q063: calories per 100 g is an ingredient value (nutrition basis), not a recipe figure. Recipe calories are summed from its lines; per portion is optional.

### Q059
> Can be per portion and whole batch

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, term]
- notes: Recipe ingredient amounts may be defined either per portion or for a whole batch. Interpreted as: the recipe declares its basis, and a planned meal counts in that basis (portions or batches). Superseded by Q068 (2026-11-23): the per-portion / per-batch basis is gone; recipe amounts are for the recipe portions (missing = 1).

### Q060
> Dont care about that, just assume that every ingredient have to have ml/g/kg/l etc. - make it as a requirement

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, constraint]
- notes: Every recipe ingredient amount must be in a mass or volume unit (g, kg, ml, l, …). Piece counts are not allowed. Supersedes the "1 tomato" example in Q043. Superseded by Q065 (2026-11-23): pieces are a unit again; an amount is in g/kg, ml/l or pieces.

### Q061
> yes

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule]
- notes: Confirms Q059 interpretation: the recipe declares its basis (per portion or per batch); a planned meal counts in that basis; per-batch recipes are planned in whole batches. Superseded by Q068 (2026-11-23): no per-batch basis and no whole batches.

### Q062
> no you cant, take it as separate things... you cant say 1ml=1g...

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, constraint]
- notes: Mass and volume are never converted into each other, for any calculation (shopping lists or nutrition).

### Q063
> We don't need total weight, wtf. We need weight only per ingredient, so we can calculate calories and amount per ingredient

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule, term]
- notes: No recipe-level total weight or per-100 g figure. Each recipe line has its own amount; line calories = line amount × the ingredient's calories per 100 g/ml. Recipe calories are the sum of line calories.

### Q064
> No, either measure

- source: user-supplied answer, archetype review
- loc: n/a
- tags: [rule]
- notes: A reusable ingredient may be used by mass in one recipe and by volume in another; it can carry calories per 100 g and/or per 100 ml.

### Q065
> Amounts can be in g/kg, ml/l or pieces. g and kg convert, ml and l convert. Mass, volume and pieces of one ingredient never combine.

- source: user-supplied answer, units review (2026-11-23)
- loc: n/a
- tags: [rule, term, constraint]
- notes: Supersedes Q060 (no pieces). Pieces never convert to g, even for an ingredient that has both (2 eggs + 100 g egg stays two amounts). Extends Q062 (mass and volume never converted) and restores Q043 ("1 tomato").

### Q066
> A reusable ingredient has one or more units (an egg in pieces or g); a recipe line picks one.

- source: user-supplied answer, units review (2026-11-23)
- loc: n/a
- tags: [rule, term]
- notes: Generalises Q064 (one ingredient by mass in one recipe, volume in another) to include pieces.

### Q067
> Calories are optional per 100 g, per 100 ml and per piece.

- source: user-supplied answer, units review (2026-11-23)
- loc: n/a
- tags: [rule]
- notes: Extends Q063 (line calories = amount x calories per 100 g/ml): a line in pieces uses calories per piece. A line whose unit has no calorie value shows nutrition as incomplete (see Q052).

### Q068
> Portions are optional on the recipe and on the planned meal. The shopping list scales by meal portions / recipe portions.

- source: user-supplied answer, units review (2026-11-23)
- loc: n/a
- tags: [rule, term]
- notes: Supersedes Q054, Q059, Q061 and the "portions required" part of Q004. Examples: Chili, 400 g beans, 4 portions, planned for 2 -> 200 g. Omelette, 3 eggs, no portions, planned for 2 -> 6 eggs. Chili planned without portions -> 100 g.

### Q069
> Missing portions count as 1.

- source: user-supplied answer, units review (2026-11-23)
- loc: n/a
- tags: [rule]
- notes: Applies to the recipe and to the planned meal. Chili (4 portions) planned without portions = 1 portion -> 400 g x 1 / 4 = 100 g.
