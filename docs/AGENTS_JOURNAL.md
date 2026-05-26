# godoc Agents Journal

Append-only session log. Each session, the agent adds one entry as its
final action. Newest entries on top. Read the last 2–3 to orient quickly.

## Entry shape

```
## YYYY-MM-DD — Short title

**Branches touched:** ...
**Shipped:** PR #N (status), PR #M (status), ...
**Issue status:** which issues moved, by how much

**Key decisions:**
- Decision + one-line rationale
- ...

**Drive-by fixes:** (optional) anything cleaned up along the way

**Next session should:**
- Concrete first action

**Open questions blocking next session:** none / list

**Known debt to track (not blocking):**
- Item + why it's not urgent
```

---

## 2026-05-26 (PM, late) — Strategic shift: brownfield-as-wedge thesis, three issues filed

**Session shape:** Post-Slice-B strategic conversation. No code shipped. The founder surfaced a sharp worry — *"we're just building a really cool, sort of self-assembling Hugo theme with some cool tools bolted on. no one is going to pay us for that"* — which is correct as stated and required a real answer, not reassurance. The conversation pivoted godoc's product thesis from "polished scaffolder" to "daily-driver CLI for docs engineers, with brownfield as the wedge."

**Branches touched:** none (no commits on `feat/embedded-theme` from this conversation; this journal entry is the only artifact).

**Shipped:**
- Issue #10 — *Interview 5 docs engineers to validate brownfield demand before Slices D/E.* The cheapest possible test of the strategic thesis. Three structured questions, target mix of Docusaurus / lean-Hugo / multi-repo enterprise interviewees. Gates #11 and #12.
- Issue #11 — *`godoc audit` MVP* spec. Read-only brownfield analyzer. Theme-respectful (never touches `layouts/`, `themes/`, `static/`). Extends `internal/project` with `Load` / `Audit` / `check.Check` interface. The Day-1 choice to name the package `project` not `scaffold` pays off here exactly as intended. Gated on #10.
- Issue #12 — *ADR-0003 brownfield-wedge* strategic ADR. Captures the product thesis explicitly so future agents and the founder don't re-derive it on every roadmap call. Provisional until #10's interviews complete.

**Issue status:**
- #1 — Slice A merged, Slice B in PR #9, Slices C–E now under strategic review pending #10
- #8 — open, closes on PR #9 merge (unchanged)
- #10 / #11 / #12 — new, filed this conversation

**Key decisions:**

1. **The MVP is the credibility ticket, not the business.** A scaffolder competes with free tutorials and free Hugo themes (Doks, Hugo Book, Lotus Docs, Docsy). One-shot tools don't generate recurring revenue. The recurring-value product is the daily-driver commands (`audit`, `fix`, `polish`, analytics) — captured in candidate ADR-0003 (#12).

2. **Brownfield is the wedge.** Far more existing Hugo docs sites than people starting new ones. Existing teams have higher willingness to pay (they have committed content, a maintenance burden, a real problem). Recurring usage is built in — `polish` runs per-PR, `audit` runs weekly. The brownfield play also defangs the SSG-agnostic question (content-layer commands are mostly engine-agnostic) and the "is the theme work wasted?" question (Slice B becomes opt-in components for brownfield users, default for greenfield).

3. **Theme-respectful is the cultural rule.** godoc never modifies a user's theme files without an explicit `--inject` flag and confirmation prompt. This is what makes godoc safe to install in any team's repo without negotiation. Three integration levels were sketched: theme-respectful (no theme touch), shortcode-additive (user opts in via `{{< godoc/* >}}` shortcodes), partial-injecting (gated, explicit consent).

