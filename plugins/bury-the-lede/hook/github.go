package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

// gitHub is what the hook reads from GitHub. A failed read never denies a
// command by itself.
type gitHub interface {
	prBody(selector, repo, dir string) (string, bool)
	// file returns os.ErrNotExist when repo has no such file.
	file(repo, path string) (string, error)
}

// ghCLI reads through the gh CLI with the user's own auth. Without gh,
// network or auth every read fails, and the hook falls back or stands aside.
type ghCLI struct{}

func (ghCLI) prBody(selector, repo, dir string) (string, bool) {
	args := []string{"pr", "view"}
	if selector != "" {
		args = append(args, selector)
	}
	if repo != "" {
		args = append(args, "--repo", repo)
	}
	gh := exec.Command("gh", append(args, "--json", "body", "--jq", ".body")...)
	gh.Dir = dir
	out, err := gh.Output()
	if err != nil {
		return "", false
	}
	return string(out), true
}

func (ghCLI) file(repo, path string) (string, error) {
	out, err := exec.Command("gh", "api", "repos/"+repo+"/contents/"+path, "-H", "Accept: application/vnd.github.raw").Output()
	var exit *exec.ExitError
	if errors.As(err, &exit) && bytes.Contains(exit.Stderr, []byte("HTTP 404")) {
		return "", os.ErrNotExist
	}
	return string(out), err
}
