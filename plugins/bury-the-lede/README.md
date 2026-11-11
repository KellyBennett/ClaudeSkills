# Bury the lede

Agent-written PR descriptions tend to bury the lede: they open on query names, packages, and file paths, so a reviewer meets the details before knowing what they're looking at. This skill fixes the order: What, Why, How, and every section starts with a plain sentence before any code.

## How to install it

In Claude Code, run:

```
/plugin marketplace add KellyBennett/ClaudeSkills
/plugin install bury-the-lede@kbennett-skills
```

If you've already added the marketplace, update it with `/plugin marketplace update kbennett-skills` before installing.

## How to use it

It kicks in whenever Claude opens a PR or writes a PR description. You can also call it directly:

```
/bury-the-lede Write the description for this branch.
/bury-the-lede Rewrite the description on PR 42.
```
