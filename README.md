# Claude Code skills

A Claude Code plugin marketplace. Add it once and the plugins below become installable.

## Install

```
/plugin marketplace add KellyBennett/ClaudeSkills
/plugin install over-coffee@kbennett-skills
```

Every `/plugin` command also runs from a shell as `claude plugin`. `claude plugin list` should then show `over-coffee@kbennett-skills` as enabled.

## Maintain

```
claude plugin marketplace update kbennett-skills
claude plugin marketplace remove kbennett-skills
```

## Plugins

### over-coffee

Claude drops the essay and talks like a co-worker across a table. For when you need to understand something rather than receive it.

Start one with `/over-coffee:coffee`, or just ask Claude to walk you through something. Anything you send marked `ooc:` is real instruction rather than conversation — context, steering, or `ooc: go implement that` to close the scene. Details in the [plugin README](plugins/over-coffee/README.md).
