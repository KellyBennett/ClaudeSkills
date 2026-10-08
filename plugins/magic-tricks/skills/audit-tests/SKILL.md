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
- **Doubles** with no executable connection to the collaborator's real API. Apply the contract check below regardless of language or framework support.

#### Keep doubles synchronized with the real API

For every stub, mock, spy, or fake used by an in-scope test, identify the production collaborator or role it replaces and the executable evidence that both conform to that role. Read the relevant interface, production wiring, and contract checks even when those files are outside the diff. Keep findings scoped to the in-scope tests that depend on them.

- Prefer an existing compiler check or verified double tied to the real collaborator (`instance_double`, `create_autospec`) when it checks the relevant methods and signatures. Accept equivalent protection; do not require a particular library or redundant assertions.
- Without that protection, require an automated contract check against the real implementation. A separately copied test interface, matching method names by inspection, or a promise to update the double manually is not synchronization. If the implementation is unavailable, report the contract as unverified and name the missing evidence; do not invent a mismatch or mark it clean.
- Distinguish API shape from behavior. Signature checks do not establish return-value meaning, errors, or state transitions. When tests rely on behavior implemented by a fake, require focused shared contract tests against the fake and real implementation for that relied-on behavior, or equivalent existing evidence. Do not recreate the collaborator's entire test suite or add call expectations to outgoing queries merely to verify a double.

For Go, follow the interface the production consumer actually accepts. Confirm that both the real implementation and the double are checked against that same interface in compiled production/test code. Ordinary assignments or constructor calls can already provide this guarantee. Where the connection is otherwise absent, use compile-time assertions such as:

```go
var _ Sender = (*RealSender)(nil)
var _ Sender = (*StubSender)(nil)
```

Use the actual types and pointer/value forms used by the program. Do not demand these declarations when existing wiring already proves conformance. Generated mocks receive the same check; generation alone is not proof that the current production implementation still satisfies the role. A test-only interface checked only against a fake leaves the real API unchecked. Use the repository's build/test checks to compile both sides, including relevant build tags. See [Go's interface checks](https://go.dev/doc/effective_go#blank_implements).

For each double, record the real role, the evidence location, what the check covers (signatures and/or behavior), and any remaining gap. In a finding, name the missing connection and the smallest automated check that would detect drift. Do not classify a stub or mock as unsafe solely because it is handwritten or does not use a mocking framework.

A test with no single subject, such as an integration, system, or end-to-end test, sits outside the grid. Mark it out of scope and move on.

Done when every in-scope test case has a verdict: clean, violations listed, or out of scope.

### 4. Present

Turn the violations into findings, in the order you would fix them, highest cost first. Violations with the same cause and the same fix are one finding that names every test it covers. Each finding's full-report fields are the tests with `path:line`, the message, its cell, what the tests do, and the fix, stated as the assertion or stub to use instead.

Load the `over-coffee:present-findings` skill and hand it the findings. The set-asides are the clean tests by name, the out-of-scope tests, and anything worth knowing that sits outside the grid, with one line of totals.

Done when it hands back the chosen findings.

### 5. Deliver

When the target is a PR and findings were chosen, offer in one `AskUserQuestion` dialog to post them as a review, with **Done** as the last option. Post only after the user approves the comment text, following [post-comments.md](post-comments.md).
