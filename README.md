# Claude Code skills

A Claude Code plugin marketplace. Add it once and the plugins below become installable.

## Install

```
/plugin marketplace add KellyBennett/ClaudeSkills
/plugin install over-coffee@kbennett-skills
/plugin install magic-tricks@kbennett-skills
/plugin install squint-review@kbennett-skills
```

Every `/plugin` command also runs from a shell as `claude plugin`. `claude plugin list` should then show each installed plugin as enabled.

## Maintain

```
claude plugin marketplace update kbennett-skills
claude plugin marketplace remove kbennett-skills
```

## Plugins

### over-coffee

Claude drops the essay and talks like a co-worker across a table. For when you need to understand something rather than receive it.

Start one with `/coffee`. Anything you send marked `ooc:` is real instruction rather than conversation — context, steering, or `ooc: go implement that` to close the scene. Details in the [plugin README](plugins/over-coffee/README.md).

### magic-tricks

Audits the tests added or changed in a PR or branch against Sandi Metz's Magic Tricks of Testing grid: assert what comes in, expect the commands that go out, and leave everything else alone.

Run it with `/audit-tests <pr-or-branch>`. Details in the [plugin README](plugins/magic-tricks/README.md).

### squint-review

Reviews a PR or branch the way a careful human with plenty of time would. It squint-tests the changed code, names each smell with the refactoring that would cure it, and offers to post the findings as a PR review. Ask it to refactor and it carries the recommendations out, one commit per refactoring with tests green.

Run it with `/squint-review <pr-or-branch>`. Details in the [plugin README](plugins/squint-review/README.md).
