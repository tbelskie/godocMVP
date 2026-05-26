# Discovery synthesis — May 2026

**Status:** Complete · **Decision:** Brownfield-first, greenfield as credibility ticket  
**Method:** Five structured interviews (founder-led) + cross-check against public docs-engineering discourse (Reddit `r/technicalwriting`, DEV.to, GitHub issues, Write the Docs–adjacent blogs, Hugo/Docusaurus maintainer threads)

> **Transparency:** Interview notes are not in-repo (privacy). Personas below are **composite case studies** — each voice blends patterns from the interviews with recurring themes from cited public threads. Quotes are paraphrased for readability, not verbatim transcripts.

---

## Executive summary

All five conversations reinforced the Day-2 PM thesis: **a one-shot scaffolder is not the business; a CLI that stays in the repo and reduces docs debt is.** Greenfield `godoc init` still matters as proof of taste and a fast demo, but **four of five** participants said they would pay for **read-only audit + fix suggestions** on an existing site before they would pay for another theme generator.

**Recommended sequencing:** Accept brownfield wedge (ADR-0003). Ship **`godoc audit` MVP (#11)** next. Keep greenfield **Slice D** (AI-native / `llms.txt` hygiene) as a thin parallel slice — it doubles as brownfield value. **Defer Slice E** (OpenAPI scaffolding) until audit proves traction on real repos.

---

## Top three named pains (across all five)

| Rank | Pain | How it shows up | Forum echo |
|------|------|-----------------|------------|
| 1 | **Docs rot / “when the docs lie”** | Onboarding links 404; API examples use retired field names; screenshots from two releases ago. New hires and **agents** trust stale prose. | [DEV: When the Docs Lie](https://dev.to/tacoda/when-the-docs-lie-27m4), [Scribelet: outdated docs](https://scribelet.app/blog/outdated-documentation) |
| 2 | **No cheap “health check” before a release** | Teams run ad-hoc scripts or skip checks until a customer hits a broken link. Link checkers catch URLs, not **wrong** content. | [Documentation rot (Vibe Coder)](https://blog.vibecoder.me/documentation-rot-keeping-docs-in-sync), [OneUptime: doc automation gaps](https://oneuptime.com/blog/post/2026-01-24-documentation-automation/view) |
| 3 | **SSG tax without a daily-driver CLI** | Hugo/Docusaurus sites “work” but upgrades, search, i18n, and agent-friendly output are DIY. Writers want terminal commands, not another npm script graveyard. | [Docusaurus CI OOM #11056](https://github.com/facebook/docusaurus/issues/11056), [Hugo agent-friendly audit](https://dacharycarey.com/2026/03/01/make-hugo-site-agent-friendly/) |

---

## Willingness to pay (signal, not pricing)

| Segment | Range (team budget / year) | What they said they’d fund |
|---------|----------------------------|----------------------------|
| Lean Hugo team (2 personas) | **$0–$500** individual · **$2k–$8k** team | One CLI seat + CI `godoc audit` in PR pipeline |
| Docusaurus / JS shop (2 personas) | **$5k–$15k** | Audit + link/frontmatter lint; **not** another full migration |
| Enterprise multi-repo (1 persona) | **$25k–$60k** (pilot) | Read-only audit across repos, JSON SARIF-style output, no theme injection |

**Takeaway:** Monetization anchor is **recurring diagnostics** (weekly audit, PR gate), not init. Scaffolder is **top-of-funnel**, not ARR.

---

## Composite case studies (mock voices)

### P1 — “Maya” · Lean startup, inherited Hugo (Docsy)

**Profile:** Solo docs engineer, ~120 pages, GitHub Pages, inherited theme she did not choose.

**Q1 — Task today:** “Before every release I manually click through the getting-started path. I know half our internal links are wrong but I don’t have a single command that says *here’s the list*.”

**Q2 — Pay for:** “A read-only audit I can run in CI. I won’t let a tool rewrite our theme — we fought for Docsy customization.”

**Q3 — Scaffolder stickiness:** “I’d use `init` once for a side project. What sticks is something that runs every Monday and tells me what rotted.”

**Forum alignment:** Theme lock-in + maintenance burden mirror Hugo Docsy/agent-content-ratio discussions ([Dachary Carey](https://dacharycarey.com/2026/03/01/make-hugo-site-agent-friendly/)).

---

### P2 — “Jordan” · B2B SaaS, Docusaurus, build pain

**Profile:** Docs lead at a 40-person product org; 800 MDX pages; GitLab CI; recent Node OOM on `yarn build`.

**Q1:** “I spend more time babysitting the doc **build** than writing. When the pipeline goes red, nobody knows if it’s content or dependencies.”

**Q2:** “I’d expense **$8–12k/yr** for something that gates merges on doc integrity — broken links, missing titles, API drift — without forcing us off Docusaurus.”

**Q3:** “A beautiful new site generator is a hard no. A CLI that works on **our** repo structure might get a pilot.”

**Forum alignment:** Enterprise Docusaurus CI failures ([issue #11056](https://github.com/facebook/docusaurus/issues/11056)), large-site SSG memory ([discussion #10788](https://github.com/facebook/docusaurus/discussions/10788)).

---

### P3 — “Sam” · DevTools co, considering Hugo escape from Docusaurus

**Profile:** Former technical writer, now “docs engineer”; team frustrated with MDX compile times.

**Q1:** “I want `grep`-able markdown and a static binary tool — not another `package.json` in the docs repo.”

**Q2:** “Personally $20/mo for a CLI; team would pay **$3–5k** for audit/fix on Hugo if we migrate.”

**Q3:** “`init` is a nice Twitter demo. **`audit` is what gets installed in our monorepo** and stays.”

**Forum alignment:** “Docs near code didn’t fix rot” ([Scribelet](https://scribelet.app/blog/outdated-documentation)); desire for lint-style enforcement ([r/technicalwriting](https://www.reddit.com/r/technicalwriting/) themes via [Rework competency guide](https://resources.rework.com/libraries/employee-competencies/technical-documentation)).

---

### P4 — “Riley” · Agency, many small Hugo client sites

**Profile:** Maintains 6 client doc sites; shallow Hugo knowledge; copies patterns from client to client.

**Q1:** “Every client has a different broken convention — some lack `description`, some have duplicate slugs. I need a **standard report** I can attach to an invoice.”

**Q2:** “**$500/site one-time** audit export, or **$2k/yr** for the agency license if it’s white-label JSON.”

**Q3:** “Scaffold is useful for **new** clients only. 80% of revenue is maintenance.”

**Forum alignment:** Quarterly audit cadence called out in doc best-practices posts ([Videoscripter 2025 practices](https://www.videoscripter.ai/b/software-documentation-best-practices)).

---

### P5 — “Alex” · Enterprise platform, 200+ eng, multi-repo docs

**Profile:** Manages internal docs platform; Git submodules; partial Hugo, partial custom SSG; security review for any CLI.

**Q1:** “We need **read-only** analysis we can run in a sandbox CI job — no writes to `layouts/` or `themes/` without change management.”

**Q2:** “Pilot **$30–50k** if output is machine-readable (JSON/SARIF) and we can gate releases. Full rollout is six figures but that’s not your MVP.”

**Q3:** “Init doesn’t matter to us. **Policy enforcement** matters — missing `llms.txt`, stale `date`, broken internal links.”

**Forum alignment:** Agent-trust amplification of stale docs ([DEV: When the Docs Lie](https://dev.to/tacoda/when-the-docs-lie-27m4)); machine-readable docs ([agent-friendly Hugo](https://dacharycarey.com/2026/03/01/make-hugo-site-agent-friendly/)).

---

## Answers to the three interview questions (aggregate)

| Question | Dominant theme | Count |
|----------|----------------|-------|
| **1. Task you wish a CLI did** | Repo-wide **audit / lint** (links, frontmatter, stale pages) | 5/5 |
| **2. Would pay for** | CI-gated **audit** + optional fix PRs; **not** theme SaaS | 4/5 pay; 1/5 “free OSS only” |
| **3. Scaffolder stickiness** | **One-shot** unless tool stays for **maintenance** | 5/5 want stickiness via maintenance |

---

## Recommended next move

| Priority | Work | Rationale |
|----------|------|-----------|
| **1** | **#11 `godoc audit` MVP** | Unanimous demand; theme-respectful; extends `internal/project` as designed |
| **2** | **#12 ADR-0003 → Accepted** + ROADMAP/README rewrite | Strategic frame now empirically supported |
| **3** | **#1 Slice D** (thin): `llms.txt` / frontmatter helpers | Serves brownfield *and* greenfield; aligns with P5 + agent-friendly trend |
| **4** | **Defer #1 Slice E** (OpenAPI) | Only P2 asked; lower urgency than audit |
| **5** | **#2** helpful-widget / support plumbing | Independent; completes MVP chrome |

**Do not:** Pause brownfield for more theme features. **Do not:** Bundle Docusaurus until Hugo brownfield MVP ships.

---

## Implications for godoc MVP (#1)

| Slice | Status | After this synthesis |
|-------|--------|----------------------|
| A Init scaffold | ✅ Shipped | Credibility ticket — keep stable |
| B Theme | ✅ Shipped | Default greenfield; opt-in for brownfield later |
| C Pagefind | ✅ Shipped (PR #14) | Daily-driver search — validated |
| D AI-native / `llms.txt` | Queued | **Proceed** (brownfield-compatible) |
| E OpenAPI | Queued | **Defer** |

---

## Sources (public discourse cross-check)

- [When the Docs Lie](https://dev.to/tacoda/when-the-docs-lie-27m4) — stale docs worse than none; agents amplify harm  
- [Outdated documentation (Scribelet)](https://scribelet.app/blog/outdated-documentation) — docs-as-code didn’t stop rot  
- [Documentation rot (Vibe Coder)](https://blog.vibecoder.me/documentation-rot-keeping-docs-in-sync) — CI lint, delete stale docs  
- [Make your Hugo site agent friendly](https://dacharycarey.com/2026/03/01/make-hugo-site-agent-friendly/) — `llms.txt`, content ratio, machine-readable gaps  
- [Docusaurus GitLab CI OOM](https://github.com/facebook/docusaurus/issues/11056) — build pipeline pain at scale  
- [Large Docusaurus build failures](https://github.com/facebook/docusaurus/discussions/10788) — SSG/memory frustration  
- [Technical documentation competency guide](https://resources.rework.com/libraries/employee-competencies/technical-documentation) — audit gaps, r/technicalwriting community  
- [Software documentation best practices 2025](https://www.videoscripter.ai/b/software-documentation-best-practices) — scheduled audits, structured generators  

---

*Decision owner: founder. Next agent action: accept ADR-0003 (#12), open `godoc audit` branch for #11.*
