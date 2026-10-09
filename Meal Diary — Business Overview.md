# Meal Diary — Business Overview

## 1. Product Vision

Meal Diary is a private household application for managing recipes and keeping track of meals eaten over time.

The application is intended primarily for a couple or small household that wants one shared place to:

- store recipes,
- maintain reusable ingredients,
- plan meals,
- record meals that were actually eaten,
- remember when particular meals were last eaten,
- avoid repeating the same meals too frequently,
- quickly decide what to eat next.

The core idea is not simply to create another recipe manager.

The application combines:

```text
Recipe Collection
        +
Meal Planning
        +
Meal History
```

into one system.

The result should answer a very practical question:

> What can we cook, what are we planning to eat, and what have we eaten recently?

---

# 2. Problem

Meal planning in a household usually happens across several disconnected places.

Recipes may be stored in:

- notes,
- websites,
- screenshots,
- bookmarks,
- memory.

Meal planning may happen through:

- verbal discussion,
- calendar notes,
- shopping lists,
- messages.

There is usually no easy way to answer questions such as:

- What meals do we normally cook?
- When did we last eat lasagne?
- What dinners haven't we eaten recently?
- What should we cook this week?
- Which recipes contain a particular ingredient?
- Which meals fit particular nutrition requirements?
- What ingredients do we need for meals planned this week?

Meal Diary aims to centralize this information.

---

# 3. Target Users

The initial target is a **single household with two people**.

Both users share:

- recipes,
- ingredients,
- meal calendar,
- categories,
- tags,
- meal history.

The application is therefore household-oriented rather than individual-oriented.

Future versions may support:

- larger households,
- individual preferences,
- individual portions,
- individual nutrition targets.

These are not required initially.

---

# 4. Main Business Objects

The application revolves around four main concepts.

```text
Ingredient
    ↓
Recipe
    ↓
Planned Meal
    ↓
Meal History
```

## Ingredient

A reusable food item.

Examples:

```text
Chicken breast
Potato
Rice
Mozzarella
Olive oil
Skyr
```

Ingredients contain nutritional information and classification metadata.

---

## Recipe

A reusable definition of a meal.

Examples:

```text
Chicken Curry
Protein Pancakes
Salmon with Potatoes
Lasagne
Tomato Soup
```

Recipes consist of ingredients and their quantities.

A recipe also contains information such as:

- meal types,
- categories,
- tags,
- nutrition,
- preparation instructions,
- preparation time.

---

## Planned Meal

A recipe assigned to a particular day.

Example:

```text
Monday

Breakfast
Protein Pancakes

Lunch
Chicken Curry

Dinner
Tomato Soup
```

Planning does not mean the meal was actually eaten.

---

## Meal History

A historical record of meals that were actually eaten.

Example:

```text
Chicken Curry

Last eaten:
8 days ago

Previously eaten:
12 September
28 August
17 August
4 August
```

This historical information becomes one of the application's most valuable features.

---

# 5. Core User Journey

The normal flow of the application should be:

```text
Create Ingredients
        ↓
Create Recipes
        ↓
Choose Meals
        ↓
Put Meals on Calendar
        ↓
Eat Meal
        ↓
Mark as Eaten
        ↓
Build Meal History
        ↓
Use History to Choose Future Meals
```

The system becomes more useful over time as more history is collected.

---

# 6. Ingredient Management

Users can maintain a shared database of ingredients.

Each ingredient may contain:

- name,
- calories per 100 g,
- protein,
- carbohydrates,
- fat,
- category,
- tags,
- vegetarian information,
- vegan information.

Example:

```text
Chicken Breast

165 kcal / 100g
Protein: 31g
Carbs: 0g
Fat: 3.6g

Category:
Meat

Tags:
high-protein
lean
chicken
```

Once created, the ingredient can be reused in many recipes.

---

# 7. Recipe Management

Users can create and maintain recipes.

Example:

```text
Chicken Curry

Ingredients:

500g chicken breast
300g rice
200g coconut milk
150g onion
10g olive oil
```

