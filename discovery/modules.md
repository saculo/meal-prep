# Modules — Meal Prep

## Module map

```text
quantity (shared kernel) ← used by every module below

ingredients  ←  recipes  ←  meal-calendar  ←  recipe-finder
                   ↑             ↑
                   └── shopping ─┘
```

- Dependencies point one way only; there are no cycles.
- `quantity` is a versioned, language-neutral policy with conformance fixtures;
  it is implemented locally by its consumers, not deployed as a remote service.
- Each module exposes commands, queries, views and (when needed) events; its
  aggregates stay internal.
- Modules are independently runnable processes with separate data stores and
  versioned HTTP/JSON contracts. `ingredients`, `recipes`, and `meal-calendar`
  are Gleam; `recipe-finder` is Go; `shopping` is Rust. See `design.md`.
- Archetype fits and reasoning: `domain-archetypes.md`.

Changes from the previous version:
- `meal-history` is folded into `meal-calendar`.
- `recipe-suggestions` is renamed `recipe-finder`.
- `recipe-collection` is renamed `recipes`, and `ingredient-management` is renamed `ingredients`.
- `shopping-preparation` is renamed `shopping`.
- `quantity` is new.

## `quantity` — Quantity

- source context: shared kernel (Recipe Collection, Shopping Preparation)
- responsibility: Represent ingredient amounts as a number with a mass or volume unit, and do the only arithmetic the domain allows on them.
- archetypes: Quantity (strong) — see `domain-archetypes.md`
- readiness: ready
- evidence: Q043, Q045, Q060, Q062

### Owns

- `Unit` (g, kg, ml, l) and its `Dimension` (mass, volume)
- `Quantity` = amount + unit; non-negative
- Conversion inside one dimension (g↔kg, ml↔l)
- Scaling by a whole count
- `MixedAmount`: a mass part and a volume part kept side by side

### Does not own

- Ingredient identity or names
- Calories or any nutrition
- Mass↔volume conversion, which is forbidden [Q062]
- Piece or count units, which are not allowed [Q060]

### Public entry points / capabilities

- `Quantity.of(amount, unit)` — create an amount; rejects units outside mass/volume.
- `add(a, b)` — add two amounts; converts within a dimension, keeps mass and volume apart as a `MixedAmount`.
- `scale(amount, count)` — multiply an amount by a planned count.

### Publishes

- None (value objects only).

### Consumes

- None.

### Queries needed from other modules

- None.

### Invariants / rules

- Every amount has a mass or volume unit. [Q060]
- Compatible units are combined into one amount: 100 ml + 0.5 l = 0.6 l. [Q043]
- Mass and volume are never converted into each other. [Q062]

### Integration boundaries

- Shared kernel used by `ingredients`, `recipes`, `meal-calendar` and `shopping`. Changes need agreement from every user.

### Open questions

- None

### Implementation hypotheses, not requirements

- Display normalisation, e.g. 1500 g shown as 1.5 kg, is a presentation choice, not a kernel rule.

## `ingredients` — Ingredients

- source context: Ingredient Management
- responsibility: Keep a catalogue of reusable ingredients with their optional calorie values, so recipes can be built from them.
- archetypes: Product (partial) — see `domain-archetypes.md`
- readiness: ready
- evidence: Q024, Q030, Q031, Q032, Q033, Q037, Q052, Q057, Q064

### Owns

- Reusable ingredient: name
- Calories per 100 g and/or per 100 ml, each optional [Q052, Q064]
- Macros, optional [Q024]
- Policy for suggesting similar ingredients

### Does not own

- Recipes, or which recipes use an ingredient
- Temporary ingredients, which are free-text lines inside one recipe
- Pantry stock [Q015, Q047]
- Recalculating anything after an edit [Q057]

### Public entry points / capabilities

- `CreateIngredient(name, caloriesPer100g?, caloriesPer100ml?, macros?)` — add a reusable ingredient and return it, so the recipe being created can use it immediately [Q031, Q032].
- `FindIngredients(text)` — look up ingredients by name while building a recipe.
- `FindSimilarIngredients(name)` — return existing ingredients that look alike, as a tip [Q033].

