---
name: present-findings
description: Present a code review's findings, all at once or one at a time over coffee, and collect the user's decision on each. Use when a review skill has its findings ready to show.
---

# Present findings

A review skill hands you its **findings**, in the order it would fix them. Each has one or more locations (`path:start-end`, plus the PR URL and head commit when there is one), what is wrong, what it costs, and the change the reviewer would ask for, along with any fields the review skill lists for its full report. It may also hand you **set-asides** (things it looked at and would leave alone) and one line of CI context.

Every choice the user makes here is an `AskUserQuestion` dialog, so they pick with Enter. Its built-in "Other" is where they type when they want to talk one through.

The findings are your working notes until the user picks how to see them. Your first message is that dialog alone: "I've read the code. Do you want everything at once, or one at a time, starting with what I'd fix first?" with the options **Everything at once** and **One at a time**.

## Everything at once

Each finding with the fields the review skill lists for its report. Then the set-asides, and the CI line. Every finding counts as in the review; the user trims when approving what gets posted.

## One at a time

Load the `over-coffee:coffee` skill. Each finding opens with a header, stepped out of the conversation because a location has to be exact: each location with, for a PR, a link beside it. The user looks at the code there while you talk. Written as a bullet per location, like these:

- api/internal/events/reads.go:46-51 · [See in diff](<pr-url>/files#diff-<anchor>R46-R51)
- api/internal/events/handler.go:169-176 · [See in file](<repo-url>/blob/<head-commit>/api/internal/events/handler.go#L169-L176)

**See in diff** is for lines the PR changed, and **See in file** for lines it did not. Compute each diff anchor with `printf '%s' '<path>' | shasum -a 256 | cut -d' ' -f1` and paste it straight in: a link is one unbroken URL.

Then talk it over at the table: what is wrong, what it costs, the change you would ask the author for, and what it depends on. You are a reviewer recommending a change to someone else's code, and you sound like one: "I'd ask them to pull that into one function."

Each turn hands the conversation back the coffee way. When the user's view lands, or they say to move on, confirm it with a dialog: **Talk it over** first, then **Put it in the review**, **Mention it as optional**, **Leave it out**. When the review skill says the user is fixing rather than reviewing, the decisions are **Fix it** and **Leave it**. Talk it over picks the conversation back up.

After the last finding, give CI one line, then ask one dialog: "There are <n> things I'd leave alone. Want to go through them?" with **Skip them** first and **Go through them** second. Going through them is the same walk, one set-aside per turn with its header, shorter, each ending on a dialog: **Leave it alone**, then **Put it in the review** (**Fix it** when fixing).

## Hand back

The **chosen** findings are every one the user did not leave out, each marked as in the review, optional, or to fix. Hand that list back to the review skill.

Done when the user has seen every finding and made a choice for each.
