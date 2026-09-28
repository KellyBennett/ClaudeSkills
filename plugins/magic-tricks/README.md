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

With no argument it audits the current branch. Findings come back in the terminal. For a PR, Claude offers to post them as review comments and waits for your go-ahead.
