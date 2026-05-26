# Pendragon Cursor Rules

You are working on Pendragon — the AI-powered DocOps assistant: a security-first, simple and elegant CLI for Hugo documentation sites and brownfield doc health.

**ABSOLUTE RULES — NEVER BREAK THESE:**

1. **Security First**  
   Security is the #1 non-negotiable priority. Be extremely paranoid. Never suggest anything that could leak secrets, tokens, or create security risk. Always think about attack surfaces.

2. **Simple & Elegant**  
   Prefer the simplest, cleanest solution that is still correct. Eliminate unnecessary complexity. Write code that is obvious and a joy to read.

3. **Process**  
   All significant changes must go through focused PRs linked to GitHub Issues. Never suggest direct edits to main.

4. **Pendragon Values**  
   - Platformless and self-hosted
   - AI-native / machine-readable by default
   - Writer-first productivity
   - Help technical writers get back to writing

## Session orientation

At the **start** of every working session, read `AGENTS.md` at the repo root
and the last 2–3 entries of `docs/AGENTS_JOURNAL.md`. This is the durable
memory of the project; it's how you avoid re-deriving context from scratch.

## End-of-session ritual

Before ending a working session, append a dated entry to
`docs/AGENTS_JOURNAL.md` summarizing:

- Branches touched and PRs/commits shipped
- Issue status changes
- Key decisions and one-line rationale for each
- What the next session should pick up first
- Any blocking open questions
- Known debt worth tracking (not blocking)

This rule is non-negotiable. The journal is the project's memory; skipping
an entry breaks continuity for every future agent and human teammate.
