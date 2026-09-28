# Posting findings as PR review comments

Post every chosen finding in one review, so the author gets one notification:

```
gh api repos/{owner}/{repo}/pulls/<n>/reviews --method POST --input review.json
```

`review.json`:

```json
{
  "commit_id": "<head sha>",
  "event": "COMMENT",
  "body": "Test audit against the Magic Tricks of Testing grid: <totals line>",
  "comments": [
    { "path": "<file>", "line": <line>, "side": "RIGHT", "body": "<cell>: <what the test does>. <the fix>" }
  ]
}
```

- Open an optional finding's comment with `Optional:`.
- A finding covering several tests anchors on the first and names the others.
- `event` stays `COMMENT`. Approving or requesting changes is the user's call.
- `line` must be a line in the PR diff on the head side. A finding on an unchanged line goes in the review `body` instead.
- Write `review.json` to the scratchpad directory, and show the user the comment bodies before sending.
