package oval

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasoval"
)

// Oval represents an oval in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/scale
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/width
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Oval -pkg=canvas/zorder
type Oval struct {
	element.Ele
}

// Creates a new oval.
func New(parent element.Element) *Oval {
	oval := &Oval{}
	oval.SetParent(parent)
	oval.SetType(Type)

	return oval
}

// SetCoords updates the item coordinates.
func (el *Oval) SetCoords(x1, y1, x2, y2 float64) {
	tk.Get().Eval("%s coords %s [list %v %v %v %v]", el.GetParent().GetID(), el.GetID(), x1, y1, x2, y2)
}
