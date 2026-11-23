# meal-prep — technical design

From the discovery in `all-event-storms/` (problem brief, R001–R011). No code yet. Decisions
marked **open** are yours.

## Shape: services in one repository

Each part of the domain is a **service**: its own program, image, deployment and schema in
Postgres. Services talk over **gRPC** (requests) and **NATS JetStream** (events); none reads
another's data. The UI talks to one **gateway**, which calls the services. Go first; a
service can be rewritten in another language (Gleam) behind the same contract.

```
meal-prep/
  recipes/      Go: go.mod, proto/, cmd/, internal/, Dockerfile, AGENTS.md
  planning/
  shopping/
  gateway/      Go: HTTP + JSON for the UI (OpenAPI), calls the services over gRPC
  web/          Astro: the UI
  go.work       the Go services as one workspace, for the editor and local builds
```

Each service folder is complete on its own: its contract (`proto/`), code, tests,
Dockerfile, and an `AGENTS.md` whose `## Checks` lists the commands orq's gate runs for it.
Its image is `ghcr.io/saculo/meal-prep-<service>:<commit>`.

**Contracts live with the service that serves them** (`recipes/proto/`). The service
generates its Go server and client into its own module (`recipes/api`); a caller imports
that package (`go.work` locally, a version tag otherwise). A contract change is then one
pull request in one folder, and `buf breaking` in that service's check guards its callers.

## Services

| Service | Owns | Work candidates | Calls / listens |
|---|---|---|---|
| **recipes** | Recipes (name, tags, portions, ingredient lines), capture of a recently cooked recipe; reusable ingredients (name, unit, optional calories), similar-name warnings | R005, R006, R007 | — |
| **planning** | Meal entries (date, meal type, portions, recipe ID, `planned`/`eaten`), moving, history; suggestions: never eaten first, then oldest last eaten, filters, ingredient match, random fallback | R001–R004, R008, R009 | recipes; listens `recipes.*` |
| **shopping** | List for a date range from planned meals: scaled by portions, compatible units combined, incomplete items flagged, check-off, text to share | R010, R011 | planning, recipes |
| **gateway** | No domain: the UI's API, auth later | — | all |

Rules from the discovery that shape the contracts:
- A meal entry is `planned` or `eaten` (Q051). History = eaten entries; none = never eaten.
- Calories are optional everywhere; missing nutrition is shown as incomplete, never blocks (Q052).
- An ingredient line is a reusable ingredient (ID) or temporary free text (R006); shopping
  scales and combines only the reusable ones, free text is listed as is.
- Repeat and calorie warnings never block (Q012, Q013).
- Random fallback picks from all recipes (Q053).
- Unit combining: same ingredient, compatible units only (g/kg, ml/l, pcs); no fuzzy matching (Q043).

## How services talk

- **Requests: gRPC**, generated with buf. Only IDs cross service lines (`planning` keeps a
  recipe ID, asks `recipes` for the rest).
- **Events: NATS JetStream.** A service publishes what happened in its own subjects
  (`recipes.recipe.deleted`, `planning.meal.eaten`); a listener keeps what it needs in its
  own schema (planning's copy of recipe names and ingredients for suggestions). Durable
  consumers, so a restarted service catches up. Trying it from the start; dropped if it
  does not pay off.
- **Kept swappable for Kafka or RabbitMQ**: every service publishes and listens only through
  a small interface of its own (`Publish(event)`, `Listen(name, handler)`), events are plain
  protobuf messages in an envelope (id, type, time, data) that names no broker, and a
  handler can take the same event twice (it checks the id). An outbox table in each
  schema publishes after the commit. No NATS-only features (request-reply, KV, object
  store). Swapping the broker is then one adapter and the infra, not the services' code.
- Each service answers `/version` (the commit) and a health check.

## Data

- **One Postgres, one schema and one role per service** (`recipes`, `planning`,
  `shopping`); a role can use only its own schema. Migrations per service, in its folder.
- A preview gets its own database (the same schemas, empty), dropped with the preview.
- Postgres and NATS run in the cluster (orqestrai-infra), not in this repository.

## UI

- **Astro** (`web/`), the only way the app is used; no CLI.
- **Gateway → UI: HTTP + JSON**, described by OpenAPI; the UI's client is generated from it.
- **open**: GraphQL in the gateway instead, if the screens need many combined views.

## Deployment

- Every service, the gateway and the web app is its own deployment, with its own image.
- A preview of a pull request runs **all** of them (the changed ones from the PR's commit,
  the others from `main`), so the UI can be tested end to end.
- CI (set up separately) checks and builds the image only for the folders a pull request changed.

## Testing

- Per service: unit tests of the rules (each R00x scenario), store tests against Postgres
  (testcontainers), gRPC tests through the generated client, event tests against NATS.
- Contracts: `buf lint`, `buf breaking` against `main`, in the service's check.
- End to end: e2e against the preview, through the UI.

## orqestrai in a repo like this

Each service's `AGENTS.md` has a `## Checks` section with its commands; orq's gate runs the
checks of each service the branch changed, in that service's folder.

## First steps

1. Postgres and NATS in the cluster; `recipes` with its contract (R005–R007).
2. `gateway` and `web` with the first screen (capture a recipe).
3. `planning` (R001–R004, R008, R009), then `shopping` (R010, R011).
