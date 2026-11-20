---
name: evidence-corpus
description: Build a quote-backed evidence corpus from discussions, notes, transcripts, issues, docs or Event Storming sessions. Every candidate requirement must cite evidence. Preserve quotes verbatim and source each claim.
---

# Evidence Corpus

## Purpose

Prevent requirements from being invented from vibes. Every important claim should be traceable to something someone said or wrote.

## Output file

`discovery/<slug>/corpus.md`

## Format

```md
# Evidence Corpus

## <source path or user-supplied>

### Q001
> <verbatim quote>

- source: <file/path, message, interview answer, board image, etc.>
- loc: <line number, timestamp, section, or n/a>
- tags: [pain|goal|rule|constraint|metric|actor|event|risk|integration|term]
- notes: <short interpretation outside the quote>
```

## Rules

- Preserve source language and exact wording inside quotes.
- Keep interpretation outside quotes.
- Prefer many small quotes over one huge quote.
- A candidate requirement should cite at least two quotes when possible.
- A candidate with no evidence is an assumption, not a requirement.
- User answers gathered during the session become `user-supplied` evidence.

## Useful tags

- `pain` — something hurts today
- `goal` — desired outcome
- `rule` — business rule
- `constraint` — limitation, compliance, compatibility
- `metric` — measurable target
- `actor` — user/stakeholder/system
- `event` — something that happened in the domain
- `risk` — failure, cost, ambiguity
- `integration` — external dependency
- `term` — domain vocabulary
