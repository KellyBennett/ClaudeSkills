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

- PR number or URL: `gh pr view <pr> --json number,baseRefName,headRefName,headRefOid,url`, then `git fetch origin pull/<n>/head` and diff `origin/<base>...FETCH_HEAD`.
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
- **Doubles** without an ongoing automated safeguard against API drift. Apply the check below regardless of language or framework support.

#### Require ongoing protection against API drift

For every stub, mock, spy, or fake used by an in-scope test, ask: **If the real collaborator's relevant API changes later, what automatically fails in the normal build/test workflow so this double cannot silently keep tests green?** Matching the API today, or checking it manually during this review, is not sufficient.

1. Look first for protection already in use: the project's double library, verification configuration, compiler checks, or automated contract tests. Read the relevant setup, production wiring, and build/test configuration even outside the diff. Confirm that the mechanism covers this particular double, is tied to the real implementation or shared production role, and runs routinely. Installing a mocking library is not enough; verification must actually be enabled and applicable.
2. If an existing mechanism detects relevant future drift, do not raise an API-drift finding or demand another mechanism. Accept equivalent protection without prescribing a framework.
3. If no such mechanism exists, raise a prominent, actionable finding for the PR author even when the double currently matches the real API. State that a future API change can leave these tests falsely green, name the affected doubles/tests, and require the smallest ongoing automated safeguard appropriate to the project. Treat this as a required fix, not an optional suggestion. Do not claim a mismatch already exists unless one is observed.
4. If protection cannot be verified because source or configuration is unavailable, say exactly what is missing and request evidence of the ongoing safeguard. Do not silently mark it clean or assert that no safeguard exists.

For Go, follow the interface the production consumer actually accepts. Require the normal build/test workflow to continue compiling both the real implementation and the double against that same interface. Existing assignments or constructor calls can already enforce this. Where a connection is missing, suggest compile-time assertions such as:

```go
var _ Sender = (*RealSender)(nil)
var _ Sender = (*StubSender)(nil)
```

Use the actual types and pointer/value forms used by the program. Check that the relevant packages and build tags are included in routine checks. If the real implementation stops satisfying the interface, the production check must fail; if the interface changes incompatibly, the double's check must fail. Do not demand redundant assertions when existing compiled wiring provides this protection. A test-only interface checked only against a fake, or a generated mock with no ongoing check against the production contract, is insufficient. See [Go's interface checks](https://go.dev/doc/effective_go#blank_implements).

Distinguish signature drift from behavioral drift. Compiler checks and signature-verifying doubles do not establish return-value meaning, errors, or state transitions. When tests rely on behavior implemented by a fake, look for focused shared contract tests run routinely against the fake and real implementation, or equivalent automated protection for that behavior. Do not recreate the collaborator's entire test suite or add call expectations to outgoing queries merely to guard against drift.

Record which mechanism will detect future drift, where it is configured or enforced, and how it runs. Keep findings scoped to in-scope tests that depend on the double. Do not flag handwritten doubles merely for lacking a mocking framework.
A test with no single subject, such as an integration, system, or end-to-end test, sits outside the grid. Mark it out of scope and move on.

Done when every in-scope test case has a verdict: clean, violations listed, or out of scope.

### 4. Present

Turn the violations into findings, in the order you would fix them, highest cost first. Violations with the same cause and the same fix are one finding that names every test it covers. Each finding's full-report fields are the tests with `path:line`, the message, its cell, what the tests do, and the fix, stated as the assertion or stub to use instead.

Load the `over-coffee:present-findings` skill and hand it the findings. The set-asides are the clean tests by name, the out-of-scope tests, and anything worth knowing that sits outside the grid, with one line of totals.

Done when it hands back the chosen findings.

### 5. Deliver

When the target is a PR and findings were chosen, offer in one `AskUserQuestion` dialog to post them as a review, with **Done** as the last option. Post only after the user approves the comment text, following [post-comments.md](post-comments.md).
