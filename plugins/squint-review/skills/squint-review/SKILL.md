---
name: squint-review
description: Squint-test the code in a PR or branch and refactor it to patterns in small green steps, the way Joshua Kerievsky's Refactoring to Patterns lays out. Use when the user asks to refactor a PR or branch to patterns, run a squint test, or hunt code smells and fix them with named refactorings.
---

# Squint review

Work like a careful human with all afternoon: find the smells, name them, and remove them one named refactoring at a time, with the tests green after every step. Refactoring changes structure only. Behavior stays exactly as it is.

## The five questions

Ask these of every class in scope, while squinting and again at the end:

1. Does this code duplicate any behavior?
2. Does everything have just one responsibility?
3. Does this code depend on something that changes more often than it does?
4. Does everything in this class change at the same rate?
5. Is this class open for extension and closed for modification?

## Steps

### 1. Set up the worktree

- PR number or URL: `gh pr view <pr> --json number,headRefName,baseRefName,isCrossRepository`, fetch the head branch, and `git worktree add ../<repo>-squint-<branch> <headRefName>`.
- Branch: the same, with that branch against the default branch.
- Nothing given: work in place on the current branch.

If git refuses because the branch is checked out elsewhere, stop and tell the user. The **scope** is the production files the branch changes relative to its base. Tests are touched only to keep them compiling or to add characterization tests.

Done when you are in the worktree and have the list of files in scope.

### 2. Go green

Find the project's test command from its manifests, scripts, or `CLAUDE.md`, and run the tests covering the scope. A red suite ends the review: report the failures and stop.

For each file in scope, judge whether its tests pin the behavior you are likely to move. Where they do not, plan **characterization tests** that record what the code does now, right or wrong.

Done when the suite is green and every file in scope is either covered or has characterization tests planned.

### 3. Squint

Read each file in scope as a shape, not as text.

- **Changes in shape**: indentation drifting right, nesting, long methods. Shape is where the conditionals live.
- **Changes in color**: runs of different syntax clustered together (literals beside calls beside operators). Color marks mixed levels of abstraction.

Then read closely where the squint pointed, and ask the five questions of each class. Record every smell by its standard name (Fowler's or Kerievsky's), with `path:line` and what it costs in concrete terms: the next change it makes harder.

Done when every file in scope has been squinted at and every class has answers to the five questions.

### 4. Plan

Turn the smells into an ordered list of refactorings. Choose each one from [smells-to-patterns.md](smells-to-patterns.md).

- Characterization tests come first, then the small cleanups that make a pattern visible (Compose Method, Extract Method), then the pattern refactorings.
- Say for each one whether you refactor **to** the pattern, **toward** it, or **away** from it. Stop toward a pattern once the smell's cost is gone. Go away from one that the code does not earn (Speculative Generality).
- List the smells you are deliberately deferring, each with the reason.

Show the user the plan: smell, location, cost, refactoring, direction. Wait for them to approve or trim it.

Done when the user approves a plan.

### 5. Refactor

For each approved refactoring, in order:

1. Announce it: the smell, the location, the refactoring, and the direction.
2. Follow the book's mechanics as a series of small steps. Run the tests after each step.
3. Green: take the next step. Red: undo that step and take a smaller one. If the refactoring cannot stay green, reset to the last commit, mark it abandoned with the reason, and move on.
4. Commit once the refactoring is complete and green. One refactoring per commit, in this shape:

   ```
   <Refactoring name> in <class or method>

   <What changed structurally, in a sentence or two.>

   Smells addressed:
   - <Smell>: <what it cost>. Cured with <refactoring>.
   ```

Keep new smells found along the way for the report, and leave them unfixed: the approved plan is the whole job.

Done when every approved refactoring is committed or abandoned with a reason.

### 6. Report

- Each refactoring with its smell, direction, and commit SHA, plus any abandoned ones and why.
- The five questions answered again for each class you touched, naming any compromise to SOLID that remains.
- Deferred smells and the new ones found while refactoring.
- The worktree path, the command to push from it, and `git worktree remove <path>` for when they are done. Pushing is the user's call.
