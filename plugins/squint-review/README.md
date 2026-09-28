# Squint review

Reviews a PR or branch the way a careful human with plenty of time would. Claude starts with Sandi Metz's squint test from *All the Little Things*, names each smell, and recommends the refactoring from Joshua Kerievsky's *Refactoring to Patterns* that would cure it.

```
/squint-review 123
/squint-review my-feature-branch
```

1. It reads the change where it is. The review checks nothing out, runs nothing, and changes nothing.
2. It squints at every changed file, tests and config included, and at the diff as a whole, where duplication across files shows up.
3. You choose the full report, or one recommendation at a time talked through over coffee. Installing squint-review installs [over-coffee](../over-coffee/README.md) for this.
4. For a PR, it offers to post what you kept as a review comment, after you approve the text.

Ask it to refactor and it checks out the branch and carries out the recommendations instead, one named refactoring per commit with tests green throughout. Pushing stays with you.
