# Squint review

Refactors a PR or branch the way a careful human with plenty of time would. Claude starts with Sandi Metz's squint test from *All the Little Things*, then refactors to patterns from Joshua Kerievsky's *Refactoring to Patterns*, one named refactoring per commit, with tests green throughout.

```
/squint-review 123
/squint-review my-feature-branch
```

1. It sets up a separate git worktree on the branch, so your checkout isn't touched.
2. It squints at the changed code, names each smell, and proposes an ordered plan of refactorings.
3. Once you approve or trim the plan, it carries it out without stopping.
4. It reports each commit and the smells it deferred, and leaves pushing to you.
