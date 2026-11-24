---
id: T-0001
---

# Research: units, how services talk, deployment (DESIGN.md and all-event-storms/)

Questions were asked without the request, so the findings describe the documents as they are.
The repo has no code and no tests; there is no GLOSSARY.md and no docs/adr/.

## The documents disagree on piece amounts

**Question:** What do the documents say about amount units, and how amounts combine?

Q043 has piece amounts ("200g tomato + 1 tomato") listed side by side and ml/l converted.
Q060 later required every amount to be mass or volume; C004 resolved the clash for Q060
("no piece amounts"). DESIGN.md still lists `pcs`, and R010's scenario still keeps "grams plus
pieces" visible together. Q062: mass and volume are never converted, for shopping or nutrition.
Q063: line calories = line amount × calories per 100 g/ml. Q064: one ingredient may be used by
mass in one recipe and by volume in another; DESIGN.md gives the ingredient a single "unit".

**Evidence:**
- `all-event-storms/corpus.md:349-355`: Q043.
- `all-event-storms/corpus.md:487-493`: Q060, notes "Piece counts are not allowed".
- `all-event-storms/contradictions.md:85-103`: C004, "Resolved in favour of Q060".
- `all-event-storms/corpus.md:503-525`: Q062, Q063, Q064.
- `DESIGN.md:36`: ingredient "(name, unit, optional calories)"; `DESIGN.md:48`: "g/kg, ml/l, pcs".
- `all-event-storms/work-candidates.md:459-463`: R010 "grams plus pieces remain visible together".
- `all-event-storms/problem-brief.md:23,42,56,60`: compatible units; evidence index Q001–Q050 only.

**Tests:** none found.

## Q060 and C004 are cited only in corpus.md and contradictions.md

**Question:** Where is Q060 recorded and who cites Q060 or C004?

**Evidence:**
- `all-event-storms/corpus.md:487`: the only record of Q060 (under user-supplied answers).
- `all-event-storms/contradictions.md:85,95`: the only citations. DESIGN.md, work-candidates.md,
  problem-brief.md cite neither. questions.md has only Q001–Q003 (= corpus Q051–Q053).

**Tests:** none found.

## DESIGN.md has gRPC requests plus NATS events with a copy of recipe data in planning

**Question:** How do the documents describe services talking to each other?

Only DESIGN.md covers it. gRPC for requests, "only IDs cross service lines". NATS JetStream
events, durable consumers, outbox per schema, a broker-neutral interface (swappable for Kafka or
RabbitMQ). Planning listens to `recipes.*` and keeps its own copy of recipe names and
ingredients for suggestions. This copy conflicts with "none reads another's data" and "only IDs
cross service lines". NATS also appears in Data (runs in the cluster), Testing (event tests) and
First steps (step 1).

**Evidence:**
- `DESIGN.md:8-9`, `DESIGN.md:37-39`, `DESIGN.md:52-64`, `DESIGN.md:71-72`, `DESIGN.md:89-90`, `DESIGN.md:101`.

**Tests:** none found.

## Suggestions: ordering uses planning's own history; recipe data came from the copy

**Question:** How are planning's suggestions described and what data do they need from recipes?

Never eaten first, then oldest last eaten (R002, Q017, Q018); ingredient match is "any selected
ingredient" (R004, Q016); result cards show name, last eaten date, matching ingredients (Q019);
random fallback from all recipes (R009, Q053). History is planning's own `eaten` entries.
Recipe names/ingredients came from the event-fed copy; which service does the matching is not
stated. Recipe deletion is not in the discovery at all (only as an example event subject).

**Evidence:**
- `DESIGN.md:37,43,47,52-57`; `all-event-storms/work-candidates.md:56-83,153-171,398-420`;
  `all-event-storms/corpus.md:127-157`.

**Tests:** none found.

## Five folders, each its own deployment and image; web has no Dockerfile described

**Question:** How are deployment and packaging described?

recipes, planning, shopping, gateway, web: each its own deployment and image; previews run all
five; CI builds only changed folders. The image rule says "service" and the layout shows a
Dockerfile only for recipes.

**Evidence:**
- `DESIGN.md:7-8,13-25,82-85`.

**Tests:** none found.

## Open

- DESIGN.md's recipe "portions" disagrees with Q054/Q059/Q061 (the recipe declares a basis: per portion or per batch).
