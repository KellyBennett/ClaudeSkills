# Posting the review

Post every chosen recommendation in one review, so the author gets one notification:

```
gh api repos/{owner}/{repo}/pulls/<n>/reviews --method POST --input review.json
```

`review.json`:

```json
{
  "commit_id": "<headRefOid>",
  "event": "COMMENT",
  "body": "Squint review: <one-paragraph summary, plus the smells not worth acting on>",
  "comments": [
    { "path": "<file>", "line": <line>, "side": "RIGHT", "body": "<smell>: <what it costs>. <refactoring, direction, and what it buys>" }
  ]
}
```

- Open an optional recommendation's comment with `Optional:`.
- `event` stays `COMMENT`. Approving or requesting changes is the user's call.
- `line` must be a line in the PR diff on the head side. A recommendation spanning files anchors on its first occurrence and names the others. One on an unchanged line goes in the review `body`.
- Write `review.json` to the scratchpad directory, and show the user the comment bodies before sending.
