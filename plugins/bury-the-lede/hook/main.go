// Command bury-the-lede is a PreToolUse hook for Bash. It denies a
// `gh pr create|edit` whose body skips the bury-the-lede skill's shape, and
// stays silent for everything else. It always exits 0.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// wholeCallDenied is said on every denial: a PreToolUse deny drops the whole
// Bash call, so a step chained before the gh command didn't run either.
const wholeCallDenied = "Nothing in this Bash call ran, including any steps before or after this command; rerun them once it is fixed."

func main() {
	var event struct {
		CWD       string `json:"cwd"`
		ToolInput struct {
			Command string `json:"command"`
		} `json:"tool_input"`
	}
	raw, err := io.ReadAll(os.Stdin)
	if err != nil || json.Unmarshal(raw, &event) != nil {
		return
	}
	if reason := denial(event.ToolInput.Command, event.CWD, repoRoot, ghCLI{}); reason != "" {
		printDenial(reason)
	}
}

// denial returns why cmd may not run, or "" when it may.
func denial(cmd, cwd string, rootOf func(string) string, gh gitHub) string {
	if !strings.Contains(cmd, "gh") || !strings.Contains(cmd, "pr") {
		return ""
	}
	pr, ok := parsePRInvocation(cmd)
	if !ok || !pr.setsBody {
		return ""
	}
	dir := commandDir(cmd, cwd, pr.at)
	files, readable := visibleBody(cmd, dir)
	if !readable {
		return ""
	}
	o, ok := outlineFor(rootOf(dir), pr.repo, gh)
	if !ok {
		return ""
	}
	r := review{body: cmd + "\n" + files, outline: o}
	if pr.edit {
		if before, ok := gh.prBody(pr.selector, pr.repo, dir); ok {
			r = r.asEditOf(before)
		}
	}
	return r.problems()
}

func repoRoot(dir string) string {
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func printDenial(reason string) {
	out, _ := json.Marshal(map[string]any{
		"hookSpecificOutput": map[string]string{
			"hookEventName":            "PreToolUse",
			"permissionDecision":       "deny",
			"permissionDecisionReason": reason + "\n\n" + wholeCallDenied,
		},
	})
	fmt.Println(string(out))
}
