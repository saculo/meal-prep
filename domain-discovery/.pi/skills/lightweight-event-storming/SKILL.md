---
name: lightweight-event-storming
description: A lightweight Event Storming facilitation method for discovery. Extracts events, commands, actors, policies, data, hotspots, external systems and work seeds from real user conversation.
---

# Lightweight Event Storming

## Purpose

Use when the domain/process is unclear and the user has tacit knowledge. This is a facilitation/scribing skill, not a fake-stakeholder simulator.

## Core concepts

- **Event** — something meaningful that happened, past tense: `Task was approved`.
- **Command/action** — intent that tries to cause an event: `Approve task`.
- **Actor** — person/system initiating a command.
- **Policy** — rule connecting events/conditions to commands.
- **Read model/data** — information needed to decide or act.
- **External system** — outside dependency.
- **Hotspot** — pain, risk, confusion, conflict, manual workaround, compliance concern.

## Interview flow

1. Ask for the process boundaries:
   - Where does this start?
   - Where does it end?
   - What outcome matters?
2. Ask for the happy-path events in order.
3. For each event, ask:
   - What caused it?
   - Who/what did it?
   - What rule allowed/blocked it?
   - What data was needed?
   - What could go wrong?
4. Ask for unhappy paths and exceptions.
5. Ask which hotspots are most expensive or frequent.
6. Convert strong hotspots into work seeds.

## Output table

```md
| # | Event | Command/action | Actor | Policy/rule | Data | Hotspot? | Source |
|---|---|---|---|---|---|---|---|
```

## Facilitation prompts

- "What happened next?"
- "What made that possible?"
- "Who decides that?"
- "What do people do manually today?"
- "Where do you wait, retry or escalate?"
- "What would be bad if we automated this wrong?"
- "Which term here is overloaded or ambiguous?"

## Work seed rule

A hotspot becomes a work seed only when it can be phrased as:

```text
Given <current/context>
When <trigger>
Then <observable better outcome>
```

Otherwise keep it as an open question or domain note.
