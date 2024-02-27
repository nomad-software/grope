package image

import "embed"

var (
	//go:embed png
	//go:embed gif
	FS embed.FS
)
