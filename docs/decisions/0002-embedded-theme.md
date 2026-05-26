# ADR 0002: Embedded MVP 1.0 Theme — CSS, Type, Brand Asset, Cross-Slice Placeholders

**Date:** 2026-05-26
**Status:** Accepted
**Deciders:** Tom (Founder) + Claude (Implementation)
**Supersedes:** none
**Refines:** ADR-0001 (no Node/npm in scaffolded sites)
**Related issues:** #1 Slice B, #8

## Context

Slice A of #1 shipped a scaffold that Hugo will build cleanly but
renders zero HTML pages because there are no templates. Slice B's
job is to embed the theme that makes `hugo server` produce a
visually-finished, on-brand documentation site in seconds —
without breaking ADR-0001's "single static binary, no Node, no
npm" promise.

The MVP 1.0 theme spec (provided in the Day-2 PM session) calls
for: dark-first with toggle, collapsible left sidebar, top nav
with logo and search input, system-sans typography, emerald-
accented code blocks, "Was this helpful?" widget, mobile-
responsive with hamburger, subtle indigo links/CTAs, a stacked-
pages brand mark, and a fixed brand palette
(`#0F172A` / `#6366F1` / `#10B981` / `#F8FAFC` / `#64748B`).

Three of those requirements ("Was this helpful?" widget, Pagefind
search, brand-mark asset) interact with other roadmap items —
#2 owns support/analytics flow, Slice C of #1 owns Pagefind
wiring, and the brand mark didn't previously exist in the repo.

This ADR captures the cross-cutting decisions Slice B makes that
will outlive any single PR.

## Decisions

### D1. Hand-written CSS with design tokens over pre-built Tailwind

Slice B ships a single `assets/css/main.css` (~450 lines)
organized by CSS `@layer` (reset → tokens → base → layout →
components). Brand colors are CSS custom properties on `:root`;
semantic names (`--bg`, `--text`, `--link`) resolve through them.

**Considered:** committing pre-built Tailwind output as a static
asset (the option floated in the Day-2 morning journal). Rejected
because:

1. **Footgun.** A pre-built Tailwind stylesheet only contains the
   utility classes used at our build time. Any class a downstream
   writer adds in their own markdown or layouts silently fails to
   apply. The whole "utility-first composition" pitch of Tailwind
   evaporates without a build step, and the no-Node constraint
   from ADR-0001 forbids that build step at user-init time.
2. **Bytes shipped.** A safe-list-padded Tailwind output is
   several times larger than what we actually need for this
   theme's footprint. Custom CSS is leaner and easier to audit.
3. **Audit surface.** Every line of CSS we ship to every user is
   security-relevant for the "platformless and self-hosted"
   promise. Hand-written CSS we can read in one sitting; a
   Tailwind utility dump we cannot.
4. **Override story.** Users who want to re-skin can target a
   small, named set of semantic tokens. Versus a Tailwind config
   they would need to learn and then would have no way to apply
   without `tailwindcss` on PATH.

The downside — losing utility-class composition for users — is
mitigated by exposing semantic tokens (rule #2 of godoc Values,
"Simple and Elegant"). Users opting in to Tailwind in their own
project can always do so on top of our CSS; we just don't bake it
into the scaffold.

### D2. System sans typography, no Google Fonts, no embedded font files

The font stack is `-apple-system, BlinkMacSystemFont, "Segoe UI",
Inter, system-ui, ...`. Inter is named so installed-Inter users
benefit, but **no font is loaded from the network and no font
file is embedded**.

**Considered:** linking Google Fonts for guaranteed Inter rendering.
Rejected for three reasons:

1. **Security First (rule #1).** Google Fonts is a known privacy
   surface — every page view of every scaffolded site would hit
   `fonts.googleapis.com` and pass user IP / referer / UA to
   Google. GDPR-jurisdiction sites have been fined for exactly
   this without disclosure. Defaulting every godoc user into that
   posture is unacceptable.
2. **Bytes shipped.** Inter variable is ~150 KB. System sans is
   0 bytes and renders instantly.
3. **Aesthetic.** Modern macOS / Windows / Android system sans
   stacks have closed the gap with Inter for our use case
   (medium-length technical prose). The taste cost is small.

**Considered:** self-hosting Inter WOFF2 as an embedded asset.
Deferred. A future ADR can revisit if real users report typography
complaints. We will not pre-emptively ship the bytes.

### D3. Brand mark as inline-SVG partial + separate favicon SVG

The godoc mark (stacked pages with `</>` glyph) lives in two
places by design:

- `layouts/partials/godoc-mark.html` — inline SVG using
  `stroke="currentColor"` so the mark inherits text color and
  adapts to whichever theme is active. Used in the header and the
  footer credit.
- `assets/img/godoc-mark.svg` — standalone file with brand colors
  baked in (deep-navy fill, neutral-light strokes, emerald accent
  on the rear page). Referenced as the favicon, where it must
  look right against arbitrary browser chrome.

**Considered:** a single SVG file used both inline and as favicon.
Rejected because `currentColor` doesn't work when the SVG is
referenced via `<img src>` or `<link rel="icon">`; the favicon
needs colors baked in. Two files with the same geometry is the
cleanest separation and totals under 2 KB.

**Considered:** embedding the PNG the founder supplied. Rejected.
SVG renders sharp at every size, theme-adapts automatically (in
the inline case), and is smaller than the smallest useful PNG.

The wordmark "godoc" is **not** part of the mark we ship into the
scaffold. The site's own title fills that role in the header; the
godoc name appears only in the small footer credit ("Built with
godoc"). This keeps user sites reading as the user's, not godoc's.

### D4. Visual placeholders for Pagefind and the helpful-widget

Slice B ships the **visual chrome** for two features whose
functional implementations belong to other PRs:

- **Search input** in the header — present, styled, but `disabled`
  with a `title` attribute pointing to Slice C of #1 (Pagefind
  wiring). Hidden on mobile (responsive concern deferred with the
  feature).
- **"Was this page helpful?" widget** at the bottom of every
  `single` page — present, styled with brand buttons, but
  `disabled` with a small "Feedback collection ships with godoc
  #2" footnote.

**Rationale.** Working agreement: one focused PR per slice. If
Slice B absorbed Pagefind and the helpful-widget submission flow,
it would mix three issues' worth of work into one PR and lose
review focus. But shipping Slice B without these UI affordances
would mean the rendered site looks visually incomplete on day one
— which defeats the whole "looks professional in seconds"
promise.

The compromise — ship the chrome, mark it explicitly inert —
preserves both the focused-PR rule and the first-impression
goal. When Slice C lands, it removes `disabled` from the search
input and wires the JS. When #2 lands, it removes `disabled`
from the helpful buttons and wires the submit handler. Each is a
small, focused diff against this stable chrome.

### D5. Layouts copied verbatim, not run through Go's `text/template`

Hugo template syntax (`{{ ... }}`) collides with Go's
`text/template` syntax. The render pipeline already only treats
files with the `.tmpl` suffix as Go templates; everything else is
copied byte-for-byte. Slice B adds no `.tmpl` suffix to any
layout, partial, JS, CSS, or SVG file. The only `.tmpl` files in
the scaffold are configuration and content files (`hugo.toml.tmpl`,
`llms.txt.tmpl`, `_index.md.tmpl`, …) where Go-template
interpolation of `{{.Title}}` etc. is actually wanted.

This is a convention, not a code change. Documenting it here so
future agents don't accidentally `.tmpl`-suffix a Hugo layout and
break the scaffold mysteriously.

### D6. Tiny vanilla JS for interactivity — no framework, no build

Three pieces of interactivity ship in `assets/js/theme.js`
(~50 lines, no dependencies, no build):

- Theme toggle (read/write `localStorage['godoc-theme']`, flip
  `data-theme` attr on `<html>`)
- Sidebar section collapse (toggle `aria-expanded` and `hidden`)
- Mobile hamburger (toggle a `data-open` attr on the sidebar)

An additional ~6 lines of inline `<script>` in `<head>` apply the
stored theme before the stylesheet loads, eliminating
flash-of-unstyled-content. Inline-only-because-it-must-be-inline,
and it does nothing but read `localStorage` and `matchMedia`.

**Considered:** any framework (Alpine, htmx, Stimulus, lit).
Rejected. Three trivial behaviors do not justify a runtime
dependency, even a small one. Vanilla JS keeps the page weight
and supply-chain surface minimal.

## Consequences

### Positive

- Zero third-party runtime dependencies in user sites beyond Hugo
  itself.
- Total embedded theme weight (CSS + JS + SVGs) under ~20 KB
  uncompressed, well under any reasonable budget.
- Brand is consistent and recognizable across every site `godoc
  init` produces.
- Customization story is clear: override semantic tokens in your
  own CSS, override the favicon by dropping `assets/img/godoc-
  mark.svg` in your project.
- Future agents reading this ADR + `docs/theme/BRANDING.md` can
  resume Slice B work or extend the theme without re-deriving
  context.

### Negative / risks

- We own all the CSS we ship — bug reports against the theme
  rendering go through us, not a Tailwind upstream.
- Users wanting a wildly different look will eventually want a
  first-class customization surface (`params.brand.*`), which
  doesn't exist yet.
- The disabled search input and helpful widget are visible no-ops
  until Slice C and #2 land. Mitigated by explicit `title`
  attributes and the helpful-widget footnote; user confusion
  remains a possible drive-by issue.

### Out of scope (intentionally)

- Pagefind index generation, search UI behavior, search analytics.
- Helpful-widget submission, ticket flow, support-channel
  routing.
- Per-user palette / logo / font customization through
  configuration.
- `prefers-reduced-motion` handling (motion in Slice B is too
  subtle for the cost; revisit on user feedback).
- WCAG AAA conformance (we target AA at base size in Slice B and
  will iterate on specific findings).
- A second-tier marketing landing page; the homepage layout is
  optimized for "first docs site" not "company front page".

## Notes for future ADRs

When the customization surface gets designed, it will deserve its
own ADR (ADR-0003 candidate). Likely shape: `[params.brand]` block
in `hugo.toml` with documented fields for palette overrides, logo
asset path, and wordmark text, plus a small `partial "brand-vars.html"`
that injects `:root` overrides at runtime. That work is **not**
implied by Slice B.