### Publishes

- None needed. Recipes copy ingredient data at creation time, so nothing listens for changes [Q057].

### Consumes

- None.

### Queries needed from other modules

- None.

### Invariants / rules

- An ingredient can be saved with no calories at all. [Q052]
- An ingredient may carry calories per 100 g, per 100 ml, or both. [Q064]
- Similar ingredients are a tip and never block creation. [Q033]

### Integration boundaries

- `recipes` queries it synchronously while a recipe is being created and copies what it needs.

### Open questions

- None

### Implementation hypotheses, not requirements

- Similarity can start as case-insensitive name containment.

## `recipes` — Recipes

- source context: Recipe Collection
- responsibility: Keep reusable recipes, meaning what goes into a meal and how much, together with calories calculated when the recipe is created.
- archetypes: Product (partial), Pricing for calories (partial) — see `domain-archetypes.md`
- readiness: ready
- evidence: Q002, Q021, Q022, Q023, Q025, Q028, Q034, Q037, Q052, Q057, Q058, Q059, Q060, Q062, Q063, Q064

### Owns

- Recipe: name, tags/categories, meal types [Q002, Q022, Q028]
- Amount basis: **per portion** or **per batch** [Q059]
- Recipe lines: ingredient name + `Quantity`, plus the calories per 100 g / 100 ml copied from the ingredient (or typed in for a temporary line)
- Temporary lines: free-text ingredients that exist only in this recipe [Q023, Q034]
- Calorie calculation: per line, whole recipe, and per portion when the basis is per portion

### Does not own

- The reusable ingredient catalogue
- Dated meals, planned or eaten
- Shopping lists
- Recency or ordering of recipes

### Public entry points / capabilities

- `CaptureRecipe(name, tags, mealTypes, basis, lines)` — save a recipe to the household base, with or without putting it on the calendar [Q021, Q025].
- `AddIngredientLine(recipeDraft, ingredient, quantity)` — copy a reusable ingredient into a recipe line [Q057].
- `AddTemporaryLine(recipeDraft, name, quantity, calories?)` — add a free-text line without touching the catalogue [Q023, Q034].
- `SaveTemporaryLineAsIngredient(line)` — create a reusable ingredient from a temporary line for future recipes [Q023, Q037].
- `GetRecipeSummaries(ids | all)` — name, tags, meal types, line names; used by the calendar views and the finder.
- `GetRecipeLines(id)` — basis + lines with quantities; used by shopping.
- `GetRecipeCalories(id)` — whole recipe, per portion if available, and an incomplete marker.

### Publishes

- None needed in the MVP. Every consumer queries when it needs data.

### Consumes

- None.

### Queries needed from other modules

- `ingredients` → find and create ingredients during recipe creation (local port `IngredientCatalog`).

### Invariants / rules

- Saving a recipe never requires a calendar entry. [Q025]
- Every line amount is mass or volume. [Q060]
- Line calories = amount × the matching calories (per 100 g for mass, per 100 ml for volume). No conversion between mass and volume. [Q062, Q063, Q064]
- If the matching value is missing, the line and the total are marked incomplete. Saving is never blocked. [Q052]
- Recipe calories = sum of line calories. There is no recipe total weight. [Q063]
- Per-portion calories exist only when the basis is per portion. [Q058, Q059]
- Ingredient data is copied when the recipe is created and never recalculated afterwards. [Q057]

### Integration boundaries

- Customer of `ingredients` at creation time only.
- Supplier of summaries, lines and calories to `meal-calendar`, `recipe-finder` and `shopping`.
- "Capture and add to calendar" is a flow in the application layer: `CaptureRecipe`, then `meal-calendar.RecordEatenMeal`. No event is involved.

### Open questions

- None

### Implementation hypotheses, not requirements

- A line may keep the id of the ingredient it was copied from, for display only. Nothing reads through it.
- Editing a saved recipe is not evidenced and stays out of scope until asked for.

## `meal-calendar` — Meal Calendar

