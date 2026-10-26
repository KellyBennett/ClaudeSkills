# Checkpoint and resume

The checkpoint is a small run note that lets a later session pick the work up. Keep it in an allowed task location, outside the product changes, and update it after every externally visible transition.

## Contents

- Scope, acceptance criteria, dependency edges, and ownership.
- The target checkout and its original branch, head, and modifications.
- Per task: branch, worktree, PR URL, current head and base SHAs, and state.
- Which agent filled each role.
- Check commands and results.
- Report links with the head and base each one reviewed, finding dispositions, and fix commits.
- Merge gates: required approvals, reviewer identities and review IDs, the commits they approved and whether that approval is still valid, merge permission evidence, and auto-merge or queue state with who requested it.
- Merge evidence, local update and rebuild evidence.
- The next action, or the exact blocker.

## States

`planned`, `implementing`, `published`, `reviewing`, `fixing`, `rechecking`, `merge-ready`, `awaiting-external-approval`, `awaiting-authorized-merge`, `merged`, `blocked`, `done`.

The two waiting states mean the internal review cycle is complete and an external gate remains. They are neither a failure nor completion.

## Resume

Live Git and GitHub state is the truth; the checkpoint is a claim to reconcile against it.

- After an uncertain push, review publication, or merge, read back what happened before retrying. Reuse existing PRs and reports.
- For each open PR, read new feedback, the current head and base, and whether approvals are still valid. New commits or a changed base make internal reviews stale and may dismiss external approvals under the repository's rules. Address feedback and repeat fix and recheck before merging. An approval counts only if it is valid on the live PR now.
- Read back auto-merge and queue requests before making another.
- For a PR merged while the session was paused, confirm the merged revision and record it. Reconcile any unexpected changes or missing acceptance evidence as explicit gaps. Downstream tasks now start or update against that actual merge.
- An incomplete local update or rebuild means the run is not done.
