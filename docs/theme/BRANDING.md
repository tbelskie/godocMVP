# Pendragon Theme Branding Guide

This is the living reference for the visual identity that every site
scaffolded by `pendragon init` ships with. It applies to the embedded MVP
1.0 theme in `internal/project/templates/`. Architectural decisions
behind these choices live in
[`docs/decisions/0002-embedded-theme.md`](../decisions/0002-embedded-theme.md).

> Users get Pendragon chrome out of the box. The wordmark, palette, and
> typography are Pendragon's. The content, site title, and (eventually)
> overrides are the user's.

**Product line:** Pendragon — *The AI-powered DocOps assistant.*

---

## 1. Brand essence

Pendragon is a writer-first, docs-engineer's DocOps tool. The visual identity
should feel:

- **Technical without being cold.** Deep navy primary, generous
  whitespace, no decorative flourishes.
- **Premium by default.** Subtle blur, considered type scale, smooth
  transitions. The first impression must read as "professional"
  without further configuration.
- **Calm.** One accent, one highlight. No rainbow palettes. The
  reading surface stays quiet so the writing carries the page.

---

## 2. Color palette (official)

| Role | Hex | Where used |
| --- | --- | --- |
| Primary — Deep Navy | `#0F172A` | App background (dark mode), brand surfaces, foreground text on light mode |
| Accent — Indigo | `#6366F1` | Links, CTAs, focus rings, active nav state |
| Highlight — Emerald | `#10B981` | Code-block left accent, "new" badges, hover underline emphasis |
| Neutral light | `#F8FAFC` | Foreground text on dark mode, app background on light mode |
| Neutral mid | `#64748B` | Muted text, faint metadata, dividers in light mode |

These five hex values are the **only** brand colors. Anything else in
the CSS (elevated surface `#1E293B`, code background `#0B1224`, border
`#1E293B`, link hover `#4F46E5`, etc.) is a derived semantic token, not
a brand color. Derived tokens live in `assets/css/main.css` under
`@layer tokens` and can be tuned without touching the brand list.

### Semantic tokens

Brand hexes are exposed as CSS custom properties on `:root` so that
downstream users can override either layer:

```css
:root {
  --color-primary:        #0F172A;
  --color-accent:         #6366F1;
  --color-highlight:      #10B981;
  --color-neutral-light:  #F8FAFC;
  --color-neutral-mid:    #64748B;
}
```

Layouts reference semantic names (`--bg`, `--text`, `--link`, `--bg-code`,
…) which resolve to the brand tokens above. A user who wants to re-skin
in their own colors only needs to override the semantic layer; the
component CSS doesn't need to change.

### Dark-first, light on demand

Dark is the **brand default**. The site loads dark unless:

1. The visitor has previously toggled to light (stored in
   `localStorage` under key `godoc-theme`); OR
2. The visitor's OS reports `prefers-color-scheme: light` and they have
   never toggled.

The theme toggle button in the top nav cycles between dark and light
and persists the choice. An inline `<script>` in `<head>` applies the
stored preference **before** the stylesheet loads, so there is no flash
of unstyled content. This script is intentionally tiny and inline; it
runs no third-party code and reads nothing but `localStorage`.

---

## 3. Typography

### Stack

```css
font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Inter,
             system-ui, "Helvetica Neue", sans-serif;
```

System sans first. Inter listed as a hint for installed-Inter users
but **not** loaded from the network. We do not embed font files in
the scaffold output and do not link to Google Fonts; see
`0002-embedded-theme.md` for the rationale (zero bytes shipped, no
third-party tracking surface).

### Mono stack (code)

```css
font-family: ui-monospace, SFMono-Regular, "SF Mono", Menlo,
             Consolas, monospace;
```

### Scale

| Token | Size | Use |
| --- | --- | --- |
| `--fs-h1` | `clamp(1.875rem, 1.4rem + 2vw, 2.5rem)` | Page titles (single, list) |
| `--fs-h2` | `1.5rem` | Major section headings inside content |
| `--fs-h3` | `1.25rem` | Subheadings |
| `--fs-base` | `16px` (`1rem`) | Body |
| `--fs-small` | `0.875rem` | Nav links, metadata, helpful-widget |
| Hero title | `clamp(2.25rem, 1.5rem + 3vw, 3.5rem)` | Homepage hero only |

Body line-height is `1.65`. Headings tighten to `1.25`.

### Hero gradient

The homepage `<h1>` uses a subtle vertical text gradient from
`--text` to `--color-accent` to make the first impression brand-on
without color noise on the rest of the page.

---

## 4. Logo / wordmark

Placeholder wordmark reference: [`docs/brand/pendragon-wordmark-placeholder.png`](../brand/pendragon-wordmark-placeholder.png).

Letter colors (sentence case **Pendragon**):

| Letters | Hex | Role |
| --- | --- | --- |
| **P** | `#3B82F6` | Primary brand blue |
| **e** | `#93C5FD` | Light accent blue |
| **ndragon** | `currentColor` (theme text) | Wordmark body — white on dark, navy on light |

We ship three assets:

| File | Purpose |
| --- | --- |
| `layouts/partials/pendragon-wordmark.html` | Inline SVG in header + footer |
| `assets/img/pendragon-wordmark.svg` | Full-color wordmark (dark field) for marketing exports |
| `assets/img/pendragon-mark.svg` | Favicon — **P** on `#0A0E14` |