4. **Validation before more code.** Three filed issues, only one of which (#10) acts immediately. #11 (`godoc audit` MVP) and #12 (ADR-0003) are deliberately parked until #10's interviews complete. The temptation to keep coding while uncertain about the strategic frame is exactly the trap that produces "cool tool no one pays for."

5. **Slice C still ships.** Even though the strategic frame is under review, Slice C (Pagefind wiring) is small, completes the greenfield demo story, and pays off in either world. Doing it in parallel with scheduling interviews is fine; doing D and E before interviews report would be a strategic mistake.

**Drive-by fixes:** none.

**Next session should (supersedes the prompt at the bottom of the earlier Day-2 PM entry):**

1. Confirm PR #9 is merged and Issue #8 closed.
2. Read Issues #10, #11, #12 and ADR-0001 / 0002 for the strategic + technical frame.
3. Open the Slice C issue (Pagefind search wiring) and branch `feat/pagefind-search` off `main`. Slice C is small and pays off in both greenfield and brownfield futures (search is a daily-driver feature); ship it.
4. **In parallel (founder work, not agent work):** start scheduling the five interviews in #10. The agent can help draft outreach copy and interview protocol if asked, but the conversations themselves are the founder's job.
5. Do **not** start Slice D (AI-native frontmatter) or E (API / OpenAPI scaffolding) until #10's synthesis is written and ADR-0003 is either Accepted or explicitly superseded.

**Open questions blocking next session:** none. #10 will surface the questions that matter; until then, Slice C is unblocked.

**Known debt to track:**
- Same as earlier Day-2 PM entry, plus:
- ROADMAP.md still leads with "Core Promise (MVP): `godoc init` → instant beautiful site". If ADR-0003 is Accepted, this needs rewriting to lead with brownfield. Captured in #12's acceptance criteria.
- The Day-2 PM entry's "Recommended fresh-session prompt for Day 3" is now stale (it had next-agent start Slice C-style work, framing greenfield as the only path). The prompt below supersedes it.

**Recommended fresh-session prompt for Day 3 (supersedes the earlier one — paste verbatim into a new Cursor agent chat in this workspace):**

```
Day 3 of godoc. Slice B (PR #9) shipped the embedded MVP 1.0 theme.
Day-2 PM strategic conversation also pivoted godoc's product thesis
from "polished scaffolder" to "daily-driver CLI; brownfield is the
wedge." Three new issues capture the strategic frame: #10
(validation interviews), #11 (`godoc audit` MVP, gated on #10), and
#12 (ADR-0003 strategic ADR, provisional until #10).

Before proposing anything, please:
1. Read AGENTS.md at the repo root.
2. Read the top 3 entries of docs/AGENTS_JOURNAL.md (the late-PM
   entry is the one that captures the strategic shift; don't miss it).
3. Read Issues #10, #11, and #12 for current strategic state.
4. Read docs/decisions/0001-architecture.md and 0002-embedded-theme.md
   for the technical constraints.
5. In 3-4 sentences, confirm where we are on #1 (which slices
   shipped), what the strategic thesis is (brownfield-led, with
   greenfield as credibility ticket), and what's unblocked vs
   blocked-pending-#10.

Then: open the Slice C issue (Pagefind search wiring), branch
feat/pagefind-search off main, and propose a focused implementation
plan. Slice C is the only unblocked code work; it pays off in both
greenfield and brownfield futures.

Do NOT start Slices D or E. They are blocked on #10's outcome.
Follow the rules in .cursor/rules/godoc.md and the working
agreements in AGENTS.md.
```

---

## 2026-05-26 (PM) — Slice B shipped: MVP 1.0 embedded godoc theme

**Session shape:** Implementation session. Founder supplied brand direction (palette, MVP-1.0 feature list, godoc logo image) mid-session; I translated it into the embedded theme, scaffolded the issue + branch + PR, and documented thoroughly so this slice is a clean handoff point.

**Branches touched:** `feat/embedded-theme` (created off `main` at `832ff4c`).

**Shipped:**
- Issue #8 (Spec for Slice B) — opened and then revised mid-session to reflect the locked brand palette, MVP 1.0 requirements, and the cross-slice placeholder decision.
- PR #9 — `feat(theme): embed MVP 1.0 godoc theme — branded, dark-first, responsive`. Open, linked to #8, closes #8 on merge, references #1 Slice B.

**Issue status:**
- #1 — Slice A merged (`a31da3a`), Slice B implementation in PR #9 awaiting review. Slices C / D / E queued.
- #6 — closed by `832ff4c` (agent continuity).
- #8 — open, will close on PR #9 merge.

**Key decisions (all captured in ADR-0002):**
- **D1 Hand-written CSS over pre-built Tailwind.** Tailwind-without-build is a footgun — pre-built CSS only contains classes used at *our* build time; any class a downstream writer adds silently no-ops. Hand-written design-tokens CSS (~615 lines, `@layer`-organized) is leaner, auditable, and the override surface is a small named set of semantic CSS custom properties.
- **D2 System sans over Google Fonts / embedded fonts.** Google Fonts is a known privacy surface (GDPR fines have happened); defaulting every godoc user into that posture is unacceptable under rule #1 (Security First). System sans ships zero font bytes and renders instantly. Inter is listed as a hint for installed-Inter users only.
- **D3 Brand mark as two SVGs.** Inline-SVG partial with `currentColor` for theme-adaptive header/footer use; standalone brand-colored SVG for favicon (where `currentColor` doesn't work). Total <2 KB.
- **D4 Visual placeholders for Pagefind + helpful-widget.** Search input and "Was this page helpful?" widget ship as styled but `disabled` chrome with explicit `title` attrs pointing to Slice C and #2. Keeps the focused-PR rule intact while still delivering the "looks finished" first impression.
- **D5 No `.tmpl` suffix on Hugo layouts.** Hugo's `{{ ... }}` syntax would collide with Go's `text/template` if `.tmpl`-suffixed. Documented as a convention so future agents don't break the scaffold mysteriously.
- **D6 Vanilla JS over framework.** Three trivial behaviors (theme toggle, sidebar collapse, mobile hamburger) do not justify a runtime dependency. 76 lines of dependency-free JS.

**Drive-by fixes:** none — kept the diff focused on Slice B.

**Verification performed:**
- `go vet ./... && go build ./... && go test ./...` clean. Three new tests pass (`TestEmbeddedLayouts_ParseAsTemplates`, `TestScaffoldBuildsWithHugo`, extended `TestCreate_WritesExpectedSkeleton`).
- End-to-end manual run: `godoc init demo-site` takes **14 ms**; `hugo --minify` renders 17 pages in 15 ms; `hugo server` returns HTTP 200 with 6–7 KB body on `/`, `/docs/`, `/docs/getting-started/`, `/guides/`, `/api/`, `/changelog/`, `/contributing/`; favicon served at `/img/godoc-mark.svg`; helpful widget present on single pages but absent on home (scoped via `eq .Kind "page"` in `baseof.html`); brand mark, theme toggle markup, and SRI-fingerprinted CSS/JS all present in rendered HTML.

**Documentation added:**
- `docs/decisions/0002-embedded-theme.md` — ADR for the six decisions above.
- `docs/theme/BRANDING.md` — living brand guide: palette + semantic tokens, type scale, logo dual-asset strategy, layout system, components (cards / code / helpful widget / skip link), motion, customization story, accessibility commitments.

**Founder-asset note:** The brand image the founder shared this session was saved to `~/.cursor/projects/Users-tom-repo-godoc/assets/image-fa6947cf-ee81-4135-81f5-265c201260be.png` (per-machine, not in git). The geometry it inspired was re-authored as an SVG and embedded; the PNG itself is not in the repo.

**Next session should:**
1. Review and merge PR #9.
2. Confirm Issue #8 closes on merge.
3. Start Slice C on a fresh `feat/pagefind-search` branch off `main`. Slice C removes `disabled` from the search input in `header.html` and wires Pagefind. See ADR-0002 D4 for the contract; the visual chrome is already in place.
4. In parallel, #2 (support + analytics) can pick up the helpful-widget submission flow against the disabled chrome at the bottom of every single page. Independent of Slice C; either order works.

**Open questions blocking next session:** none.

**Known debt to track (not blocking, unchanged from morning except where noted):**
- Cursor rule file format: `.cursor/rules/godoc.md` still lacks YAML frontmatter; consider renaming to `.mdc` with `alwaysApply: true` before Slice C to make auto-loading reliable for fresh agents. Workaround (paste-the-orientation-prompt) is documented and proven to work.
- Module path is still `github.com/tbelskie/godocMVP`. Branding fix is a tiny separate PR; ideally before Slice C.
- No CI workflow yet (`.github/workflows/`). Slice B grew the test surface (real-Hugo integration test). Worth landing a small CI workflow before the test count grows further. Note: the Hugo integration test correctly `t.Skip`s when Hugo isn't on PATH, so a minimal Ubuntu-Go-only runner is fine for now; richer matrix (Hugo installed) becomes valuable later.
- `prefers-reduced-motion` is not yet respected in the theme. Motion is subtle; revisit on user feedback. Captured in ADR-0002 "out of scope".
- First-class theme customization surface (`params.brand.*`) is not yet designed. Will deserve its own ADR-0003 when a real user asks for it. Captured in ADR-0002 "notes for future ADRs".
- A leftover `my-docs/` directory from Day-2 morning verification still sits untracked in the working tree. Not in any commit; safe to `rm -rf` whenever convenient.

**Cross-references for the next agent:**
- `docs/theme/BRANDING.md` is the brand guide; read this before touching any layout or CSS.
- `docs/decisions/0002-embedded-theme.md` is the ADR — read this before deviating from any decision it captures (Tailwind, fonts, JS framework, `.tmpl` suffixes, etc.).
- `internal/project/templates/layouts/partials/sidebar.html` reads `[[menu.main]]` from `hugo.toml`. To add a section to the sidebar, add a menu entry; the sidebar will pick it up automatically and (if the matched section has child pages) render a collapsible group.

**Recommended fresh-session prompt for Day 3 (paste verbatim into a new Cursor agent chat in this workspace):**

```
Day 3 of godoc. Slice B (PR #9) ships the embedded MVP 1.0 theme with
visual placeholders for Pagefind search (Slice C) and the helpful
widget (#2). Once #9 merges, the next focused PR should be Slice C:
wire Pagefind to the existing search input.

Before proposing anything, please:
1. Read AGENTS.md at the repo root.
2. Read the top 2 entries of docs/AGENTS_JOURNAL.md.
3. Read docs/decisions/0002-embedded-theme.md (the constraints
   you must respect) and docs/theme/BRANDING.md (the visual surface
   Slice C plugs into).
4. In 3-4 sentences, confirm what you understand about where we
   are, what Slice C's scope is, and the constraints from ADR-0002
   (especially: no Node/npm at init time, no third-party CDN, the
   search input chrome is already in place in header.html).

Then open a new GitHub Issue spec'ing Slice C, branch
feat/pagefind-search off main, and propose a focused implementation
plan before writing code. Follow the rules in .cursor/rules/godoc.md.
```

---

## 2026-05-26 (AM) — Slice A verified end-to-end against real Hugo

**Session shape:** Morning verification + handoff preparation. No new code; the deliverable was confidence that what shipped yesterday actually works against the real Hugo runtime, not just unit tests.

**Branches touched:** read-only on `feat/init-scaffold`; this entry committed on `feat/agent-continuity`.

**Verification performed:**
- Built `feat/init-scaffold` binary (`go build ./cmd/godoc`), ran `godoc init demo-site` into a clean tempdir.
- Ran `hugo` (build, no server) against the scaffold output: exit 0, produced `sitemap.xml`, `robots.txt`, `index.xml` RSS feeds, and category/tag indexes — but **no HTML pages**. This is correct Hugo-by-design behaviour: with no theme and no templates, Hugo refuses to invent HTML and silently emits zero pages.
- Added three throwaway HTML templates (`layouts/index.html`, `layouts/_default/single.html`, `layouts/_default/list.html`) and rebuilt. Hugo produced **9 HTML pages**: homepage, all five section indexes (`/docs/`, `/guides/`, `/api/`, `/changelog/`, `/contributing/`), `/docs/getting-started/`, and the taxonomy indexes. Homepage rendered title "Demo Site" (proving `titleFromName` works in a real Hugo run), markdown content (bold, code, lists) rendered correctly, cross-links resolved.
- Ran `hugo server --port 1314` and curled all routes: HTTP 200 across the board.

**Verdict:** Slice A ships exactly what it claimed — a valid Hugo project. The "404 on every URL" issue the founder hit at the start of the session is **not a scaffolder bug**; it is Hugo correctly refusing to render without templates. Slice B's whole job is to ship those templates.

**Mistake corrected:** Day-1 verification guidance said the founder would see "Hugo's default index rendering" — there is no such thing. Hugo has no built-in HTML rendering; it ships zero default layouts. Future advice should say: "Hugo will start the server, but every URL will 404 until Slice B ships layouts; the real verification is whether `hugo` builds cleanly against the scaffold output." That's the test we passed.

**Implications for Slice B scope (now grounded in evidence):**
- Three layout files (`index.html`, `_default/single.html`, `_default/list.html`) are the **absolute minimum** to make `localhost:1313` render. Slice B must ship at least these.
- For a "breathtaking in seconds" experience, also need: a base layout (`_default/baseof.html`), a partial for nav/header, a partial for footer, and an `assets/css/` with **pre-built Tailwind CSS** committed as a static asset.
- Hard constraint: **no Tailwind build step at `godoc init` time**. The project's "single static binary, no Node, no npm" promise (ADR-0001) requires the CSS to be pre-built and embedded.
- Sub-second `godoc init` budget still holds — pre-built CSS is just a file copy.

**Issue status:**
- #1 — Slice A verified, ready to merge.
- #6 — PR #7 still open, awaiting review.
- #5 — open, awaiting review and merge.

**Next session should:**
1. Confirm merges happened on `main` for #5 and #7.
2. Start Slice B on a fresh `feat/embedded-theme` branch off `main`.
3. Open a new GitHub Issue for Slice B (template: Spec). Title suggestion: "Slice B: embedded minimal theme — base layout, partials, prebuilt Tailwind CSS".
4. Implementation outline (proposal — agent should validate):
   - Add `internal/project/templates/layouts/_default/baseof.html`
   - Add `internal/project/templates/layouts/_default/single.html`
   - Add `internal/project/templates/layouts/_default/list.html`
   - Add `internal/project/templates/layouts/index.html`
   - Add `internal/project/templates/layouts/partials/header.html`, `footer.html`, `nav.html`
   - Add `internal/project/templates/assets/css/main.css` (pre-built Tailwind output)
   - Update `templates.go` to embed dotfile-safe paths (already handled via `all:` prefix)
   - Update `hugo.toml.tmpl` to wire the asset pipeline if needed
   - Add a test that builds the scaffolded project through real Hugo and asserts non-empty HTML in `public/index.html`
5. Acceptance criterion for Slice B: `godoc init my-site && cd my-site && hugo server` renders a styled, multi-page documentation site that a stranger would describe as "looks professional" without further configuration.

**Open questions blocking next session:** none. (Founder gave explicit green light to start Slice B.)

**Known debt to track:**
- Cursor rule file format risk: `.cursor/rules/godoc.md` has no YAML frontmatter and may not auto-load into new agent sessions. The fresh-session prompt below works around this; a tiny follow-up PR to rename to `.mdc` with `alwaysApply: true` frontmatter would make the loop fully automatic. Worth doing before Slice C.
- Module path is still `github.com/tbelskie/godocMVP` (branding drift); separate tiny PR.
- No CI workflow yet (`.github/workflows/`); local-only testing. Should land before Slice B grows the test surface.

**Recommended fresh-session prompt for Day 2 PM (paste verbatim into a new Cursor agent chat in this workspace):**

```
Day 2 of godoc. Slice A of #1 is verified end-to-end against real Hugo and
ready to merge. Now we're starting Slice B: an embedded minimal theme so
godoc init produces a visually finished, themed site in seconds.

Before proposing anything, please:
1. Read AGENTS.md at the repo root.
2. Read the top 2 entries of docs/AGENTS_JOURNAL.md.
3. In 3-4 sentences, confirm what you understand about where we are on #1,
   what Slice B's scope is, and the project's working agreements.

Then open a new GitHub Issue spec'ing Slice B, branch feat/embedded-theme off
main, and propose a focused implementation plan before writing code.
Follow the rules in .cursor/rules/godoc.md.
```

---

## 2026-05-25 — Agent continuity PR shipped

**Session shape:** Brief evening session to verify, push, and open the PR for the continuity work that a parallel agent had already committed locally (`2ae5869`).

**Branches touched:** `feat/agent-continuity`

**Shipped:**
- PR #7 — agent continuity system (the "PR for #6" referenced in the morning entry). Open, linked to #6, will close it on merge.

**Issue status:** unchanged. Roadmap unchanged.

**Key decisions:**
- Quality-checked the parallel agent's three files before pushing — `AGENTS.md`, `docs/AGENTS_JOURNAL.md` Day-1 entry, and `.cursor/rules/godoc.md` rule additions are all accurate and well-structured; shipped without modification.
- Did not gate this PR on fixing the pre-existing build break in `main` (the `cmd/godoc` package-mixing issue that #5 fixes). This PR is docs-only; `go vet/build/test` failure on this branch is a function of the base, not the diff. Called out in the PR description so the reviewer isn't surprised.

**Drive-by fixes:** none.

**Next session should:**
1. Watch for review on #5 and #7. Independent diffs — either order of merge works.
2. After both merge, rebase any remaining work onto `main` and start Slice B on a fresh `feat/embedded-theme` branch. Slice B is the first slice that produces a visibly premium site on `hugo server` — first demoable moment.

**Open questions blocking next session:** none.

**Known debt to track:** unchanged from the morning entry.

---

## 2026-05-25 — Day 1: Foundations + godoc init Slice A

**Branches touched:** `main`, `feat/guiding-principles`, `feat/init-scaffold`, `feat/agent-continuity`

**Shipped:**
- PR #4 (merged) — Guiding Principles: `.cursor/rules/godoc.md` + new sections in README/ROADMAP/PROCESS
- PR #5 (open, review-ready) — `godoc init` Slice A: scaffold engine + embedded skeleton
- PR for #6 (this branch) — agent continuity system: AGENTS.md + this journal + rule update

**Issue status:**
- #1 (`godoc init` spec) — Slice A complete; B–E queued
- #6 (agent continuity) — implementation in this branch

**Slice plan for #1:**
- A: scaffold engine + embedded skeleton (done, awaiting merge)
- B: embedded minimal theme (`layouts/`, `assets/`, Tailwind via prebuilt CSS)
- C: Pagefind static search wiring
- D: AI-native enrichment (frontmatter helpers, structured `data/`, expanded `llms.txt`)
- E: API section: OpenAPI 3.1 template + endpoint shortcodes

**Key decisions:**
- Package named `internal/project` (not `scaffold`) — leaves room for future `Load`/`Audit`/`Fix` operations on existing projects without renaming.
- Strict ASCII-only project name allowlist. Avoids Unicode normalization, filesystem case-folding, and shell-quoting pitfalls. Can be relaxed later if real users need it.
- File layout: `project.go` (public API + safety) / `render.go` (writeSkeleton + walker + per-file write) / `name.go` (validation + title) / `templates.go` (embed FS).
- Pure stdlib for the scaffold engine — no new third-party dependencies. Minimizes attack surface.
- Atomicity via `O_EXCL` writes + success-flag deferred rollback. Half-written projects are impossible by construction.
- Errors echo the offending input (`"my-project!" contains invalid character '!'`) — most common UX win for CLI validation.
- Adopted layered continuity system over canvas-based or transcript-based memory: `AGENTS.md` + journal + rule. Token-efficient, in-git, multi-machine, audit-able.

**Drive-by fixes:**
- `cmd/godoc` previously mixed `package main` (main.go) with `package godoc` (root.go, init.go) in the same directory — Go cannot compile that. Unified under `package main`. The original skeleton commit had never been built.
- `go.mod` was missing cobra's indirect deps; `go mod tidy` materialized them. No new third-party deps; only what cobra already required.

**Next session should:**
1. If PR #5 is unmerged, monitor and rebase if needed.
2. If PR for #6 is unmerged, same.
3. Once both are merged, start Slice B on a fresh branch `feat/embedded-theme` off `main`. Slice B is the first slice that produces a visibly premium site on `hugo server` — first demoable moment.

**Open questions blocking next session:** none.

**Known debt to track (not blocking):**
- `go.mod` module path is still `github.com/tbelskie/godocMVP` — branding fix not yet applied. Should be a tiny, separate PR.
- No CI workflow yet (`.github/workflows/`); tests run locally only. Worth a focused PR before slice B grows the test surface.
- No release / version-injection pipeline; `version` in `cmd/godoc/root.go` is hardcoded `"dev"`.
- `feat/guiding-principles` branch is leftover after PR #4 merged — safe to delete locally and on origin.
- `cmd/godoc/init.go` has an unused-looking `name := args[0]` extraction; intentional (used twice). Not debt, just noting.

**Artifacts you can look at outside the repo:**
- `~/.cursor/projects/Users-tom-repo-godoc/canvases/day-1-founders-report.canvas.tsx` — Day-1 founder dashboard. Per-machine, not in git. Not a continuity mechanism — that's what this journal is for.
