# Problem Brief — Meal Prep Discovery Across Event Storms

## Current pain

- The household wants to keep useful recipes, plan meals across flexible ranges, decide what to eat today, and generate shopping lists without re-entering the same information repeatedly. [Q001, Q014, Q021, Q039]
- Meal decisions should avoid repetition and use meal-history signals such as never eaten and oldest last eaten. [Q017, Q018, Q010]
- Recipe and ingredient capture can be blocked or weakened by missing ingredients, temporary/free-text data, or unclear nutrition completeness rules. [Q023, Q030, Q034, Q038]
- Shopping preparation currently needs a practical list from planned meals, with manual handling for at-home items and sharing. [Q039, Q047, Q050]

## Actors / users

- Either household user / any household user in a shared two-person household. [Q021]
- The household acts as a shared owner of recipes, ingredients, calendar meals, and meal history. [Q001, Q003, Q025]

## Desired observable outcomes

- Users can plan meals for a week or selected dates, choose meal type, place planned meals, and later move them. [Q001, Q002, Q003, Q004, Q007]
- Recipe suggestions support filters and variety by showing never-eaten recipes first and oldest-last-eaten recipes next. [Q008, Q009, Q010]
- A user can find dinner for today by manually selecting available ingredients and seeing recipes that match any selected ingredient. [Q014, Q015, Q016]
- Search/result displays include decision data such as recipe name, last eaten date, and matching ingredients. [Q019]
- Recently cooked recipes can be saved to the recipe collection with name, ingredients, and tags, with optional calendar addition. [Q021, Q022, Q025]
- Missing ingredients can be handled during recipe creation/capture through reusable or temporary/free-text ingredients. [Q023, Q030, Q031, Q034]
- Shopping lists can be generated from planned meals in a selected range, scale quantities by portions, combine compatible units, show incomplete items, allow review/edit/check-off, and generic copy/share. [Q040, Q041, Q043, Q044, Q045, Q046, Q047, Q048, Q049, Q050]

## Rules and constraints

- Planning range is flexible; it may be a week or only selected dates. [Q001, Q003]
- Meal type should support all desired meal types, with dinners as the main use case. [Q002]
- Planned meal placement requires date, meal type, portions, and recipe. [Q004]
- Multiple planned meals per day are allowed. [Q006]
- Repeat recipes should warn but not block. [Q012]
- Calorie warnings should not block and may be deferred. [Q013]
- Available ingredients are manually supplied; automatic pantry tracking is out of current scope. [Q015, Q047]
- Similar ingredient suggestions should warn/show tips but not block creation. [Q033]
- Shopping-list sharing is generic copy/share; app-specific integrations are not initially required. [Q050]

## Risks / hotspots

- Calendar state was clarified: meals should support planned or eaten state, resolving the planning/shopping-list dependency on planned meals. [Q005, Q027, Q041, Q051]
- Ingredient calorie requirement was clarified: reusable and temporary ingredients can be saved without calories, so incomplete nutrition must be visible rather than blocking. [Q024, Q035, Q038, Q052]
- Random fallback was clarified as choosing from the recipe/receipt base. [Q011, Q020, Q053]
- Unit conversion and mixed amount expressions need careful acceptance examples for shopping lists. [Q043]

## Metrics or success signals

- Household can complete weekly/few-day planning and produce planned meals. [Q001, Q005]
- Household can choose dinner faster with visible recency and ingredient-match signals. [Q014, Q019]
- Household can preserve recently cooked recipes before forgetting them. [Q021]
- Household can generate and share an actionable shopping list from planned meals. [Q039, Q050]

## Out of scope / non-goals

- Automatic pantry inventory and stock tracking. [Q015, Q047]
- Specific communicator/list-app integrations for shopping lists. [Q050]
- Blocking calorie validation for planning. [Q013]
- Fuzzy ingredient matching for shopping list combination beyond exact/same ingredients and compatible-unit conversion. [Q043]

## Evidence index

See `corpus.md` Q001–Q050.
