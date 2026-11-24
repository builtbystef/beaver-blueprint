package core

// Severity of a validation diagnostic.
type Severity int

const (
	Warning Severity = iota
	Error
)

// Diagnostic is one finding from Check.
type Diagnostic struct {
	Severity Severity
	Path     string
	Message  string
}

// CheckOptions controls Check.
type CheckOptions struct {
	Strict bool // promote warnings to errors
	Fix    bool // apply mechanical repairs
}

// Check validates the whole blueprint.
func (b *Blueprint) Check(opts CheckOptions) ([]Diagnostic, error) {
	return nil, ErrNotImplemented
}
