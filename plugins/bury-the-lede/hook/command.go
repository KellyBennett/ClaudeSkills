package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	commandStart      = `(?:^|&&|\|\||[;|\n(])\s*(?:\w+=\S*\s+)*`
	bodyFileReadLimit = 256 << 10
)

var (
	prCommand    = regexp.MustCompile(commandStart + `(gh\s+pr\s+(create|edit)\b)([^;&|\n]*)`)
	heredocStart = regexp.MustCompile(`^<<-?\s*['"]?(\w+)['"]?`)
	cdSegment    = regexp.MustCompile(`(?:^|&&|\|\||;|\n)\s*cd\s+("[^"]*"|'[^']*'|[^\s;&|]+)`)
	bodyFlag     = regexp.MustCompile(`(^|\s)(--body|--body-file|-b|-F|--fill|--fill-first|--fill-verbose|-f)([\s=]|$)`)
	repoFlag     = regexp.MustCompile(`(?:^|\s)(?:-R|--repo)[\s=]+(\S+)`)
	bodyFileFlag = regexp.MustCompile(`(?:^|\s)(?:--body-file|--file|-F)[\s=]+("[^"]*"|'[^']*'|\S+)`)
	expandedBody = regexp.MustCompile(`(?:^|\s)(?:--body|-b)[\s=]+"?\$(\(cat\s+<<)?`)
)

// prInvocation is the `gh pr create|edit` a command runs. selector and repo
// are empty when the command doesn't name them.
type prInvocation struct {
	at       int
	edit     bool
	setsBody bool
	selector string
	repo     string
}

func parsePRInvocation(cmd string) (prInvocation, bool) {
	code := shellCode(cmd)
	m := prCommand.FindStringSubmatchIndex(code)
	if m == nil {
		return prInvocation{}, false
	}
	args := code[m[6]:m[7]]
	pr := prInvocation{at: m[2], edit: code[m[4]:m[5]] == "edit", setsBody: bodyFlag.MatchString(args)}
	if fields := strings.Fields(args); pr.edit && len(fields) > 0 && !strings.HasPrefix(fields[0], "-") {
		pr.selector = fields[0]
	}
	if r := repoFlag.FindStringSubmatch(args); r != nil {
		pr.repo = r[1]
	}
	return pr, true
}

// commandDir follows every `cd` before offset end. The hook's cwd is the
// session's, so `cd other-repo && gh pr edit` would otherwise be judged by the
// session repo's PR template.
func commandDir(cmd, cwd string, end int) string {
	dir := cwd
	for _, m := range cdSegment.FindAllStringSubmatch(cmd[:end], -1) {
		target := strings.Trim(m[1], `"'`)
		if target == "~" || strings.HasPrefix(target, "~/") {
			home, err := os.UserHomeDir()
			if err != nil {
				continue
			}
			target = filepath.Join(home, strings.TrimPrefix(target, "~"))
		}
		if !filepath.IsAbs(target) {
			target = filepath.Join(dir, target)
		}
		dir = filepath.Clean(target)
	}
	return dir
}

// shellCode blanks quoted strings and heredoc bodies, keeping every byte
// offset, so a command that is only mentioned doesn't count as run. Command
// substitution inside double quotes is blanked with them, so
// `x="$(gh pr create …)"` goes unchecked: that is the price of not parsing
// shell.
func shellCode(cmd string) string {
	b := []byte(cmd)
	var quote byte
	var heredocs []string
	for i := 0; i < len(b); i++ {
		c := b[i]
		switch {
		case quote == '"' && c == '\\' && i+1 < len(b):
			b[i], b[i+1] = ' ', ' '
			i++
		case quote != 0 && c == quote:
			quote = 0
		case quote != 0:
			b[i] = ' '
		case c == '\\':
			i++
		case c == '\'' || c == '"':
			quote = c
		case strings.HasPrefix(cmd[i:], "<<<"):
			i += 2
		case c == '<':
			if m := heredocStart.FindStringSubmatch(cmd[i:]); m != nil {
				heredocs = append(heredocs, m[1])
				i += len(m[0]) - 1
			}
		case c == '\n' && len(heredocs) > 0:
			i = blankHeredocs(b, i+1, heredocs) - 1
			heredocs = nil
		}
	}
	return string(b)
}

func blankHeredocs(b []byte, from int, delimiters []string) int {
	at := from
	for _, delim := range delimiters {
		for at < len(b) {
			start, eol := at, len(b)
			if n := bytes.IndexByte(b[at:], '\n'); n >= 0 {
				eol = at + n
			}
			at = eol + 1
			if strings.TrimSpace(string(b[start:eol])) == delim {
				break
			}
			for j := start; j < eol; j++ {
				b[j] = ' '
			}
		}
	}
	return min(at, len(b))
}

// visibleBody returns the body files' contents, and false when the hook can't
// see the whole body: a body file it can't read (its path may hold a shell
// variable) or a --body the shell expands. An unseen body is never judged. A
// $(cat <<EOF …) heredoc is visible, since its text is in cmd.
func visibleBody(cmd, dir string) (string, bool) {
	var b strings.Builder
	readable := true
	for _, m := range bodyFileFlag.FindAllStringSubmatch(cmd, -1) {
		path := strings.Trim(m[1], `"'`)
		if path == "-" {
			continue
		}
		if !filepath.IsAbs(path) {
			path = filepath.Join(dir, path)
		}
		f, err := os.Open(path)
		if err != nil {
			readable = false
			continue
		}
		data, _ := io.ReadAll(io.LimitReader(f, bodyFileReadLimit))
		f.Close()
		b.Write(data)
		b.WriteByte('\n')
	}
	for _, m := range expandedBody.FindAllStringSubmatch(cmd, -1) {
		if m[1] == "" {
			readable = false
		}
	}
	return b.String(), readable
}
