# Shared PR review and fix cycle

Status: prepared agent pilot, disabled and not executed. The public repository's controller test workflow passed all 11 tests in [Actions run 39154360423](https://github.com/KellyBennett/ClaudeSkills/actions/runs/39154360423/job/117754393800). That run tested the PR merge ref (`205a15f...`), not an exact source-head status or the paid agent cycle. The reported account/private-repository billing or spending limit remains a setup concern; it did not block this public controller test run. No billing settings, credentials, repository permissions or branch protections were changed for this work.

This repository already owns the actual Squint and Magic Tricks of Testing skills, including their supporting files. The reusable workflow reads those sources at an immutable shared commit. It does not substitute a short generic review prompt. The user's standing instruction to publish full reports replaces the skills' interactive presentation and posting approval steps for this workflow only. The review criteria remain intact.

## Required visible sequence

Unless the user explicitly skips this cycle for a particular task, follow this order
for every participating repository, including manual bootstrap:

1. **Implementer:** finish the initial implementation, push it, and open the PR first.
   Record its published head SHA. Do not precombine independent-review fixes into
   this initial publication.
2. **Different reviewer:** start a fresh agent context after publication. Run both
   the actual Squint review and existing Magic Tricks of Testing audit against
   that published head. Post both FULL reviews, including clean areas, set-asides
   and blockers, on the PR before a fixer starts.
3. **Yet another fixer:** use an agent distinct from the implementer and reviewer.
   Evaluate the posted findings and fix valid ones (or coordinate separate fix
   workers in the manual flow). Push visible FOLLOW-UP commits to the same PR;
   never squash, amend or force-push away the reviewed implementation. There is
   one bounded fix round. No findings means no forced edits or empty commit.
4. **Independent final reviewer:** in a fresh context distinct from implementation
   and fixing, re-run both full reviews on the latest published head and post the
   full recheck. If anything changes afterward, that review is stale.
5. Only then hand the PR to the user for final review. Unresolved findings or
   missing evidence remain explicit blockers; this sequence never authorizes merge.

The shared workflow starts at step 2, not implementation. Its snapshot requires an
already-open PR. Fresh review and fixer jobs enforce separate agent processes;
they do not require different GitHub accounts. A trusted publication job records
the exact head, report digest and returned review IDs only after every report
chunk is posted. The fixer consumes that publication artifact, not an unpublished
model result. This receipt is run-local controller evidence, not a cryptographic
attestation or proof of review quality.

A skip must be explicitly requested by the user and recorded with its scope in the
PR; missing permissions, exhausted quota or a failed run is a blocker, not a skip.
Do not claim a skipped cycle passed or bypass required checks.

## What runs

`pr-cycle.yml` is a `workflow_call` entry point. It snapshots an open, non-draft, same-repository PR. `pr-review-stage.yml` starts a fresh Codex process that reads both skills, the full changed files and diff, and the test subjects and configuration. A separate job posts the complete reports as COMMENT reviews tied to the reviewed SHA. Long reports are split without dropping content. The report includes set-asides, unresolved required findings and blockers.

If findings remain, a different Codex job evaluates them and edits the target checkout. This pilot fixes directly; it does not fan out. It has exactly one fix round. The agent has no repository write token. Its patch passes to a fresh job that applies it and runs the caller's setup and tests without an API key or write token. Another fresh job publishes the fixer's evaluation and applies and pushes the same immutable artifact. It checks for concurrent changes and never force pushes. It refuses changes to workflows, agent instructions, environment files, symlinks and submodules.

A new reviewer process runs both complete skills again on the final diff, even when the first review was clean. Remaining findings or missing evidence fail the cycle. The last job verifies the live PR head and publishes the `Independent PR review cycle` commit status on precisely that SHA. No merge or deployment occurs. The user makes the final review and merge decision.

A structured model report is evidence of what the reviewer returned, not proof of review quality. An agent process failure, malformed report, missing artifact, failed test, stale head or unresolved finding cannot produce a successful final status. There is no mock agent invocation in the workflow.

## Setup that still needs the owner

