package core

import "github.com/bmatcuk/doublestar/v4"

// Stale is an anchored element or document whose code changed after it.
type Stale struct {
	ID        string
	Committed []string // files changed in commits since the element
	Working   []string // modified or untracked files in the work tree
}

// Drift computes staleness for every anchored element and document.
func (b *Blueprint) Drift() ([]Stale, error) {
	return nil, ErrNotImplemented
}

// MatchAnchor reports whether a project-relative path matches an anchor glob.
func MatchAnchor(anchor, path string) bool {
	ok, err := doublestar.Match(anchor, path)
	return err == nil && ok
}