The system calculates nutritional information based on the ingredients.

Recipes can also contain:

- instructions,
- cooking time,
- preparation time,
- meal types,
- tags,
- categories,
- final cooked weight.

---

# 8. Recipe Classification

Recipes can belong to multiple meal types.

Example:

```text
Protein Pancakes

Breakfast
Snack
Dessert
```

Recipes may also have categories such as:

```text
Pasta
Soup
Rice Dish
Salad
Pizza
Curry
Dessert
```

And flexible tags such as:

```text
quick
high-protein
cheap
meal-prep
weekend
comfort-food
spicy
```

This makes the recipe collection searchable and useful rather than becoming a large unorganized list.

---

# 9. Recipe Search

Users should be able to search recipes based on practical criteria.

Examples:

```text
Show dinners
```

or:

```text
Show meals below 160 kcal / 100g
```

or:

```text
Dinner
High protein
Quick
Not eaten recently
```

Filtering may include:

- calories,
- macros,
- meal type,
- category,
- tags,
- vegetarian status,
- ingredients,
- preparation time,
- last eaten date.

---

# 10. Meal Calendar

The calendar represents both:

```text
What are we planning to eat?
```

and:

```text
What did we actually eat?
```

Example:

```text
MONDAY

Breakfast
Protein Pancakes
✓ Eaten

Lunch
Chicken Curry
✓ Eaten

Dinner
Tomato Soup
Planned
```

Calendar entries may have states such as:

```text
PLANNED
EATEN
SKIPPED
```

This distinction is important because plans often change.

---

# 11. Meal History

Whenever a meal is marked as eaten, it becomes part of the household's eating history.

This allows the system to calculate:

- last eaten date,
- number of times eaten,
- times eaten this month,
- days since last eaten.

Example:

```text
Lasagne

Last eaten:
23 days ago

Times eaten:
8

Times eaten during last 30 days:
1
```

---

# 12. Meal Rotation

One of the main practical features is helping users avoid repeatedly eating the same meals.

The application should allow users to find:

```text
Meals not eaten for 14+ days
```

or:

```text
Dinner recipes sorted by longest time since last eaten
```

Example:

```text
Thai Curry
Last eaten 31 days ago

Lasagne
Last eaten 26 days ago

Salmon with Potatoes
Last eaten 18 days ago
```

This becomes a simple meal recommendation mechanism without requiring AI.

---

# 13. "What Should We Eat?"

A major future business feature is helping users decide what to cook.

Initially this can simply be a specialized recipe search.

Example:

```text
We need dinner.

Requirements:

< 700 kcal / portion
high protein
< 30 min preparation
not eaten during last 10 days
```

The application returns matching recipes.

Later recommendations may also consider:

- meals already eaten recently,
- ingredients available at home,
- household preferences,
- nutrition,
- planned meals,
- ingredient expiration dates.

---

# 14. Recipe History

Recipes may change over time.

For example:

```text
Chicken Pasta v1

200g mozzarella
30g olive oil
```

later becomes:

```text
Chicken Pasta v2

120g mozzarella
10g olive oil
```

The current recipe may change, but historical meal entries should still represent what was eaten at the time.

Therefore recipe history/versioning should eventually be supported.

---

# 15. Weekly Meal Planning

Users should be able to plan an entire week.

Example:

```text
MONDAY
Chicken Curry

TUESDAY
Pasta Carbonara

WEDNESDAY
Salmon + Potatoes

THURSDAY
Tomato Soup

FRIDAY
Homemade Pizza
```

Planning should be lightweight.

Moving or replacing meals should be easy.

---

# 16. Future Shopping List

Once recipes exist and meals are planned, the application can generate shopping lists automatically.

Example meal plan:

```text
Monday
Chicken Curry

Tuesday
Chicken Pasta

Wednesday
Salmon + Potatoes
```

The application can aggregate ingredients:

```text
Shopping List

Chicken breast
1.2 kg

Rice
300 g

Pasta
400 g

Salmon
500 g

Potatoes
1 kg

Tomato passata
500 g
```

