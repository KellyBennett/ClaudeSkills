# Magic tricks

Audits the tests added or changed in a PR or branch against Sandi Metz's [Magic Tricks of Testing](https://www.youtube.com/watch?v=URSWYvyc42M) grid.

| Origin       | Query            | Command                     |
|--------------|------------------|-----------------------------|
| Incoming     | Assert result    | Assert public side effect   |
| Sent to self | Ignore           | Ignore                      |
| Outgoing     | Ignore           | Expect it to be sent        |

```
/audit-tests 123
/audit-tests my-feature-branch
/audit-tests
```

With no argument it audits the current branch. You choose the findings all at once or one at a time over coffee, and for a PR, Claude offers to post the ones you keep as review comments after you approve the text. Installing magic-tricks installs [over-coffee](../over-coffee/README.md) for this.
