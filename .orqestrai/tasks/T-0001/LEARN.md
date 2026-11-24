---
id: T-0001
status: applied
---

# Learn: gRPC-only decision recorded, discovery superseding rule

## What happened
Documents-only chore: DESIGN.md and all-event-storms/ brought in line with the 2026-11-23
answers (units g/kg, ml/l, pieces; optional portions; gRPC only, no events or copies; five
apps). Three parallel tasks split by file, with Q065–Q069 fixed in the spec up front. Verify
passed in round 1 with no findings; two notes below threshold (evidence index left out
current entries Q055–Q058, Q062–Q064; a Q045 citation dropped). No gate configured, V1 manual.

## Proposals
| # | Kind | Change | Evidence in this task | Applied in |
|---|------|--------|-----------------------|------------|
| L1 | adr | Record "services ask over gRPC and never copy; no events until a feature needs them", with the reversed NATS/copy design and its costs (recipes must serve filtering, runtime dependency). DESIGN.md states the rule but not why or what it costs. | REQUEST; BRIEF Q5, Decisions; SPEC AC4, Q10; RESEARCH "DESIGN.md has gRPC requests plus NATS events" | docs/adr/0001-services-ask-over-grpc-and-never-copy.md |
| L2 | agent-rules | How to change a discovery rule: new corpus entry, old entry keeps text with "Superseded by Qxxx (date)", settled contradiction reopened, evidence index lists all current entries. | BRIEF Q2; SPEC Design; VERIFY round 1 hypothesis 2 and "For the PR body" (evidence index gap) | AGENTS.md |

## Considered, not proposed
- L3 (glossary "App" vs "service", AGENTS.md wording): dropped by the orchestrator. It made gateway "not a service" while DESIGN.md, this task's deliverable, lists gateway in its Services table; it would add a contradiction instead of settling one.
- Testing / verify skill: no code and no app; Q11 chose no checks. Nothing to record yet.
- Out-of-scope: the user rejected nothing; the dropped NATS design is covered by L1.
- Gate "not configured": expected for a docs-only repo with no app folders yet.
- DESIGN.md's services table still has a `gateway` row ("No domain"); left as is (task
  deliverable), a later task can rename the table to apps if it matters.
