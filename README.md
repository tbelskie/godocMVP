# godoc

**The fastest way to beautiful Hugo documentation sites.**

One command. Zero friction. Production-ready docs in minutes.

```bash
godoc init my-project
cd my-project
godoc serve
```

## Guiding Principles

1. **Security First**  
   Security is non-negotiable. Treat every design and implementation choice as a potential attack surface.
2. **Simple and Elegant**  
   Prefer the cleanest solution that is still correct. Remove unnecessary complexity.
3. **Focused Process**  
   Significant changes flow through focused PRs linked to GitHub Issues.
4. **godoc Values**  
   Platformless and self-hosted, AI-native and machine-readable, writer-first, and built to help technical writers get back to writing.

## Why godoc

As a tech writer turned founder, I built this to solve the exact pain I lived with for years: Hugo is powerful but the setup is painful.

`godoc init` gives you:
- Smart Information Architecture (IA)
- Premium default theme (Tailwind + modern UX)
- Pagefind search
- Deploy-ready GitHub Actions
- Everything embedded — no theme hunting

## Quick Start

```bash
# Install
go install github.com/tbelskie/godocMVP@latest

# Create your first docs site
godoc init my-docs
cd my-docs
godoc serve
```

## Development

See `PROCESS.md` for our SDLC.

We dogfood godoc to document godoc itself.

## License

MIT (core CLI). Pro features coming soon.
