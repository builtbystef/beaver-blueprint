package core

// Render serializes an element with canonical frontmatter field order.
// The body bytes are emitted unchanged.
func Render(e *Element) ([]byte, error) {
	return nil, ErrNotImplemented
}

// NewElement creates an element file with correct frontmatter and headings.
func (b *Blueprint) NewElement(t Type, id, title string) (*Element, error) {
	return nil, ErrNotImplemented
}

// Move renames an element and rewrites every reference to it.
func (b *Blueprint) Move(oldID, newID string) error {
	return ErrNotImplemented
}

// Remove deletes an element. With force, references to it are stripped;
// without, it fails if the element is referenced.
func (b *Blueprint) Remove(id string, force bool) error {
	return ErrNotImplemented
}

// Verify sets `verified` to HEAD on each named element or document.
func (b *Blueprint) Verify(ids ...string) error {
	return ErrNotImplemented
}
