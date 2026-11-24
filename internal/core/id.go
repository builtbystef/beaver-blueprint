package core

import "regexp"

var idPattern = regexp.MustCompile(`^[a-z][a-z0-9-]{0,59}$`)

// ValidID reports whether s is a valid element id: lowercase ASCII letters,
// digits, and '-', starting with a letter, at most 60 characters.
func ValidID(s string) bool {
	return idPattern.MatchString(s)
}
