# Take it home

Hand Claude a set of agreed tasks and it carries them the rest of the way: implementation, PRs, independent reviews, fixes, merges, and a rebuild of your local checkout.

```
/take-it-home Issues #41, #42 and #43. #43 depends on #41.
```

1. Each task gets its own agent in its own worktree, and its own PR.
2. Two fresh agents review every PR, one with [squint-review](../squint-review/README.md) and one with [magic-tricks](../magic-tricks/README.md), and post their full reports. Installing take-it-home installs both.
3. A separate agent fixes the findings, and another fresh pair rechecks, until the reviews are clean.
4. PRs merge in dependency order, then your checkout is fast-forwarded and rebuilt.

It never switches accounts, bypasses protections, or merges past a required human approval. When it has to stop, it leaves a checkpoint; run `/take-it-home` again to pick up where it left off.
