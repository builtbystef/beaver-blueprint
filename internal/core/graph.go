package core

// Edge is a typed, directed relationship stored in From's frontmatter.
type Edge struct {
	From  string
	To    string
	Field string // e.g. "depends_on", "owned_by"
}

// Graph is the resolved relationship graph with inverses derived.
type Graph struct {
	Edges []Edge
}

// Graph resolves every relationship in the blueprint.
func (b *Blueprint) Graph() (*Graph, error) {
	return nil, ErrNotImplemented
}
