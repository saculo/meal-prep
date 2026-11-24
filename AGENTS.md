# meal-prep

Meal planning for a two-person household, as services in one repository. Read `DESIGN.md`
first, then the discovery in `all-event-storms/` (start with `problem-brief.md`).

- Each service is a top-level folder (`recipes/`, `planning/`, `shopping/`, `gateway/`,
  `web/`) with its own `AGENTS.md`; its `## Checks` are what orq's gate runs for it.
- A changed discovery rule gets a new corpus entry; the old entry keeps its text and gets
  "Superseded by Qxxx (date)", a contradiction it settled is reopened, and the problem
  brief's evidence index lists every entry still current.
- No service exists yet.
