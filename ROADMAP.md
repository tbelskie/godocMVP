# godoc Roadmap

## Overall Vision
Godoc is the **docs engineer’s Swiss Army Knife** — a platformless CLI tool that helps technical writers and teams get back to writing great documentation instead of fighting infrastructure and formatting.

## Guiding Principles
1. **Security First**  
   Security is the first constraint for every roadmap decision.
2. **Simple and Elegant**  
   Prioritize outcomes that reduce complexity for writers and teams.
3. **Focused Process**  
   Ship meaningful changes through focused PRs linked to GitHub Issues.
4. **godoc Values**  
   Platformless and self-hosted, AI-native and machine-readable, writer-first, and designed to get technical writers back to writing.

### Core Promise (MVP)
`godoc init` → Instant, beautiful, production-ready Hugo site with smart IA, AI-native features, and API documentation scaffolding.

### Phase 1.5: Support + Analytics Plumbing
- Seamless "This didn’t solve my problem" → ticket flow
- Basic insightful analytics for doc teams
- All plumbing stays self-hosted and static-first

### Phase 2: Writer Productivity + Infrastructure
- Style-guide aware commands (`godoc polish`, `godoc make selection htmltable`, etc.)
- Automatic CI/CD scaffolding (GitHub Actions first)
- Dev / preview / prod workflows
- Brownfield support (`godoc audit`, `godoc fix`)

### Long-term
Become the daily driver CLI for docs-as-code teams.

## Current Priorities
- Issue #1: Core `godoc init` (in progress)
- Issue #2: Support + analytics plumbing
- Issue #3: Writer productivity + CI/CD vision

---

*Last updated: May 22, 2026*