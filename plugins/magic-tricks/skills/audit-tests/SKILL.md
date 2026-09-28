---
name: audit-tests
description: Audit the tests added or modified in a PR or branch against Sandi Metz's Magic Tricks of Testing grid (incoming, sent-to-self, outgoing × query, command). Use when the user asks to audit, review, or check the tests in a PR or branch, or asks whether tests over-mock, test private methods, or assert the wrong thing.
---

# Audit tests against the Magic Tricks grid

A unit test watches one **subject** from the outside. Every message the test touches has an **origin** relative to that subject and a **kind**, and the pair decides what the test may do with it.

| Origin       | Query                | Command                     |
|--------------|----------------------|-----------------------------|
| Incoming     | Assert return value  | Assert direct public side effect |
| Sent to self | Ignore               | Ignore                      |
| Outgoing     | Ignore (stub freely) | Expect it to be sent        |

- **Incoming**: sent to the subject by the test, through its public interface.
- **Sent to self**: the subject calling its own methods, private or not.
- **Outgoing**: the subject sending to a collaborator.
- **Query**: returns something, changes nothing. **Command**: changes something. A message that does both is a command, and as incoming also has a return value worth asserting.

## Steps

### 1. Resolve the target

- PR number or URL: `gh pr view <pr> --json number,baseRefName,headRefName,url`, then `git fetch origin pull/<n>/head` and diff `origin/<base>...FETCH_HEAD`.
- Branch: diff `<default-branch>...<branch>`.
- Nothing given: the current branch against the default branch.

Read files at the head revision with `git show <rev>:<path>`. Leave the working tree and checked-out branch as they are.

Done when you have the list of changed test files and the base and head revisions.

### 2. Collect the test cases in scope

A test case is in scope when the diff adds it or changes any line inside it, including shared setup (`before`, `setUp`, fixtures, `let`) it depends on. Read each test file in full, and read the subject's source, so you can tell public from private and collaborator from self.

Done when every in-scope test case is on a list with its file and line.

### 3. Place every message on the grid

For each test case, name the subject, then every message the test sends, stubs, spies on, or asserts about. Give each an origin and a kind, and compare what the test does with it to its cell. The violations to hunt:

- **Incoming query** asserted by anything other than its return value, such as expecting internal calls or inspecting internal state.
- **Incoming command** asserted through private state (reflection, `instance_variable_get`, reaching into internals) or through the calls it makes, when a public side effect is observable.
- **Sent to self** tested at all: private methods called directly, the subject stubbed or expected to receive its own messages (partial mocks).
- **Outgoing query** asserted as sent (`expect(x).to receive(:query)`, `verify(x).query()`). Stubbing it is correct.
- **Outgoing command** left unasserted, or asserted by checking its effect inside the collaborator (a database row, a delivered email) instead of expecting the message.
- **Doubles** that can drift from the collaborator's real interface, when the framework offers verified doubles (`instance_double`, `create_autospec`, typed mocks).

A test with no single subject, such as an integration, system, or end-to-end test, sits outside the grid. Mark it out of scope and move on.

Done when every in-scope test case has a verdict: clean, violations listed, or out of scope.

### 4. Report

Group by file. For each test case give `path:line`, the test name, and the verdict. For each violation give the message, its cell, what the test does, and the fix, stated as the assertion or stub it should use instead. End with one line of totals.

List clean tests by name only, so the report shows every in-scope test was examined.

### 5. Offer PR comments

When the target is a PR and there are violations, offer to post them as inline review comments. Post only after the user approves, following [post-comments.md](post-comments.md).
