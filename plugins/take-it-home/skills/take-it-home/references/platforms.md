# Codex and Claude Code

Read the section for the active runtime before dispatching agents. Discover its available capabilities rather than inventing tool names. Keep orchestration in the main conversation. Do not launch another paid CLI runtime, change model choices, or enable experimental features merely because a tool is missing.

## Codex

The portable skill is `skills/take-it-home`, including its references and `agents/openai.yaml`. For local discovery, place that whole folder in a supported skill directory such as the target repository's `.agents/skills/take-it-home`; invoke `$take-it-home`. Installation is a separate user action. Do not overwrite an existing skill. This repository's Claude marketplace manifest does not itself prove Codex installation or dependency resolution.

Use the current session's native subagent tools. In a runtime exposing `spawn_agent`, start reviewers with a fresh context (`fork_turns: "none"` where supported), and give each its explicit task, skill locations and revisions. Do not assume spawning isolates files. Prepare separate Git worktrees for writers using the available worktree tool or `git worktree add`, then pass each writer its absolute worktree path and verify its branch and HEAD before edits. Preserve user model choices and the runtime's defaults. This skill does not require custom agent configuration files.

Locate actual `squint-review`, `audit-tests` and their presentation resources through skill discovery or installed paths. Prefer installed Codex adaptations when present. If using this repository's Claude-oriented sources, use the same review criteria and files; map `AskUserQuestion` to the available user-input tool or chat, and resolve `over-coffee:present-findings` to the actual presentation skill. Preserve its approval requirements. Never fabricate unavailable tools or treat a UI adaptation as permission to skip review.

See [Codex skill discovery](https://developers.openai.com/codex/build-skills) and [subagents](https://developers.openai.com/codex/agent-configuration/subagents).

## Claude Code

Use the repository's marketplace mechanism to install `take-it-home@kbennett-skills` when the user requests installation. The manifest declares `squint-review` and `magic-tricks` dependencies. Verify that their skills and the transitive `over-coffee` resources are actually available. The fully qualified invocation is `/take-it-home:take-it-home`; the skill slug is `take-it-home`.

Use the native Agent tool and fresh calls, not `resume`, for each review pass. For implementation, request worktree isolation in the call when supported and verify the returned worktree. Otherwise prepare a Git worktree and require the writer to operate there. Do not assume an agent name, a background session, or a new context provides isolation. Pass actual skill names or source paths into each task; do not assume the parent skill's content is inherited. Reviewers read existing revisions without creating worktrees or running tests. The main conversation handles publication and any required interactive confirmation.

Read applicable `CLAUDE.md`, `.claude/skills`, and repository instructions as well as the shared workflow. Use the configured model and permission mode. Do not enable agent teams or bypass permissions for this workflow.

See [Claude skills and plugin invocation](https://code.claude.com/docs/en/skills), [subagents and worktree isolation](https://code.claude.com/docs/en/sub-agents), and [plugin manifests](https://code.claude.com/docs/en/plugins-reference).

## Capability limits

A valid package proves neither agent independence nor external permissions. If the runtime cannot provide fresh agents, revision access, publication, or required checks, record the exact unavailable capability and block the dependent action. Continue independent authorized work. Test package discovery separately from model execution; never describe local fixtures or simulated PR events as a completed run in both products.

## External GitHub gates in either runtime

Use the connected GitHub tools or authenticated `gh` to inspect the target rules and the current account's permissions. Agent identity is separate from GitHub account identity. `gh pr view <pr> --json headRefOid,baseRefOid,reviewDecision,reviews,reviewRequests,autoMergeRequest,mergeStateStatus,statusCheckRollup` supplies useful PR evidence, but does not alone prove complete rules or permission coverage. Read applicable rulesets, branch protection and queue requirements through the available API; report restricted visibility.

When the shared workflow's prerequisites hold, use the repository's allowed merge method and existing auto-merge/queue mechanism. With `gh`, `--auto` waits for repository requirements and `--match-head-commit <sha>` checks the head at the request; the latter does not pin all future automated merges. Never use `--admin` or switch accounts. Read back the request and its outcome. For a queued PR, verify actual queue membership and removal; a null `autoMergeRequest` does not prove it has left the queue. If no permitted mechanism is available, return the approval/merge handoff and checkpoint, rather than installing a bot or enabling repository features.

See [GitHub auto-merge permissions and behavior](https://docs.github.com/en/pull-requests/how-tos/merge-and-close-pull-requests/automatically-merging-a-pull-request) and [protected branches and stale reviews](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches).
