// Package core holds all Beaver Blueprint logic: discovery, parsing,
// writing, schema, validation, graph resolution, and drift. The CLI and the
// HTTP server are thin interfaces over it.
package core

import "errors"

// ErrNotImplemented marks scaffolded functionality that has not been written yet.
var ErrNotImplemented = errors.New("not implemented")

// SchemaVersion is the only supported value of `schema` in overview.md.
const SchemaVersion = 1
