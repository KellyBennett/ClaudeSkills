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

## The hook

The plugin also installs a hook that checks every `gh pr create` and `gh pr edit` before it runs. If the description skips a section or opens a section with code, the hook blocks the command and tells Claude what to fix. When a repo has its own PR template, the hook checks against that instead.

The hook is a small compiled program. The first time it runs, it downloads the build for your machine from this repo's GitHub releases. Nothing else needs to be installed. If the download fails, the hook steps aside and lets every command through.

## Releasing

Bump `version` in `.claude-plugin/plugin.json`, merge, then push a tag named `bury-the-lede-v<version>`. That tag builds the hook for each platform and publishes the release. Until the release exists, the hook lets every command through.
