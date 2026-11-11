package main

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	heading         = regexp.MustCompile(`(?m)(?:^|["'])\s*#{1,6}\s+(.+?)["']?\s*$`)
	optionalSection = regexp.MustCompile(`(?i)\bdelete (this|the) section\b|\bremove (this|the) section\b`)
	proseSlot       = regexp.MustCompile(`(?s)^(<!--.*-->)?$`)
	ticketSlot      = regexp.MustCompile(`^<!--\s*[A-Z][A-Z0-9]+-[X0-9]+\s*-->$`)

	prTemplatePaths = []string{
		".github/pull_request_template.md",
		".github/PULL_REQUEST_TEMPLATE.md",
		"pull_request_template.md",
		"PULL_REQUEST_TEMPLATE.md",
		"docs/pull_request_template.md",
	}
)

type section struct {
	name     string
	required bool
	prose    bool
}

// outline is the structure a PR body must follow: the repo's PR template when
// it has one, otherwise the skill's What, Why and How.
type outline struct {
	source   string
	template bool
	sections []section
}

var skillOutline = outline{source: "the bury-the-lede skill", sections: []section{
	{name: "What", required: true, prose: true},
	{name: "Why", required: true, prose: true},
	{name: "How", required: true, prose: true},
}}

// outlineFor reports false when -R names a repo whose template can't be read:
// guessing would judge the body by the wrong outline.
func outlineFor(root, repo string, gh gitHub) (outline, bool) {
	if repo != "" {
		for _, p := range prTemplatePaths {
			template, err := gh.file(repo, p)
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			if err != nil {
				return outline{}, false
			}
			return templateOutline(repo+"'s PR template ("+p+")", template), true
		}
		return skillOutline, true
	}
	if root != "" {
		for _, p := range prTemplatePaths {
			if data, err := os.ReadFile(filepath.Join(root, p)); err == nil {
				return templateOutline("this repo's PR template ("+p+")", string(data)), true
			}
		}
	}
	return skillOutline, true
}

// templateOutline reads each section's role from its placeholder: one that
// says it may be deleted is optional, and one that is empty or a comment
// (other than a ticket-key placeholder) is prose the author writes.
func templateOutline(source, template string) outline {
	o := outline{source: source, template: true}
	matches := heading.FindAllStringSubmatchIndex(template, -1)
	for i, m := range matches {
		end := len(template)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		placeholder := strings.TrimSpace(template[m[1]:end])
		o.sections = append(o.sections, section{
			name:     template[m[2]:m[3]],
			required: !optionalSection.MatchString(placeholder),
			prose:    proseSlot.MatchString(placeholder) && !ticketSlot.MatchString(placeholder),
		})
	}
	return o
}

// keptBy narrows the outline to the sections body has, so an edit is held to
// the sections already there rather than to ones the author removed.
func (o outline) keptBy(body string) outline {
	var kept []section
	for _, s := range o.sections {
		if _, found := firstLineOfSection(body, s.name); found {
			kept = append(kept, s)
		}
	}
	o.sections = kept
	return o
}

func firstLineOfSection(body, name string) (line string, found bool) {
	lines := strings.Split(body, "\n")
	for i, l := range lines {
		m := heading.FindStringSubmatch(l)
		if m == nil || !strings.EqualFold(m[1], name) {
			continue
		}
		for _, next := range lines[i+1:] {
			next = strings.Trim(strings.TrimSpace(next), `"'`)
			if heading.MatchString(next) {
				return "", true
			}
			if next == "" || (strings.HasPrefix(next, "<!--") && strings.HasSuffix(next, "-->")) {
				continue
			}
			return next, true
		}
		return "", true
	}
	return "", false
}

func headingsIn(body string) map[string]bool {
	present := map[string]bool{}
	for _, m := range heading.FindAllStringSubmatch(body, -1) {
		present[strings.ToLower(m[1])] = true
	}
	return present
}
