---
name: squint-review
description: Squint-test the code a PR or branch changes and review it for smells and the refactorings, from Joshua Kerievsky's Refactoring to Patterns, that would cure them. Use when the user asks to review a PR or branch for code smells, duplication, or design, to run a squint test, or to refactor a PR or branch to patterns.
---

# Squint review

Review like a careful human with all afternoon: find the smells, name them, and say which named refactoring would remove each one and what that buys.

A review is **read-only**. Read the code where it already is and change nothing: no checkout, no worktree, no test run. Refactoring is a separate mode, entered only when the user asks for it.

## The five questions

Ask these of every unit in scope (class, module, file), while squinting and again at the end:

1. Does this code duplicate any behavior?
2. Does everything have just one responsibility?
3. Does this code depend on something that changes more often than it does?
4. Does everything in this class change at the same rate?
5. Is this class open for extension and closed for modification?

## Steps

### 1. Read the target

- PR number or URL: `gh pr view <pr> --json number,url,headRefName,headRefOid,baseRefName`, `gh pr diff <pr>`, and one line of `gh pr checks <pr>` for context.
- Branch: diff `<default-branch>...<branch>`.
- Nothing given: the current branch against the default branch.

Read whole files at the head revision with `git show <rev>:<path>`, fetching the PR head first if it is not local. The **scope** is every file the branch changes: production code, tests, and configuration alike. A failing check is context for the report, never a reason to stop.

Done when you have the diff and the full head-revision text of every file in scope.

### 2. Squint

Look at the code as a shape, not as text, at two distances.

- **Each file**: changes in shape (indentation drifting right, nesting, long methods) are where the conditionals live. Changes in color (runs of different syntax clustered together) mark mixed levels of abstraction.
- **The whole diff**: the same shape repeating across files is Duplicated Code, and the same edit repeated across hunks is Shotgun Surgery. The fact that this PR had to change many places for one reason is the smell.

Then read closely where the squint pointed, and ask the five questions of each unit. Record every smell by its standard name, with `path:line` and what it costs in concrete terms: the next change it makes harder.

Done when every file in scope and the diff as a whole have been squinted at, and every unit has answers to the five questions.

### 3. Recommend

Turn the smells into an ordered list of recommended refactorings, choosing each from [smells-to-patterns.md](smells-to-patterns.md).

- Small cleanups that make a pattern visible (Compose Method, Extract Method) come before the pattern refactorings. Within each group, highest cost first.
- Name each refactoring's prerequisites among the earlier ones.
- Say for each whether it goes **to** the pattern, **toward** it, or **away** from it. Toward a pattern stops once the smell's cost is gone. Away from one is for a pattern the code does not earn (Speculative Generality).
- List the smells not worth acting on in this PR, each with the reason.

The list is your working notes. The user first sees it in step 4, in the form they choose.

Done when every recorded smell is either a recommendation or listed with its reason.

### 4. Present

Your first message after reading is the question alone, in about these words: "I've read the code. Do you want everything at once, or one at a time, starting with what I'd fix first?"

- **Everything at once**: each recommendation with its smell, location, cost, refactoring, direction, and prerequisites. Then the smells not worth acting on, and one line on CI.
- **One at a time**: load the `over-coffee:coffee` skill and walk the recommendations in order, one per turn: the smell, what it costs, the refactoring you would reach for, and what it depends on. The user keeps, softens, or drops each one. Finish with the smells not worth acting on, briefly, so any can be pulled back in.

Done when the user has seen every recommendation and the list reflects what they kept.

### 5. Deliver

Offer the next moves the target allows:

- For a PR: post the kept recommendations as a review, following [post-comments.md](post-comments.md), only after the user approves the comment text.
- Refactor them: follow [refactor.md](refactor.md), only when the user asks for it.

When the user asked to refactor from the start, the kept recommendations are the plan: go straight to [refactor.md](refactor.md).
