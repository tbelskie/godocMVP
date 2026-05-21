# ADR 0001: Overall Architecture & Tech Stack

**Date:** 2025-05-21  
**Status:** Accepted  
**Deciders:** Tom (Founder) + Grok (Technical Co-Founder)

## Context

We are building a CLI tool that delivers the "magic moment" of instant beautiful Hugo docs sites. Goal: $140k ARR lifestyle business.

## Decision

We chose:

- **Language**: Go 1.23+
- **CLI Framework**: Cobra + Viper
- **Structure**: Standard `cmd/` + `internal/` layout
- **Embedding**: `go:embed` for the Hugo skeleton and theme
- **SDLC**: GitHub Issues + Projects + ADR + Spec templates
- **Dogfooding**: godocMVP documents itself from day 1
- **Licensing**: MIT for core, closed for Pro features

## Rationale

- Go + Cobra = single static binary, battle-tested in Hugo/gh CLI
- `internal/` keeps code clean and private
- Embedding = zero user friction for `init`
- Structured process builds Tom's founder skills across all areas

## Alternatives Considered

- Pure shell script / Python: Too slow, not professional
- Docusaurus base: Heavier, not our Hugo wedge

## Consequences

- Positive: Fast, professional, great portfolio piece
- Positive: High involvement learning path for Tom
- Risk: Go learning curve (mitigated by step-by-step)

