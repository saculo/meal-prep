---
name: guided-conversation
description: Ask small batches of evidence-anchored clarification questions, record answers, and stop when enough information exists to produce useful craftpath-ready candidates.
---

# Guided Conversation

## Principle

Ask only questions that change the output. Do not interrogate the user for completeness when the current evidence is enough for useful small work candidates.

## Question rules

Ask when:

- a contradiction blocks a candidate
- an actor/user is unknown
- the observable outcome is unknown
- acceptance cannot be written as Given/When/Then
- scope boundaries are unclear
- a safety/compliance constraint may exist

Do not ask when:

- the answer would only affect implementation details
- a reasonable assumption can be recorded safely
- the candidate can be marked `needs-questions`

## Batch format

```md
I need <N> answers to make the candidates executable:

1. <question>
   Why: <what this decides>
   Evidence: <quote/source>

2. ...
```

## Recording answers

Write to `discovery/<slug>/questions.md`:

```md
# Questions

## Q001 — <question>

- status: answered|open|deferred
- why: <what this decides>
- answer: <user answer or blank>
- affects: [R001, R003]
```

Also add answered content to the evidence corpus under `user-supplied`.

## Termination

Stop asking when one of these is true:

- at least one useful candidate is `ready`
- all remaining candidates are clearly `needs-questions`
- the user says to stop
