package rectangle

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasrectangle"
)

// Rectangle represents a rectangle in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/scale
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/width
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Rectangle -pkg=canvas/zorder
type Rectangle struct {
	element.Ele
}

// Creates a new rectangle.
func New(parent element.Element) *Rectangle {
	rect := &Rectangle{}
	rect.SetParent(parent)
	rect.SetType(Type)

	return rect
}

// SetCoords updates the item coordinates.
func (el *Rectangle) SetCoords(x1, y1, x2, y2 float64) {
	tk.Get().Eval("%s coords %s [list %v %v %v %v]", el.GetParent().GetID(), el.GetID(), x1, y1, x2, y2)
}
