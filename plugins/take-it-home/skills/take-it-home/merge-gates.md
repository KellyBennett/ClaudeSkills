# Merge gates

## Read the gates

Read the target branch's rulesets, branch protection, required checks, merge queue, required reviewers and CODEOWNERS, and the authenticated account's merge permission. For a PR, `gh pr view <pr> --json headRefOid,baseRefOid,reviewDecision,reviews,reviewRequests,autoMergeRequest,mergeStateStatus,statusCheckRollup` is useful evidence, but it does not cover every rule; read rulesets and protection through `gh api`. When a rule or permission can't be read, record it as unknown and report the restricted visibility. Unknown never means permitted.

Agent identity and GitHub identity are separate. Agents working under the author's account never satisfy an approval that must come from a different account. Satisfy gates as they stand: the account, repository settings, and protections stay as they are.

## Waiting on someone else

Finish both internal reviews, fixes, and rechecks even while external approval is missing. Then record the waiting state:

- `awaiting-external-approval`, naming the exact missing reviewer requirement.
- `awaiting-authorized-merge`, when approval is sufficient but this account or runtime cannot merge. Approval and merge permission are separate gates.

Continue independent tasks meanwhile. A dependent PR is never merged, or armed to merge unattended, before its prerequisites are merged and it has been revalidated on the resulting base.

## Merge directly

Before each merge, read the live PR: head and base, published report coverage, finding dispositions, required checks, and approvals. Merge only when the reviewed head and base still match and every acceptance criterion and required piece of feedback is satisfied. Use the repository's allowed merge method with `gh pr merge --match-head-commit <sha>`, then confirm the result. A timeout is not a merge: read back before retrying.

`--admin` and force-pushing are never part of this workflow.

## Auto-merge and merge queues

Use the repository's existing auto-merge or required merge queue only when all of these hold:

- The account has permission and the repository already offers the mechanism.
- The complete internal cycle is done on the current head and base, and every non-external prerequisite is met.
- Enforced checks would stop a later unreviewed head or base from merging. A `COMMENT` review or a checkpoint is not an enforced check. `--match-head-commit` checks the head only at request time.

When any of these can't be established, leave a manual handoff instead.

`gh pr merge --auto` waits for the repository's requirements. Record the request, its revision, and its pending requirements, and read it back. A pending request is not a merge. For a queued PR, verify actual queue membership: a null `autoMergeRequest` does not prove the PR has left the queue.

Before pushing a fix or updating the base of a PR with a pending request this run created, cancel the request or remove it from the queue. Confirm the cancellation, and check whether it already merged, so it cannot merge ahead of fresh rechecks. When cancellation is unavailable or uncertain, hold the update and reconcile the live outcome. A request owned by another actor stays theirs: reconcile it, and require a safe handoff if it could race the review cycle.

See [GitHub auto-merge](https://docs.github.com/en/pull-requests/how-tos/merge-and-close-pull-requests/automatically-merging-a-pull-request) and [protected branches](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches).
