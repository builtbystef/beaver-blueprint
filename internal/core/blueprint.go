package core

import (
	"errors"
	"os"
	"path/filepath"
)

// DirName is the blueprint directory at the project root.
const DirName = "blueprint"

// ErrNoBlueprint is returned when walk-up discovery finds no blueprint.
var ErrNoBlueprint = errors.New("no blueprint found, run beaver-blueprint init")

// Blueprint is a loaded blueprint directory.
type Blueprint struct {
	Root      string // project root, the directory containing blueprint/
	Elements  map[string]*Element
	Documents map[string]*Document
}

// Dir returns the absolute path of the blueprint directory.
func (b *Blueprint) Dir() string {
	return filepath.Join(b.Root, DirName)
}

// Find walks up from start to locate the project root containing blueprint/.
func Find(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		info, err := os.Stat(filepath.Join(dir, DirName))
		if err == nil && info.IsDir() {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", ErrNoBlueprint
		}
		dir = parent
	}
}

// Load reads every element and document under root/blueprint.
func Load(root string) (*Blueprint, error) {
	return nil, ErrNotImplemented
}

// Init scaffolds an empty blueprint skeleton at root.
func Init(root string) error {
	return ErrNotImplemented
}
