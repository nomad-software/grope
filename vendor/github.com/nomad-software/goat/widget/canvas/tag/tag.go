package tag

import (
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvastag"
)

// Tag represents a tag in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/anchor
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/scale
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/width
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Tag -pkg=canvas/zorder
type Tag struct {
	element.Ele
}

// Creates a new tag.
func New(parent element.Element) *Tag {
	tag := &Tag{}
	tag.SetParent(parent)
	tag.SetType(Type)

	return tag
}
