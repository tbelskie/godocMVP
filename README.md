# Pendragon

**The AI-powered DocOps assistant.**

Beautiful Hugo documentation sites in seconds. Daily-driver CLI workflows for keeping docs healthy in brownfield repos.

```bash
pendragon init my-project
cd my-project
hugo server
```

## Guiding Principles

1. **Security First**  
   Security is non-negotiable. Treat every design and implementation choice as a potential attack surface.
2. **Simple and Elegant**  
   Prefer the cleanest solution that is still correct. Remove unnecessary complexity.
3. **Focused Process**  
   Significant changes flow through focused PRs linked to GitHub Issues.
4. **Pendragon Values**  
   Platformless and self-hosted, AI-native and machine-readable, writer-first, and built to help technical writers get back to writing.

## Why Pendragon

Docs engineers inherit Hugo sites that rot quietly — broken links, stale frontmatter, themes nobody wants touched. Pendragon meets teams where they are: **audit and fix at the content layer**, with a polished scaffolder when you need a credible greenfield demo.

`pendragon init` gives you:

- Smart Information Architecture (IA)
- Branded default theme with Pagefind search
- `llms.txt` and AI-native hooks
- Deploy-ready structure — no theme hunting

## Quick Start

```bash
# Install (module path still godocMVP until rename PR)
go install github.com/tbelskie/godocMVP/cmd/pendragon@latest

# Create your first docs site
pendragon init my-docs
cd my-docs
hugo server
```

After your first build, enable search:

```bash
hugo --minify
pagefind --site public
```

## Development

See `PROCESS.md` for our SDLC.

We dogfood Pendragon to document Pendragon itself.

## License

MIT (core CLI). Pro features coming soon.
