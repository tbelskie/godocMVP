# Agents Guide

This file orients new AI agents (and human teammates) to the Pendragon project (repo: godocMVP).
Read these in order at the start of every session.

## 1. Principles (auto-loaded)

`.cursor/rules/godoc.md` — the four absolute rules: Security First, Simple
& Elegant, Process (focused PRs linked to issues), and Pendragon Values.

## 2. Recent context

`docs/AGENTS_JOURNAL.md` — append-only session log, newest entries on top.
Read the last 2–3 entries to understand what shipped, what's in flight,
and what the next session should pick up.

## 3. Where we're going

- `ROADMAP.md` — overall vision, current priorities, slice plans
- GitHub Issues — the spec system (each significant change has one)
- Open PRs (`gh pr list`) — current branches and review state

## 4. How we work

- `PROCESS.md` — SDLC: idea → issue → spec → branch → PR → review → merge
- `docs/decisions/` — Architecture Decision Records (the "why" behind the "what")

## 5. Working agreements

- Branch off `main` for every focused PR. Never edit `main` directly.
- Each PR addresses one slice of one issue. Keep them small and shippable.
- Reference the GitHub Issue in the commit body (`Refs #N` or `Closes #N`).
- Run `go vet ./... && go build ./... && go test ./...` before pushing.
- Add new ADRs in `docs/decisions/` when making architectural calls.

## 6. End-of-session ritual

Before ending a working session, append a dated entry to
`docs/AGENTS_JOURNAL.md`. This is the durable memory of the project — see
the journal for the entry shape.
