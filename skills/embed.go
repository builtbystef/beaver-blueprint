// Package skills embeds the Agent Skills installed by `beaver-blueprint init`.
package skills

import "embed"

//go:embed */SKILL.md
var FS embed.FS