1. Confirm Actions capacity for each caller repository and resolve any applicable account or private-repository billing/spending block. Public controller tests have run successfully; paid agent execution remains untested. This change does not alter billing or buy capacity.
2. Approve an API-funded Codex setup, with access to the default Codex model and sufficient API quota. The desktop login and the GitHub connector do not authenticate an Actions runner. Provision `OPENAI_API_KEY` as a repository or approved organization secret separately. No credential is created or copied here. ChatGPT-managed auth is an advanced alternative, not implemented here, and should not be copied into this public repository's workflow.
3. Review and merge the shared workflow, then pin a reviewed full commit SHA in both `uses` and `shared-ref` of the [caller](examples/repository-caller.yml). Keep those pins identical. The official actions are SHA-pinned. Codex CLI/proxy currently use the pinned Action's default version; select and qualify a CLI release before enabling production use.
4. Configure a suitable complete setup and test command for each repository. The shared runner is Ubuntu. It does not assume that Python tests cover every project or install arbitrary runtimes for you.
5. Approve the caller's proposed `contents: write`, `pull-requests: write` and `statuses: write` permissions, and confirm Actions policy allows the pinned actions and reusable workflow. No PAT or GitHub App is required by this pilot: it uses the short-lived `GITHUB_TOKEN` for same-repository comments, statuses and pushes. Protected branches may reject its push. Do not bypass protection to make it work.
6. Set `PR_CYCLE_ENABLED=true` only after setup is approved. A disabled invocation fails explicitly. The concrete caller in this repository is manual `workflow_dispatch`; no agent is automatically invoked by this PR. Run it against an ordinary code PR after merging. Forks, drafts and protected-file fixes require today's manual cycle.

A workflow file is not enforcement. Each repository needs a caller, and required-check or ruleset configuration needs separate approval. Require the exact-head status `Independent PR review cycle` if enforcement is desired; do not rely solely on the original Actions run's check. A push using `GITHUB_TOKEN` normally does not start another PR workflow, so the workflow explicitly runs tests and final review after its own push. Other required CI checks may still need a human-triggered run. Do not weaken those checks.

## Trust boundaries and limits

Only trusted write-access actors may invoke the Codex Action; forks are rejected. Do not change the caller to `pull_request_target`, allow arbitrary bot accounts, or add secrets to job-level environment variables. Codex runs through the official Action's API proxy with `drop-sudo`; review uses `read-only`, fix uses `workspace-write`. Target checkout credentials are not persisted. Publication and push jobs do not execute PR code. Tests execute PR code on a separate disposable hosted runner with read-only repository access. Treat generated output and patches as untrusted.

The immutable shared checkout supplies the controller and prompts. The fixer can influence its own output and working files, so the fresh push job independently validates its patch. Artifacts come only from the current run. Tests and push both consume the same immutable patch artifact. Review and fix jobs use separate fresh runners and do not resume each other's sessions.

The initial base SHA is the review baseline. A moving base is not silently retested or merged into the branch. Repository merge checks still need to handle base changes. Final success is scoped to the exact PR head, not a claim that the PR has no defects or is safe to merge regardless of other checks. Protected files can be reviewed but cannot be automatically fixed. This restriction makes bootstrap review of this workflow necessarily manual.

Failures before report publication appear in Actions logs and the failing final gate; an agent that never returns cannot supply a full review. Inspect the job failure and any report artifacts and continue manually. A failing test prevents the patch from being pushed. No automatic retry can extend the one-round bound.

## Local validation and today's manual cycle

Run `python3 -m unittest discover -s automation/tests -v`. Policy tests cover stale and unsafe PRs, missing skill reports, report preservation and protected paths. Real-Git integration tests cover patch round trips and forbidden changes. These tests do not call a model or GitHub and do not claim end-to-end Actions execution.

Today, first have the implementer push the initial commit and open the PR. Then use separate agent contexts manually: a different reviewer reads both actual skills and posts both complete reports against the current SHA; yet another fixer evaluates each finding and pushes visible follow-up commits; run relevant tests; a new reviewer reads both skills again on the latest full diff and posts the full recheck. Record remaining findings and blockers. Stop after the agreed bound and leave merge to the user. Do not mark an Actions check green on the strength of this manual process.

This PR needs that manual bootstrap cycle. The new workflow cannot independently establish its own correctness before it exists on the shared branch and has approved credentials, API quota and Actions capacity in its caller repository.

## Verified execution references

- [Official Codex non-interactive execution](https://developers.openai.com/codex/noninteractive): fresh executions, structured output and automation authentication.
- [Official Codex GitHub Action](https://developers.openai.com/codex/github-action): proxy, sandbox, write-access gating and API secret requirements.
- [Pinned Action source](https://github.com/openai/codex-action/blob/bdf19a4a223ec2549a3e2274a0cf61556bc07675/action.yml): actual supported inputs used here.

