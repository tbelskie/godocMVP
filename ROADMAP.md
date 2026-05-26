# Pendragon Roadmap

## Overall Vision

**Pendragon: The AI-powered DocOps assistant** — a platformless CLI that helps documentation teams ship, audit, and maintain docs-as-code sites without fighting infrastructure.

Brownfield is the wedge (`pendragon audit` → `fix` → `polish`). Greenfield `pendragon init` is the credibility ticket that proves taste on day one.

## Guiding Principles

1. **Security First**  
   Security is the first constraint for every roadmap decision.
2. **Simple and Elegant**  
   Prioritize outcomes that reduce complexity for writers and teams.
3. **Focused Process**  
   Ship meaningful changes through focused PRs linked to GitHub Issues.
4. **Pendragon Values**  
   Platformless and self-hosted, AI-native and machine-readable, writer-first, and designed to get technical writers back to writing.

### Core Promise (MVP)

`pendragon init` → instant, beautiful, production-ready Hugo site.  
`pendragon audit` → read-only health report on existing Hugo repos (theme-respectful).

### Phase 1.5: Support + Analytics Plumbing

- Seamless "This didn’t solve my problem" → ticket flow
- Basic insightful analytics for doc teams
- All plumbing stays self-hosted and static-first

### Phase 2: Writer Productivity + Infrastructure

- Style-guide aware commands (`pendragon polish`, selection helpers, etc.)
- Automatic CI/CD scaffolding (GitHub Actions first)
- Dev / preview / prod workflows
- Brownfield remediation (`pendragon fix`)

### Long-term

Become the daily driver CLI for docs-as-code teams.

## Current Priorities

- Issue #1: Core `pendragon init` (Slices A–C shipped; D/E sequenced per discovery)
- Issue #11: `pendragon audit` MVP (brownfield wedge — next code slice)
- Issue #2: Support + analytics plumbing
- Issue #12: ADR-0003 strategic frame (accept after #10 — done)

---

*Last updated: May 26, 2026*
