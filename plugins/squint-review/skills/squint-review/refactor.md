# Refactor mode

Carry out the chosen recommendations one named refactoring at a time, with the tests green after every step. Refactoring changes structure only. Behavior stays exactly as it is.

## 1. Check out the branch

- PR: `gh pr checkout <pr>`.
- Branch: `git switch <branch>`.
- Nothing given: stay on the current branch.

A dirty working tree blocks the checkout. Say so in one line and stop.

Done when `HEAD` is the target branch's head.

## 2. Go green

Find the project's test command from its manifests, scripts, or `CLAUDE.md`.

When the target is a PR, `HEAD` is its `headRefOid`, and every check in `gh pr checks` passed, CI has already shown the suite green: run only the tests covering the first refactoring, to prove this machine can run them. Otherwise run the tests covering the scope. Red ends refactor mode: report the failures and stop.

For each file the plan touches, judge whether its tests pin the behavior you are about to move. Where they do not, add **characterization tests** that record what the code does now, right or wrong, as the first commit.

Done when the tests are green and every file the plan touches is covered.

## 3. Refactor

For each chosen recommendation, in order:

1. Announce it: the smell, the location, the refactoring, and the direction.
2. Follow the book's mechanics as a series of small steps. Run the tests after each step.
3. Green: take the next step. Red: undo that step and take a smaller one. If the refactoring cannot stay green, reset to the last commit, mark it abandoned with the reason, and move on. Anything that depends on it is abandoned with it.
4. Commit once the refactoring is complete and green. One refactoring per commit, in this shape:

   ```
   <Refactoring name> in <unit>

   <What changed structurally, in a sentence or two.>

   Smells addressed:
   - <Smell>: <what it cost>. Cured with <refactoring>.
   ```

Keep new smells found along the way for the report, and leave them unfixed: the plan is the whole job.

Done when every chosen recommendation is committed or abandoned with a reason.

## 4. Report

- Each refactoring with its smell, direction, and commit SHA, plus any abandoned ones and why.
- The five questions answered again for each unit you touched, naming any compromise to SOLID that remains.
- New smells found while refactoring.

Pushing is the user's call.
