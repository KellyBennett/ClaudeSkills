# Claude Code skills

A Claude Code plugin marketplace. Add it once and the plugins below become installable.

## Install

```
/plugin marketplace add KellyBennett/ClaudeSkills
/plugin install over-coffee@kbennett-skills
/plugin install magic-tricks@kbennett-skills
/plugin install squint-review@kbennett-skills
/plugin install readme-writer@kbennett-skills
/plugin install take-it-home@kbennett-skills
```

Every `/plugin` command also runs from a shell as `claude plugin`. `claude plugin list` should then show each installed plugin as enabled.

## Maintain

```
claude plugin marketplace update kbennett-skills
claude plugin marketplace remove kbennett-skills
```

## Plugins

### readme-writer

Writes READMEs for people with little time. Explains the value, installation, and first use in a direct, opinionated voice, with brief contribution guidance when useful. Leaves implementation details to other documents.

Run it with `/readme-writer`, or ask it to write, rewrite, or review a README. Details in the [plugin README](plugins/readme-writer/README.md).

### over-coffee

Claude drops the essay and talks like a co-worker across a table. For when you need to understand something rather than receive it.

Start one with `/coffee`. Anything you send marked `ooc:` is real instruction rather than conversation — context, steering, or `ooc: go implement that` to close the scene. Details in the [plugin README](plugins/over-coffee/README.md).

### magic-tricks

Audits the tests added or changed in a PR or branch against Sandi Metz's Magic Tricks of Testing grid: assert what comes in, expect the commands that go out, and leave everything else alone.

Run it with `/audit-tests <pr-or-branch>`. Details in the [plugin README](plugins/magic-tricks/README.md).

### squint-review

Reviews a PR or branch the way a careful human with plenty of time would. It squint-tests the changed code, names each smell with the refactoring that would cure it, and offers to post the findings as a PR review. Ask it to refactor and it carries the recommendations out, one commit per refactoring with tests green.

Run it with `/squint-review <pr-or-branch>`. Details in the [plugin README](plugins/squint-review/README.md).

### take-it-home

Carries an agreed set of tasks all the way: an agent and a PR per task, two fresh reviewers per PR running squint-review and audit-tests, fixes and rechecks until clean, merges in dependency order, then a rebuild of your local checkout. Installing it installs squint-review and magic-tricks.

Run it with `/take-it-home <tasks>`. Details in the [plugin README](plugins/take-it-home/README.md).
