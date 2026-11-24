package core

// Frontmatter field structs are the single source of truth for the schema:
// Huma reflects them into the OpenAPI document's components, and
// `beaver-blueprint schema` prints them from the embedded document.

// Common holds the fields every element has.
type Common struct {
	Title    string   `yaml:"title" json:"title"`
	Summary  string   `yaml:"summary,omitempty" json:"summary,omitempty"`
	Verified string   `yaml:"verified,omitempty" json:"verified,omitempty"`
	Code     []string `yaml:"code,omitempty" json:"code,omitempty"`
}

// TODO: per-type frontmatter structs (UseCaseFields, EntityFields, ...).
