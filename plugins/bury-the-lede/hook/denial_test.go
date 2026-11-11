package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fakeGitHub struct {
	body      string
	bodyFound bool
	files     map[string]string
	fileErr   error
}

func (f fakeGitHub) prBody(_, _, _ string) (string, bool) { return f.body, f.bodyFound }

func (f fakeGitHub) file(repo, path string) (string, error) {
	if f.fileErr != nil {
		return "", f.fileErr
	}
	if content, ok := f.files[repo+"/"+path]; ok {
		return content, nil
	}
	return "", os.ErrNotExist
}

func noRepo(string) string { return "" }

const framedBody = `## What

Students land on their dashboard after signing in.

## Why

Sign-in ended on a blank page.

## How

The return path now survives the round trip. ` + "`auth/callback.go`" + ` reads it from the cookie.`

func create(body string) string {
	return "gh pr create --title t --body \"$(cat <<'EOF'\n" + body + "\nEOF\n)\""
}

func TestDenial(t *testing.T) {
	cases := []struct {
		name    string
		cmd     string
		gh      fakeGitHub
		wantAll []string
	}{
		{name: "unrelated command", cmd: "ls -la"},
		{name: "framed body", cmd: create(framedBody)},
		{name: "create without a body", cmd: "gh pr create --web"},
		{
			name:    "missing section",
			cmd:     create(strings.Replace(framedBody, "## Why", "## Reason", 1)),
			wantAll: []string{"missing Why", "bury-the-lede skill"},
		},
		{
			name:    "section opens on code",
			cmd:     create(strings.Replace(framedBody, "The return path now", "`auth/callback.go` now", 1)),
			wantAll: []string{"How opens with", "framing"},
		},
		{
			name:    "section opens on a path",
			cmd:     create(strings.Replace(framedBody, "Students land", "auth/callback.go lands", 1)),
			wantAll: []string{"What opens with"},
		},
		{name: "prose word that looks like code", cmd: create(strings.Replace(framedBody, "Students land", "macOS users land", 1))},
		{name: "leading ticket key then prose", cmd: create(strings.Replace(framedBody, "Sign-in ended", "QTY-1234: Sign-in ended", 1))},
		{name: "N/A section", cmd: create(strings.Replace(framedBody, "Sign-in ended on a blank page.", "N/A", 1))},
		{name: "only mentioned in a string", cmd: `echo "gh pr create --body '## What'"`},
		{name: "only mentioned in a heredoc", cmd: "cat <<EOF\ngh pr create --body x\nEOF"},
		{name: "body from a variable", cmd: `gh pr create --body "$BODY"`},
		{name: "fill replaces the body", cmd: "gh pr create --fill", wantAll: []string{"missing What, Why, How"}},
		{
			name: "remote template wins",
			cmd:  "gh pr create -R acme/app " + create(framedBody)[len("gh pr create "):],
			gh: fakeGitHub{files: map[string]string{
				"acme/app/.github/pull_request_template.md": "## Summary\n\n## Testing\n\n## Screenshots\n<!-- delete this section if none -->\n",
			}},
			wantAll: []string{"acme/app's PR template", "missing Summary, Testing", "replaces the template"},
		},
		{name: "remote template unreadable", cmd: "gh pr create -R acme/app --body x", gh: fakeGitHub{fileErr: os.ErrPermission}},
		{
			name: "edit keeps what the author removed",
			cmd:  "gh pr edit 7 --body \"## What\n\nA plain sentence.\n\n## Why\n\nAnother one.\"",
			gh:   fakeGitHub{body: "## What\n\nOld.\n\n## Why\n\nOld.", bodyFound: true},
		},
		{
			name:    "edit drops a section that was there",
			cmd:     "gh pr edit 7 --body \"## What\n\nA plain sentence.\"",
			gh:      fakeGitHub{body: "## What\n\nOld.\n\n## Why\n\nOld.", bodyFound: true},
			wantAll: []string{"This edit removes Why"},
		},
		{
			name: "edit leaves an unframed line untouched",
			cmd:  "gh pr edit 7 --body \"## What\n\n`x.go` changes.\n\n## Why\n\nNew reason.\n\n## How\n\nNew approach.\"",
			gh:   fakeGitHub{body: "## What\n\n`x.go` changes.\n\n## Why\n\nOld.\n\n## How\n\nOld.", bodyFound: true},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := denial(c.cmd, t.TempDir(), noRepo, c.gh)
			if len(c.wantAll) == 0 {
				if got != "" {
					t.Fatalf("want allowed, got denial:\n%s", got)
				}
				return
			}
			for _, want := range c.wantAll {
				if !strings.Contains(got, want) || got == "" {
					t.Fatalf("want denial containing %q, got:\n%s", want, got)
				}
			}
		})
	}
}

func TestDenialReadsBodyFileAfterCd(t *testing.T) {
	home := t.TempDir()
	other := filepath.Join(home, "other")
	if err := os.Mkdir(other, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(other, "body.md"), []byte("## What\n\n`main.go` changed."), 0o644); err != nil {
		t.Fatal(err)
	}
	got := denial("cd other && gh pr create --body-file body.md", home, noRepo, fakeGitHub{})
	if !strings.Contains(got, "missing Why, How") || !strings.Contains(got, "What opens with") {
		t.Fatalf("want missing sections and framing problems, got:\n%s", got)
	}
}

func TestDenialUsesLocalTemplateOfTheCdTarget(t *testing.T) {
	session, other := t.TempDir(), t.TempDir()
	tmpl := filepath.Join(other, ".github", "pull_request_template.md")
	if err := os.MkdirAll(filepath.Dir(tmpl), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tmpl, []byte("## Summary\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	rootOf := func(dir string) string {
		if dir == other {
			return other
		}
		return ""
	}
	got := denial("cd "+other+" && "+create(framedBody), session, rootOf, fakeGitHub{})
	if !strings.Contains(got, "this repo's PR template") || !strings.Contains(got, "missing Summary") {
		t.Fatalf("want the cd target's template applied, got:\n%s", got)
	}
}

func TestShellCodeKeepsOffsets(t *testing.T) {
	cmd := "echo 'gh pr create' && cat <<EOF\nbody\nEOF\ngh pr edit"
	code := shellCode(cmd)
	if len(code) != len(cmd) {
		t.Fatalf("length changed: %d -> %d", len(cmd), len(code))
	}
	if strings.Count(code, "gh pr") != 1 || !strings.HasSuffix(code, "gh pr edit") {
		t.Fatalf("want only the real invocation left, got %q", code)
	}
}
