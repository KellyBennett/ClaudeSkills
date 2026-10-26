---
name: take-it-home
description: Carry an agreed, bounded set of tasks through implementation, independent squint and test reviews, fixes, dependency-ordered merges, and a local rebuild.
disable-model-invocation: true
---

# Take it home

Finish work that is already defined: implement each task, publish a PR, have it reviewed independently, fix what the reviews find, merge in dependency order, then update and rebuild the user's checkout. Task discovery and standing auto-merge policy belong elsewhere.

Invoking this skill authorizes, for the agreed tasks only, creating branches, worktrees and PRs, posting both full review reports, fixing their findings, and merging. Runtime permissions, repository rules, required checks, and the user's explicit limits still bind. Where a loaded skill requires approval of comment text, prepare the reports and ask once. A real unmet permission or scope decision is a **blocker**: record it exactly, and keep independent work moving.

## Roles

The main conversation is the **orchestrator**. It plans, dispatches, publishes, merges, and owns the [checkpoint](checkpoint.md). Every other role is a fresh Agent call. A new agent knows only its prompt, so pass it the task, the skill names, and the exact revisions.

- **Implementer**: one per task, dispatched with `isolation: "worktree"`. Verify the returned worktree's branch and HEAD before trusting its work.
- **Reviewers**: two per pass, distinct from every implementer and fixer of that PR. One runs `squint-review:squint-review`, the other `magic-tricks:audit-tests`. Both are read-only: they read revisions where they are, with no worktree and no test run. Tell them the user chose the full report and that they return it to you; publication and any confirmation happen in the orchestrator.
- **Fixer**: distinct from the implementer and that pass's reviewers.

Independence is the point of the roles. Record which agent filled each one, and when a fresh context is unavailable, block that review.

## Steps

### 1. Establish the run

Read the repository's AGENTS.md or CLAUDE.md and its documented check and rebuild commands. Confirm `squint-review:squint-review` and `magic-tricks:audit-tests` are available. When either is missing, record it and block review and merge; planning and implementation can still go ahead.

Record the local target checkout's branch, head, and existing modifications, changing nothing. Look for existing branches, PRs, and a prior checkpoint before creating anything. When a checkpoint exists, resume from it following [checkpoint.md](checkpoint.md).

Write acceptance criteria for each task and a dependency graph. Give each task the files or components it owns. Tasks with overlapping ownership, or that depend on each other, run in sequence; independent tasks may run concurrently.

Read the target branch's merge gates following [merge-gates.md](merge-gates.md).

Done when the checkpoint holds the scope, acceptance criteria, dependency edges, ownership, the target checkout's original state, and each merge gate or the fact that it is unknown.

### 2. Implement and publish

Dispatch an implementer per ready task with its acceptance criteria, owned paths, dependencies, and the repository's instructions. Start a dependent task from its merged prerequisite. Stack it on an unmerged one only with an explicit base and a plan to retarget and revalidate once the prerequisite merges.

Implementers run the repository's checks and report results and gaps. Unrelated user changes stay exactly as they are. On failure, keep the commits and worktree for resuming.

Open each PR against its correct base with its validation evidence. A failing required check blocks that PR: investigate and fix within scope while other independent work continues.

Done when every ready task has a PR whose head, base, and check results are in the checkpoint.

### 3. Review

Capture the PR's exact head and base SHAs, then start both reviewers in parallel with the full diff and those revisions.

Publish each complete report as a `COMMENT` review pinned to the reviewed head, or as a PR comment naming head and base when a review can't be submitted. Publish everything: findings, clean results, set-asides, and gaps. Split a long report into parts without losing content. Before posting, save the report and the parts you expect; after posting, record the returned IDs and confirm each part landed. After a timeout, read what was posted and post only the parts proven missing.

Done when both full reports for the current head and base are published and confirmed.

### 4. Fix and recheck

Send both posted reports to a fixer. It evaluates every finding and puts substantive fixes in new commits on top of the implementation. Every finding gets a disposition with evidence: fixed, or not applicable with the reason. A required finding needs a fix or independent evidence it does not apply. A disputed required finding blocks merge until the user decides. Clean reports need no fixer.

Then dispatch a fresh pair of reviewers, distinct from every author and fixer, on the complete latest diff. Publish their full rechecks as in step 3, and have them verify each original finding against the code and tests. Fetch every PR comment and review thread, human feedback included, and account for each actionable one. Resolve a thread only after its fix is verified.

Repeat fix and recheck while each round makes measurable progress. Block the task, recording the exact decision needed, when a finding recurs without new evidence, reviewers disagree on a required change, a fix would expand scope, or a runtime or cost limit is reached.

Done when the latest recheck on the current head and base is clean, or every finding has a verified disposition.

### 5. Merge

Merge in dependency order, following [merge-gates.md](merge-gates.md). A new commit or a changed base makes a PR's reviews stale: rerun its checks and step 4's recheck on the new head and base before merging it.

Done when every PR is merged or waiting on a recorded external gate.

### 6. Bring it home

Once every task is confirmed merged, update the recorded target checkout with a fast-forward-only pull. First check its branch, dirty and untracked state, and divergence again, and compare the incoming paths with local edits and untracked paths, directory conflicts included. On any overlap, leave the checkout as it is and report the conflicting paths. The update happens in the user's checkout on the user's branch; a separate clean checkout does not count. Then run the documented rebuild there.

Done when the checkout is at the merged head and the rebuild succeeded, or the blocker is recorded.

## Report

Report **done** only with evidence that all scope is merged and the target checkout is updated and rebuilt. Otherwise report partial completion: each PR with its reviewed head and base, report links, missing approval or merge permission, dependency constraints, auto-merge state, the next authorized action, and where the checkpoint is. The run ends with the session, so tell the user to invoke `/take-it-home` again after an external review or manual merge.

Keep the checkpoint and any worktree holding unpreserved work. Remove only disposable worktrees this run created.
