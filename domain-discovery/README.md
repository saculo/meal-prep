# domain-discovery

A small discovery toolkit for turning messy product/domain input into **craftpath-ready work items**.

It is inspired by DomainDrivers `dd-ai` episodes:

- `002` — evidence corpus, contradiction finding, guided clarification
- `003` — lightweight Event Storming facilitation
- `006` — solution-free problem brief, context maps/modules and blind requirement generation
- `007` — problem-class classification
- `004/005` — optional implementation-time scouting for algorithms and coupling

The toolkit deliberately stops before implementation. Its output is meant to be copied into a project and then executed with craftpath, e.g.

```bash
/craftpath-work "Add approval audit trail"
```

## Installed shapes

Both harness layouts are present:

- Claude Code: `.claude/commands/` and `.claude/skills/`
- pi: `.pi/prompts/` and `.pi/skills/`

## Main commands

### `/domain-discover`

Use for notes, transcripts, issue threads, rough ideas, or discussion summaries.

Produces:

```text
discovery/<slug>/
  corpus.md
  contradictions.md
  questions.md
  problem-brief.md
  work-candidates.md
```

### `/event-storm`

Lightweight Event Storming interview mode. Use it to extract domain events, commands, actors, policies, hotspots and open questions from a conversation.

### `/context-map`

Turns discovery artifacts into evidence-backed bounded context candidates and a context map. Separates business areas from bounded contexts and records weak boundaries as hypotheses.

Produces:

```text
discovery/<slug>/
  context-candidates.md
  context-map.md
```

### `/module-design`

Turns a context map into module blueprints and craftpath-ready module work candidates. Does not implement code.

Produces:

```text
discovery/<slug>/
  modules.md
  module-work.md
```

## Recommended flow

```text
raw discussion / notes
  -> /event-storm            # optional, when the domain/process is unclear
  -> /domain-discover        # builds evidence + candidate work items
  -> /context-map            # bounded context candidates + relationships
  -> /module-design          # module blueprints + module work candidates
  -> choose one module work candidate
  -> /craftpath-work "..."   # in the target implementation repo
```

## Output contract for craftpath

Each candidate should contain:

- title suitable for `craftpath work new`
- problem statement
- Given/When/Then scenarios
- out of scope
- assumptions
- unresolved questions
- evidence quotes
- suggested craftpath skills, e.g. `backend`, `frontend`, `testing`, `architecture`
- implementation notes only as hypotheses, never as the requirement itself
