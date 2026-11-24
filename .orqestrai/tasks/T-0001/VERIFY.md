---
id: T-0001
status: passed
round: 1
---

# Verify: DESIGN.md and the discovery follow the 2026-11-23 decisions

## Round 1

### Quality gate
| Command | Result | Evidence |
|---------|--------|----------|
| (all) | not configured | No `AGENTS.md` with a `## Checks` section above any changed file (root `AGENTS.md` only mentions `## Checks` in prose). Not configured: `.`, `all-event-storms/`, `.orqestrai/tasks/T-0001/`. |

### Checks (verifier)
| Check | Claim | Verdict | Evidence |
|-------|-------|---------|----------|
| V1 | AC1–AC6 | manual: waits for the user at the PR | Kind `manual` (Q11); not sent to the verifier. |

### Review (reviewer)
**Hypotheses:**
1. An old rule survives ("no pieces", NATS/events/outbox/subjects/listeners, per-batch basis): not confirmed. The only event mention is the allowed sentence (DESIGN.md:73).
2. Superseded entries deleted or rewritten instead of annotated: not confirmed. Q004 (portions part), Q054, Q059, Q060, Q061 keep their text with a "Superseded by Q06x (2026-11-23)" note; C004 keeps its old resolution marked void.
3. Q numbers or examples drift between the three parallel tasks: not confirmed. Q065–Q069 and the three Chili/Omelette examples agree in all files.
4. GLOSSARY.md missing Unit or Portions, or the image name leaving out an app: not confirmed.

**Prior findings:** none (round 1).

| # | Axis | Finding | Confidence | Blocking | Where |
|---|------|---------|------------|----------|-------|
| — | — | No findings at confidence 75 or above | — | — | — |

Own check: DESIGN.md read in full after W1.3; grep for nats/outbox/listen/jetstream/swappable/pcs/"no piece"/batch over DESIGN.md, problem-brief.md, work-candidates.md finds only DESIGN.md:56 ("no per-portion or per-batch basis"), a negation.

### Outcome
passed

Root causes: none

## For the PR body
- Below the reviewer's threshold: the problem brief's evidence index reads "Q001–Q053 and Q065–Q069" and leaves out entries still current (Q055–Q058, Q062–Q064).
- Below the reviewer's threshold: the brief's shopping-list line dropped its Q045 citation, though Q045 was only clarified, not superseded.
- No automated checks run for this task (Q11); V1 is the user's read of the diff against BRIEF.md.
