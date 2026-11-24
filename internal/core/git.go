package core

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

var (
	ErrNoGit     = errors.New("git not found, beaver-blueprint requires git")
	ErrNotGitDir = errors.New("not a git repository")
)

// RequireGit checks that git is on PATH and dir is inside a work tree.
func RequireGit(dir string) error {
	if _, err := exec.LookPath("git"); err != nil {
		return ErrNoGit
	}
	out, err := git(dir, "rev-parse", "--is-inside-work-tree")
	if err != nil || out != "true" {
		return ErrNotGitDir
	}
	return nil
}

// git runs a git command in dir and returns trimmed stdout.
func git(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		if msg := strings.TrimSpace(stderr.String()); msg != "" {
			return "", errors.New(msg)
		}
		return "", err
	}
	return strings.TrimSpace(stdout.String()), nil
}