- source context: Meal Calendar + Meal History
- responsibility: Plan recipes onto dates and record what was actually eaten, keeping the plan and the eaten record as separate facts, and derive recency from what was eaten.
- archetypes: Plan vs Execution (strong), Accounting (partial) — see `domain-archetypes.md`
- readiness: ready
- evidence: Q001, Q003, Q004, Q005, Q006, Q007, Q010, Q012, Q017, Q018, Q021, Q025, Q027, Q028, Q029, Q036, Q041, Q051, Q055, Q059, Q061

### Owns

- **PlannedMeal**: recipe ref, date, meal type, count (portions or whole batches, matching the recipe's basis) [Q004, Q059, Q061]
- **EatenMeal**: recipe ref, copied recipe name, date eaten, count, copied calories, optional link to the planned meal it fulfils [Q051, Q055]
- Rescheduling planned meals [Q007]
- Repeat-warning policy [Q012]
- Recency projection: last-eaten date and never-eaten status per recipe [Q010, Q017, Q018]
- Calendar views over a chosen range: a week or selected dates [Q001, Q003, Q028, Q029]

### Does not own

- Recipe content or calorie calculation
- Shopping lists
- Search, filters or suggestion ordering
- Calorie warnings, which are deferred [Q013]

### Public entry points / capabilities

- `PlanMeal(recipe, date, mealType, count)` — put a recipe on the calendar as planned [Q004, Q005].
- `RescheduleMeal(plannedMeal, date, mealType?)` — move a planned meal [Q007].
- `MarkPlannedMealEaten(plannedMeal, adjustments?)` — create the eaten record from the plan. Date and count may be changed if things differed [Q055].
- `RecordEatenMeal(recipe, date, count)` — record a meal that was eaten without being planned, e.g. right after capturing a recipe [Q021, Q025, Q027].
- `GetPlannedMeals(range)` — open planned meals in a range; used by shopping [Q040, Q041].
- `GetCalendar(range)` — planned and eaten meals with recipe name, category cue and recency signal [Q028, Q029].
- `GetRecency(recipeIds | all)` — last-eaten date or never eaten [Q010, Q017, Q018].

### Publishes

- None needed outside the module. Internally, recording an eaten meal updates the recency projection.

### Consumes

- None.

### Queries needed from other modules

- `recipes` → names, categories and meal types for views, and calories to copy into an eaten record (local port `RecipeCatalog`).

### Invariants / rules

- A planned meal needs a recipe, date, meal type and count. [Q004]
- A per-batch recipe is planned in whole batches. [Q061]
- Several meals may share a day. [Q006]
- Planning a recipe that is already planned nearby warns but is allowed. [Q012]
- Planned and eaten are separate records. An eaten meal may exist with no plan. [Q025, Q027, Q051]
- Once a planned meal is eaten it is fulfilled. It no longer counts as planned for shopping. [Q041, Q055]
- An eaten meal keeps the name and calories it was recorded with. Later changes elsewhere never rewrite it. [Q036, Q057]
- Recency is derived from eaten meals only. [Q017, Q018]

### Integration boundaries

- Customer of `recipes`.
- Supplier of planned meals to `shopping` and of recency to `recipe-finder`.

### Open questions

- None

### Implementation hypotheses, not requirements

- Meal types can start as a fixed list (breakfast, lunch, dinner, snack), with dinner as the default [Q002].
- The repeat warning can compare against the range currently being viewed.
- Correcting a wrong eaten record can be a superseding entry (Accounting style) rather than an edit, if it ever becomes needed.

## `recipe-finder` — Recipe Finder

- source context: Recipe Suggestions
- responsibility: Help the household pick a recipe, for planning or for tonight, using filters, ingredients they have, and how long since each recipe was eaten.
- archetypes: Rules (partial) — see `domain-archetypes.md`
- readiness: ready
- evidence: Q008, Q009, Q010, Q011, Q014, Q015, Q016, Q017, Q018, Q019, Q020, Q053

### Owns

- The search criteria as a composed filter: meal type, tags, selected ingredients
- Ingredient match rule: any selected ingredient [Q016]
- Ordering: never eaten first, then oldest last eaten [Q010, Q017, Q018]
- Result cards: name, last-eaten date, matching ingredients [Q019]
- No-results recovery: relax filters, or a random recipe from the whole base [Q011, Q020, Q053]

### Does not own

- Any stored state; it only reads
- Recipes, ingredients or eaten history
- Placing the chosen recipe on the calendar, which the application flow does through `meal-calendar`
- Pantry stock; selected ingredients are typed in each time [Q015]

### Public entry points / capabilities

- `FindRecipes(criteria)` — ordered result cards for planning or for tonight [Q008, Q009, Q014].
- `RelaxCriteria(criteria)` — offer criteria with terms removed when nothing matched [Q011].
- `PickRandomRecipe()` — a random recipe from the whole base [Q020, Q053].

### Publishes

- None.

### Consumes

- None.

### Queries needed from other modules

- `recipes` → summaries with line names (local port `RecipeSource`).
- `meal-calendar` → recency per recipe (local port `RecencySource`).

### Invariants / rules

- Selected ingredients match a recipe if any one of them appears in its lines. [Q016]
- Never-eaten recipes come first, then the ones eaten longest ago. [Q010, Q017, Q018]
- An empty result always offers relaxing filters or a random recipe. [Q011]
- Random choice is from all recipes, not the filtered set. [Q053]

### Integration boundaries

- Pure read model over two local ports; it holds no data of its own.

### Open questions

- None

### Implementation hypotheses, not requirements

- Ingredient matching by case-insensitive line name is enough, since recipe lines are copies [Q057].

## `shopping` — Shopping

- source context: Shopping Preparation
- responsibility: Turn the meals planned in a chosen range into a shopping list, then let the household edit it, tick off what they have, and share it.
- archetypes: Ordering (partial), Quantity (strong) — see `domain-archetypes.md`
- readiness: ready
- evidence: Q039, Q040, Q041, Q042, Q043, Q044, Q045, Q046, Q047, Q048, Q049, Q050, Q056, Q059, Q060, Q062, Q064

### Owns

- Shopping list: a stored snapshot created at generation [Q056]
- Generation rule: open planned meals in range × recipe lines × planned count [Q040, Q041, Q045, Q059]
- Combining items: same ingredient name, convert within a dimension, mass and volume side by side [Q043, Q062, Q064]
- Item edits, already-at-home check-off, share text [Q047, Q048, Q049]

### Does not own

- Planned meals, recipes or ingredients
- Pantry stock [Q047]
- Integrations with specific apps [Q050]

### Public entry points / capabilities

- `GenerateShoppingList(range)` — build and store a list from planned meals in the range [Q039, Q040].
- `GetShoppingList(id)` — items with name and amount [Q046].
- `EditItem(list, item, name?, amount?)` — change a generated item [Q048].
- `CheckOffItem(list, item)` / `UncheckItem` — mark what is already at home [Q047].
- `ShareText(list)` — plain text for copy/share, without checked-off items [Q049, Q050].

### Publishes

- None needed.

### Consumes

- None. The list is a snapshot, so later reschedules do not reach it. [Q056]

### Queries needed from other modules

- `meal-calendar` → open planned meals in range with count (local port `PlannedMealSource`).
- `recipes` → basis and lines for each planned recipe (local port `RecipeLineSource`).

### Invariants / rules

- Only planned meals that are not yet eaten are included. [Q041, Q055]
- Each line amount is multiplied by the planned count. [Q045, Q059]
- The same ingredient from different recipes is combined; compatible units are converted (100 ml + 0.5 l = 0.6 l). [Q043]
- Mass and volume amounts of the same ingredient stay side by side. [Q062, Q064]
- Temporary and incomplete ingredients are still listed. [Q044]
- Once generated, the list is not rewritten by calendar changes. [Q056]
- Checked-off items are left out of the share text. [Q049]

### Integration boundaries

- Customer of `meal-calendar` and `recipes`, at generation time only.
- Outbound plain text to the OS copy/share target. [Q050]

### Open questions

- None

### Implementation hypotheses, not requirements

- Items are combined by case-insensitive ingredient name, since recipe lines are copies [Q057].
- The list stays editable after sharing; nothing evidences locking it.
