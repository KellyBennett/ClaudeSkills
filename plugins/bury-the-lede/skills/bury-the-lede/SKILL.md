---
name: bury-the-lede
description: A pull request description shape — What / Why / How, each section framed in plain words before any code. Use when opening a PR or writing or editing a PR description.
---

# PR descriptions

A PR description has three sections, in this order: `## What`, `## Why`, `## How`.

- **What** — the change as a reviewer would explain it to a colleague: what is different once this merges.
- **Why** — the problem or goal that makes the change worth merging. The ticket link goes here.
- **How** — the approach, then the specifics: identifiers, files, routes, migrations, trade-offs.

## Framing

Every section opens with framing: one or two plain sentences saying what the reader is about to look at and why it matters. Code identifiers, file paths, and other specifics follow the framing. A section's first words are prose; a backticked name or a path appears only after them.

Supporting sections (test notes, rollout steps, checklists) come after the three and open with framing the same way.

## Example

```markdown
## What

Students who sign in with Google now land on their dashboard instead of a blank page.
The redirect after sign-in was dropping the return path.

## Why

Every Google sign-in since Tuesday's release ended on a blank page, so students had to
re-enter the URL by hand. Fixes QTY-1234.

## How

The fix keeps the return path through the identity provider round trip. The callback handler
in `auth/callback.go` now reads `return_to` from the signed state cookie rather than the query
string, which the provider strips.
```
