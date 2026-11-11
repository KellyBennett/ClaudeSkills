package main

import (
	"fmt"
	"regexp"
	"strings"
)

// The parts of codeToken. Each is narrow enough that prose words such as
// macOS, iPad, Node.js, CI/CD and A/B are not mistaken for code.
const (
	pathToken      = `[a-z0-9_.~-]/|/[a-z0-9_.]`
	fileNameToken  = `^(?:[a-z_][\w-]*|[A-Z][A-Z0-9_-]*)\.(?:go|rb|ts|tsx|js|jsx|py|less|css|md|ya?ml|json|sql|sh)\b`
	camelIdent     = `^[a-z][a-z0-9]+[A-Z][a-z0-9]`
	qualifiedIdent = `^[A-Z]\w*\.(?:[A-Z][a-z]|[a-z]{3})`
)

var (
	codeOpening   = regexp.MustCompile("^(`|[-*+|>] |\\d+[.)] )")
	codeToken     = regexp.MustCompile(strings.Join([]string{`::`, `\(\)`, pathToken, fileNameToken, camelIdent, qualifiedIdent}, "|"))
	notApplicable = regexp.MustCompile(`(?i)^n/a\b`)
	ticketRef     = regexp.MustCompile(`^\[?[A-Z][A-Z0-9]+-\d+\]?(\(\S+?\))?[:.,]?(\s+|$)`)
)

// review is one PR body held against its outline. When editing, before is the
// PR's current body: the edit may keep what the author already chose and is
// judged only on what it changes.
type review struct {
	body    string
	outline outline
	editing bool
	before  string
}

func (r review) asEditOf(before string) review {
	r.outline = r.outline.keptBy(before)
	r.editing = true
	r.before = before
	return r
}

type rule func(review) string

var rules = []rule{missingSections, unframedSections}

func (r review) problems() string {
	var found []string
	for _, check := range rules {
		if p := check(r); p != "" {
			found = append(found, p)
		}
	}
	return strings.Join(found, "\n\n")
}

func missingSections(r review) string {
	present := headingsIn(r.body)
	var missing []string
	for _, s := range r.outline.sections {
		if k := strings.ToLower(s.name); s.required && !present[k] {
			present[k] = true
			missing = append(missing, s.name)
		}
	}
	if len(missing) == 0 {
		return ""
	}
	list := strings.Join(missing, ", ")
	if r.editing {
		return fmt.Sprintf("This edit removes %s, which the PR's current description has and %s requires. Keep those sections, or write N/A under one that no longer applies.", list, r.outline.source)
	}
	advice := ""
	if r.outline.template {
		advice = " Passing --body, --body-file or --fill replaces the template, so start from its text; write N/A under a section that does not apply."
	}
	return fmt.Sprintf("This PR description does not follow %s: it is missing %s. Rebuild the body with the bury-the-lede skill.%s", r.outline.source, list, advice)
}

// unframedSections skips checklists, ticket keys and optional sections such as
// screenshots: they are lists, identifiers or images by design. An edit that
// leaves a section's first line as it was is not asked to fix it.
func unframedSections(r review) string {
	var unframed []string
	for _, s := range r.outline.sections {
		if !s.prose || !s.required {
			continue
		}
		line, found := firstLineOfSection(r.body, s.name)
		if !found {
			continue
		}
		if r.editing {
			if was, _ := firstLineOfSection(r.before, s.name); was == line {
				continue
			}
		}
		line, skip := framingOf(line)
		if skip {
			continue
		}
		if line == "" || codeOpening.MatchString(line) || codeToken.MatchString(strings.Fields(line)[0]) {
			unframed = append(unframed, fmt.Sprintf("%s opens with %q", s.name, truncate(line, 60)))
		}
	}
	if len(unframed) == 0 {
		return ""
	}
	return fmt.Sprintf("This PR description opens with detail before framing: %s. Per the bury-the-lede skill, each section opens with one or two plain sentences saying what the reader is about to look at and why it matters; identifiers, paths, code and lists come after.", strings.Join(unframed, "; "))
}

// framingOf drops a leading ticket reference, which opens a section without
// framing it, and reports skip when nothing is left to judge.
func framingOf(line string) (rest string, skip bool) {
	if notApplicable.MatchString(line) {
		return "", true
	}
	rest = ticketRef.ReplaceAllString(line, "")
	return rest, line != "" && rest == ""
}

func truncate(s string, n int) string {
	if r := []rune(s); len(r) > n {
		return string(r[:n]) + "…"
	}
	return s
}
