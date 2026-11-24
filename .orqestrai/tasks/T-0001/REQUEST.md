---
id: T-0001
created: 2026-11-23T14:31:29.086Z
---

Update DESIGN.md and the discovery to the decisions of 2026-11-23: units are g/kg, ml/l and pieces (reverse Q060 and C004; mass, volume and pieces of one ingredient do not combine with each other); no NATS, events or outbox until a feature needs them, so services only ask each other over gRPC and never copy data (planning asks recipes for matching recipes and sorts them by its own history); keep five separate apps, each with its own image.