This is not necessary for the first version but is a natural evolution of the product.

---

# 17. Future Pantry Integration

A later version may know what ingredients are already available at home.

Example:

```text
Available:

Rice
Potatoes
Eggs
Chicken breast
Onion
```

The system could then answer:

```text
You can make:

Chicken Rice
Chicken Potato Bowl
Fried Rice
```

or:

```text
Chicken Curry

Missing:
Coconut milk
```

Pantry management is not required initially.

---

# 18. Dashboard

The application's home page should answer useful household questions immediately.

Example:

```text
TODAY
────────────────────

Breakfast
Protein Pancakes
✓ Eaten

Lunch
Chicken Curry
✓ Eaten

Dinner
Tomato Soup
Planned


RECENTLY EATEN
────────────────────

Chicken Pasta
2 days ago

Pizza
5 days ago

Salmon
8 days ago


NOT EATEN RECENTLY
────────────────────

Lasagne
28 days

Thai Curry
25 days

Burrito
19 days
```

The dashboard should focus on decisions rather than statistics.

---

# 19. Core Business Capabilities

The application's initial capabilities can be summarized as:

```text
INGREDIENTS

Create
Edit
Categorize
Tag
Store nutrition


RECIPES

Create
Edit
Add ingredients
Calculate nutrition
Categorize
Tag
Search
Filter


CALENDAR

Plan meals
Move meals
Mark eaten
Mark skipped


HISTORY

Track meals
Last eaten
Frequency
Meal rotation


DECISION SUPPORT

Find recipes
Find meals not eaten recently
Help choose next meal
```

---

# 20. MVP Scope

The first usable version should contain:

## Ingredients

- create ingredient,
- edit ingredient,
- nutritional information,
- categories,
- tags.

## Recipes

- create recipe,
- edit recipe,
- reusable ingredients,
- ingredient quantities,
- automatic nutrition calculation,
- meal types,
- categories,
- tags,
- search,
- filtering.

## Calendar

- week view,
- month view,
- add recipe to day,
- select meal type,
- planned status,
- eaten status,
- skipped status.

## Meal History

- last eaten date,
- number of times eaten,
- recipe eating history,
- sorting by least recently eaten.

---

# 21. What the MVP Is NOT

The MVP should not become:

```text
MyFitnessPal clone
```

It is not primarily about strict calorie logging.

It should also not become:

```text
warehouse management software
```

Pantry management can come later.

And it should not become:

```text
AI meal planner
```

before enough useful household data exists.

The initial value comes from organizing real recipes and building useful meal history.

---

# 22. Product Evolution

The intended progression is:

```text
V1

Ingredients
Recipes
Calendar
Meal History


        ↓


V2

Better meal planning
Meal rotation
"What should we eat?"
Recipe history


        ↓


V3

Shopping lists
Pantry
Ingredient availability
Expiration dates


        ↓


V4

Smart recommendations
Nutrition-aware suggestions
Automatic weekly planning
```

---

# 23. Main Product Principle

The application should reduce this conversation:

```text
"What should we eat today?"

"I don't know."

"What did we eat recently?"

"I don't remember."

"What recipes do we have?"

"Somewhere in Notes..."
```

to:

```text
Open app

↓

Dinner
Not eaten in 14 days
< 30 minutes

↓

Chicken Curry
Lasagne
Salmon Bowl

↓

Add to Tuesday
```

That is the core business value of the product.

---

# 24. One-Sentence Product Definition

> **A shared household recipe and meal-planning diary that remembers what you eat and helps you decide what to eat next.**

---

# 25. Short Product Pitch

Meal Diary is a private household application for couples and families who want to organize their recipes and meal planning in one place.

Users maintain a reusable ingredient and recipe database, plan meals using a shared calendar, and mark meals that were actually eaten.

Over time, the application builds a household meal history that makes it possible to see what was eaten recently, identify meals that have not been prepared for a long time, filter recipes based on nutrition or preferences, and make future meal decisions more easily.

Later, the same information can be used to automatically create shopping lists, track available ingredients, and provide intelligent meal recommendations.