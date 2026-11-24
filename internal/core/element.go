package core

import "go.yaml.in/yaml/v4"

// Element is one graph node, stored as blueprint/<type-dir>/<id>.md.
type Element struct {
	ID          string
	Type        Type
	Path        string     // relative to the project root
	Frontmatter *yaml.Node // decoded in place so unknown keys and order survive
	Body        []byte     // written back untouched unless edited
	Hash        string     // content hash for conflict detection
}

// Document is a singleton document or a type-directory README.
type Document struct {
	Name        string
	Frontmatter *yaml.Node
	Body        []byte
	Hash        string
}