Header layout: **Pendragon wordmark · {Site Title}**. The user's site
title stays visible so the page reads as theirs; Pendragon appears in the
footer credit ("Built with Pendragon").

---

## 5. Layout system

### Page chrome

```
┌─────────────────────────────────────────────────────────────┐
│  ☰   [mark] My Site            🔍 search…           ◐       │  ← sticky top nav (60px)
├─────────────┬───────────────────────────────────────────────┤
│             │                                               │
│  Docs       │   # Page title                                │
│   ▸ Start   │                                               │
│   ▸ ...     │   Body content…                               │
│  Guides     │                                               │
│  API        │   ─────                                       │
│  Change-    │   Was this page helpful?     [Yes]   [No]     │
│   log       │                                               │
│  Contrib.   │                                               │
├─────────────┴───────────────────────────────────────────────┤
│              Built with godoc ◐  ·  © 2026                  │
└─────────────────────────────────────────────────────────────┘
```

- **Header** is sticky, 60px tall, semi-transparent navy with
  `backdrop-filter: blur(8px)`.
- **Sidebar** is 260px on desktop, sticky to the top of the viewport
  under the header.
- **Content** is centered, capped at 760px reading width, padded
  generously top and bottom.
- **Footer** is centered, muted, single line on desktop, wraps on
  mobile.

### Sidebar behavior

- Top-level groups come from `[[menu.main]]` entries in `hugo.toml`.
  The scaffold ships five (`Docs`, `Guides`, `API`, `Changelog`,
  `Contributing`); users add or reorder by editing the menu entries.
- Each group's heading is a link to the section index page.
- When a section has child pages, a small chevron button appears next
  to the heading; clicking it collapses the child list. State is
  per-section, in-memory only (not persisted) — collapse is for
  reducing visual noise, not for hiding navigation across sessions.
- Active page gets the `is-active` class and an indigo-tinted background.

### Mobile (`max-width: 900px`)

- Sidebar collapses behind a hamburger button on the left of the header.
- Tapping the hamburger slides the sidebar in from the left over the
  content; tapping again or any nav link closes it.
- Search input is hidden on mobile in Slice B (Pagefind UX for small
  viewports is its own design problem, deferred to Slice C).

---

## 6. Components

### Code blocks

Background: `--bg-code` (deeper than the surrounding surface).
Left border: 4px `--color-highlight` emerald. This is the single most
distinctive visual flourish of the theme and intentionally consistent
across both light and dark modes.

Inline `code` gets a subtle `--bg-code` background with a small radius
and `0.92em` size.

### Links

- Color: `--link` (indigo on dark mode, accent on light mode).
- No underline at rest. On hover and focus, the link gains a
  **2px emerald underline** with `--color-highlight`, paired with a
  text color shift to `--link-hover`. This pairing — indigo link, emerald
  underline — is the second most distinctive flourish.

### Cards (homepage)

Grid of auto-fit minimum-220px cards. Each card has a 1px border in
`--border`, surface background `--bg-elevated`, and on hover lifts 2px
and switches its border color to `--color-accent`.

### Helpful widget

Bottom of every `single` page: a contained block with the prompt
"Was this page helpful?" and two `Yes` / `No` buttons. In Slice B
the buttons are `disabled` and tagged with `title` attributes
explaining that submission ships with Issue #2. The block also carries
a small "Feedback collection ships with godoc #2" note so writers
know what they're looking at.

### Skip link

A keyboard-accessible "Skip to content" link slides in from off-screen
on focus, sending users straight past the header and sidebar to
`#main`. This is non-negotiable for screen-reader / keyboard users.

---

## 7. Motion

A single easing curve: `cubic-bezier(0.2, 0.8, 0.2, 1)` exposed as
`--ease`. Used on:

- Card hover lift (`transform 0.15s`)
- Sidebar slide-in on mobile (`transform 0.2s`)
- Theme-toggle / collapse chevron rotation (`0.15s`)
- Skip-link reveal on focus (`top 0.15s`)

No keyframe animations. No scroll-jacking. No motion that respects
`prefers-reduced-motion` is added in Slice B because the motions used
are below the threshold most users find noticeable; a future revision
can add an explicit `@media (prefers-reduced-motion: reduce)` block
if real users report issues.

---

## 8. Customization story (Slice B = read-only defaults)

In Slice B, the brand is fixed. Users who want to re-skin can already:

- Override semantic CSS tokens by adding their own CSS file and
  loading it after `main.css` (Hugo asset pipeline supports this).
- Replace the favicon by adding their own `assets/img/godoc-mark.svg`
  (Hugo's union mount lets the user's `assets/` win).
- Edit menu entries in `hugo.toml` to reorder or rename sections.

A first-class `params.brand` customization surface (logo override,
palette override, theme name) is intentionally **out of scope for
Slice B**. It belongs to a later slice once we have at least one real
user asking for it.

---

## 9. Accessibility commitments

- All interactive elements have visible focus rings (`outline: 2px
  solid var(--color-accent); outline-offset: 3px`).
- Buttons that are decorative-only (theme toggle, hamburger, collapse
  chevron) have explicit `aria-label`.
- Skip-link as described above.
- Color contrast: brand palette pairs (dark text on light bg, light
  text on dark bg, indigo on dark navy) all clear WCAG AA at base
  body size; the muted-text color was deliberately picked above the
  `64748B` brand neutral to keep contrast on the navy surface.

Slice B does not claim full WCAG AAA. We expect to discover specific
issues once real users exercise the theme; those become focused
follow-up issues.
