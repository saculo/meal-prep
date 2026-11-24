# meal-prep — technical design

From the discovery in `all-event-storms/` (problem brief, R001–R011; corpus entries Q043–Q069, the units and portions rules from 2026-11-23). No code yet. Decisions
marked **open** are yours.

## Shape: services in one repository

Each part of the domain is a **service**: its own program, image, deployment and schema in
Postgres. Services ask each other over **gRPC** only; none reads another's data and none
copies it. The UI talks to one **gateway**, which calls the services. Go first; a service
can be rewritten in another language (Gleam) behind the same contract.

There are five apps: `recipes`, `planning`, `shopping`, `gateway` and `web`. Each is its own
deployment and its own image.

```
meal-prep/
  recipes/      Go: go.mod, proto/, cmd/, internal/, Dockerfile, AGENTS.md
  planning/     Go, same layout
  shopping/     Go, same layout
  gateway/      Go: HTTP + JSON for the UI (OpenAPI), calls the services over gRPC; Dockerfile
  web/          Astro: the UI; Dockerfile
  go.work       the Go services as one workspace, for the editor and local builds
```

Each app folder is complete on its own: its contract (`proto/`, for the services), code,
tests, Dockerfile, and an `AGENTS.md` whose `## Checks` lists the commands orq's gate runs
for it. Every app's image is `ghcr.io/saculo/meal-prep-<app>:<commit>`, with `<app>` one of
`recipes`, `planning`, `shopping`, `gateway`, `web`.

**Contracts live with the service that serves them** (`recipes/proto/`). The service
generates its Go server and client into its own module (`recipes/api`); a caller imports
that package (`go.work` locally, a version tag otherwise). A contract change is then one
pull request in one folder, and `buf breaking` in that service's check guards its callers.

## Services

| Service | Owns | Work candidates | Calls |
|---|---|---|---|
| **recipes** | Recipes (name, tags, optional portions, ingredient lines), capture of a recently cooked recipe; reusable ingredients (name, one or more units, optional calories per 100 g, per 100 ml, per piece), similar-name warnings; finds recipes by filters and selected ingredients | R005, R006, R007 | — |
| **planning** | Meal entries (date, meal type, optional portions, recipe ID, `planned`/`eaten`), moving, history; suggestions: asks recipes for the matching recipes (filters, selected ingredients), then sorts them by its own eaten history: never eaten first, then oldest last eaten, random fallback | R001–R004, R008, R009 | recipes |
| **shopping** | List for a date range: asks planning for the planned meals and recipes for their lines, scales by portions, combines compatible units, flags incomplete items, check-off, text to share. Stores its own generated list: items with name, amount and unit at generation. That is shopping's own record, not a copy kept in sync | R010, R011 | planning, recipes |
| **gateway** | No domain: the UI's API, auth later | — | all |

Rules from the discovery that shape the contracts:
- A meal entry is `planned` or `eaten` (Q051). History = eaten entries; none = never eaten.
- Calories are optional everywhere; missing nutrition is shown as incomplete, never blocks (Q052).
- **Unit** (Q065): an amount is in g or kg (mass), ml or l (volume), or pieces. g and kg
  convert, ml and l convert; mass, volume and pieces of one ingredient never combine.
  Pieces never convert to g, even for an ingredient that has both.
- A reusable ingredient has one or more units (an egg in pieces or in g); a recipe line picks
  one (Q066). Calories are optional per 100 g, per 100 ml and per piece (Q067). A line whose
  unit has no calorie value shows nutrition as incomplete.
- **Portions** (Q068, Q069): optional on the recipe and on the planned meal; missing = 1. The
  shopping list scales a line by meal portions ÷ recipe portions. There is no
  per-portion or per-batch basis and no whole batches. Examples:
  - Chili, 400 g beans, 4 portions, planned for 2: 200 g.
  - Omelette, 3 eggs, no portions, planned for 2: 6 eggs.
  - Chili planned without portions: 100 g.
- An ingredient line is a reusable ingredient (ID) or temporary free text (R006); shopping
  scales and combines only the reusable ones, free text is listed as written.
- Repeat and calorie warnings never block (Q012, Q013).
- Random fallback picks from all recipes (Q053).
- Combining: same ingredient, compatible units only; no fuzzy matching (Q043).

## How services talk

- **gRPC**, generated with buf. A service asks another for what it needs and never copies
  the answer into its own schema. `planning` keeps a recipe ID and asks `recipes` for the
  recipes that match the filters and selected ingredients, then sorts them by its own
  eaten history. `shopping` asks `planning` for the planned meals and `recipes` for their
  lines.
- Events come when a feature needs them; none today.
- Each service answers `/version` (the commit) and a health check.

## Data

- **One Postgres, one schema and one role per service** (`recipes`, `planning`,
  `shopping`); a role can use only its own schema. Migrations per service, in its folder.
- A preview gets its own database (the same schemas, empty), dropped with the preview.
- Postgres runs in the cluster (orqestrai-infra), not in this repository.

## UI

- **Astro** (`web/`), the only way the app is used; no CLI.
- **Gateway → UI: HTTP + JSON**, described by OpenAPI; the UI's client is generated from it.
- **open**: GraphQL in the gateway instead, if the screens need many combined views.

## Deployment

- Each of the five apps (recipes, planning, shopping, gateway, web) is its own deployment, with its own image.
- A preview of a pull request runs **all** of them (the changed ones from the PR's commit,
  the others from `main`), so the UI can be tested end to end.
- CI (set up separately) checks and builds the image only for the folders a pull request changed.

## Testing

- Per service: unit tests of the rules (each R00x scenario), store tests against Postgres
  (testcontainers), gRPC tests through the generated client.
- Contracts: `buf lint`, `buf breaking` against `main`, in the service's check.
- End to end: e2e against the preview, through the UI.

## orqestrai in a repo like this

Each service's `AGENTS.md` has a `## Checks` section with its commands; orq's gate runs the
checks of each service the branch changed, in that service's folder.

## First steps

1. Postgres in the cluster; `recipes` with its contract (R005–R007).
2. `gateway` and `web` with the first screen (capture a recipe).
3. `planning` (R001–R004, R008, R009), then `shopping` (R010, R011).
