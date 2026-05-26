# ADR 0004: Pagefind Static Search Integration

**Date:** 2026-05-26
**Status:** Accepted
**Deciders:** Tom (Founder) + Cursor agent (implementation)
**Refines:** ADR-0001 (no Node/npm), ADR-0002 D4 (search placeholder contract)
**Related issues:** #1 Slice C, #13

## Context

Slice B shipped a disabled search input in the header (ADR-0002 D4) as
visual chrome. Slice C's job is to make every `godoc init` site
genuinely searchable after one post-Hugo command, without breaking the
platformless, self-hosted, no-CDN constraints.

## Decisions

### D1. Pagefind over a custom Go-side indexer

Pagefind reads `public/` after `hugo` and writes `public/pagefind/`.
It is purpose-built for static sites, fuzzy-matches in the browser via
WASM, and requires no backend. A custom indexer inside the godoc binary
would duplicate that work, add maintenance surface, and couple search
quality to our release cycle.

### D2. Same-origin assets only — no CDN

`pagefind-ui.css` and `pagefind-ui.js` load from `/pagefind/...` on
the site's own origin. No third-party network calls from the rendered
page (Security First, ADR-0002 D2 ethos).

### D3. User runs `pagefind` manually for MVP — no `godoc build` wrapper

Indexing is documented in the seeded getting-started page:

`hugo --minify && pagefind --site public`

We do not shell out to Pagefind from the godoc binary in this slice.
Demand for `godoc index` can be validated later without blocking Slice C.

### D4. Pagefind UI library over a hand-rolled results dropdown

We mount `PagefindUI` on `#godoc-search` and skin it via CSS custom
properties mapped to godoc semantic tokens. The library owns keyboard
navigation, ARIA, and result rendering (~70 KB gzipped, same-origin).
A hand-rolled UI would be a separate slice if users reject the default
chrome.

### D5. Graceful degradation when the index is missing

If `pagefind` has not been run, `/pagefind/pagefind-ui.js` 404s,
`PagefindUI` is undefined, and the header keeps a styled fallback
`<input>` with a `title` hint describing the indexing step. The rest
of the site renders normally.

### D6. Index only article body; exclude chrome

`data-pagefind-body` on `<article>`, `data-pagefind-ignore` on the
helpful-widget wrapper, `data-pagefind-meta="title"` on page `<h1>`
for clean result titles.

## Consequences

### Positive

- Search works on any host that serves static files (GitHub Pages, S3,
  nginx, `hugo server`).
- Zero new Go dependencies; `godoc init` stays sub-second.
- Brownfield sites can adopt the same Pagefind + CSS-variable pattern
  without godoc touching their theme files.

### Negative / risks

- Users must install Pagefind separately until we have demand to bundle
  or wrap it.
- Mobile search remains hidden ≤900px (Slice B contract; unchanged).
- Subpath `baseURL` deployments may need path adjustments — not tested
  in MVP.

### Out of scope

- `godoc build` / `godoc index` commands
- Bundling the `pagefind` binary in godoc releases
- Search analytics, `/search/` landing page, multi-language indexes
