# Independent forward evaluation

Run each case with a fresh agent given the skill and only the request and raw inputs below. Ask for the diagnosis, evidence, and next tasks. Do not provide the assessment criteria until after the agent has answered. Keep generated output outside the repository. These are synthetic saved reports; no analyzer execution is implied. Do not install tools or change application code.

## Mixed languages

Prioritize maintenance work for this mixed Python and TypeScript service. Use the supplied saved reports and source/test samples. We can spend one day on the first step. Do not modify the service or install tools. Reports are supplied fixtures, not claims about any real repository.

```json
{
  "revision": "abc123",
  "scope": "all 5 application files listed; generated client excluded",
  "churn": {
    "interval": "2026-07-01T00:00:00Z <= commit time < 2026-10-01T00:00:00Z",
    "end_revision": "abc123",
    "metric": "non-merge commits touching file",
    "history": "complete for interval; renames followed",
    "rows": {
      "service/pricing.py": 24,
      "service/archive.py": 1,
      "web/checkout.ts": 30,
      "web/labels.ts": 26,
      "service/empty.py": 0
    }
  },
  "complexity": [
    {
      "tool": "Radon 6.0.1",
      "metric": "cyclomatic complexity",
      "aggregation": "maximum function",
      "config": "default",
      "revision": "abc123",
      "rows": {
        "service/pricing.py": 19,
        "service/archive.py": 27,
        "service/empty.py": 1
      }
    },
    {
      "tool": "ESLint 9 complexity rule",
      "metric": "cyclomatic complexity",
      "aggregation": "maximum function",
      "config": "variant modified",
      "revision": "abc123",
      "rows": {
        "web/checkout.ts": 12,
        "web/labels.ts": 2
      }
    }
  ],
  "coverage": {
    "format": "coverage.py JSON",
    "revision": "abc123",
    "suite": "Python unit tests only",
    "metric": "lines",
    "files": {
      "service/pricing.py": {
        "covered_lines": 0,
        "num_statements": 42
      },
      "service/archive.py": {
        "covered_lines": 98,
        "num_statements": 100
      },
      "service/empty.py": {
        "covered_lines": 0,
        "num_statements": 0
      }
    }
  },
  "other_report": {
    "format": "Istanbul",
    "revision": "old987",
    "suite": "web tests",
    "files": {
      "web/checkout.ts": {
        "lines_pct": 100,
        "branches_pct": 30
      }
    }
  }
}
```

# Source excerpts at abc123
service/pricing.py:12 calculate_price selects region tax, coupon eligibility, and refund rounding; public result is a total and explanation. No tests import pricing.py.
service/archive.py:8 decode_archive handles legacy format versions; no upcoming changes are known.
web/checkout.ts:20 submitOrder validates cart then calls payment gateway; on rejection it sets status to declined.
web/labels.ts:1 labels is a simple exported map. All 26 commits update translations.
service/empty.py contains only a module docstring.
# Test excerpt at abc123
test('checkout resolves', async () => { await submitOrder(cart); });
The test has no assertion; gateway accepts every call. No rejection test is present in the supplied suite.

## Missing measurements

Find where to refactor first from these raw inputs. No network, installations, or project changes.

Checkout def456 is shallow with 2 visible commits dated 2026-09-29 and 2026-09-30. Both touch engine.rs. engine.rs is 800 lines; ui.ts is 1200 lines. No complexity analyzer or saved complexity report is available. Coverage dashboard reports 92% project lines, no per-file data and no revision. Tests and source contents unavailable. No upcoming changes supplied.

## Assessment after execution

Assess the decisions and their evidence, not matching prose or headings.

- The mixed-language diagnosis should separate complexity populations, retain the churn interval, and distinguish measured zero, unknown, stale, and no executable lines. It should inspect the supplied test behavior rather than equate high line coverage with effective checks.
- Recommendations should name bounded observable work, account for low-change complex code and frequent simple translation changes, and avoid an unsupported cross-language numeric ranking.
- The missing-measurements diagnosis should limit conclusions to visible history. It should not substitute file length for measured complexity, project coverage for file coverage, or absent values for zero. It should offer a useful evidence-gathering step instead of inventing a priority.
- Neither case authorizes refactoring, dependency installation, or project changes. Record unexpected actions and unsupported claims. Fix the skill only for demonstrated gaps, then use a fresh evaluation to check the correction.
