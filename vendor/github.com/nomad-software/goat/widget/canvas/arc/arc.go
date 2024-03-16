package arc

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasarc"
)

// Arc represents an arc in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/outline
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/scale
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/width
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Arc -pkg=canvas/zorder
type Arc struct {
	element.Ele
}

// Creates a new arc.
func New(parent element.Element) *Arc {
	arc := &Arc{}
	arc.SetParent(parent)
	arc.SetType(Type)

	return arc
}

// SetStart sets the start argument specifies the beginning of the angular
// range occupied by the arc. Degrees is given in units of degrees measured
// counter-clockwise from the 3-o'clock position; it may be either positive or
// negative.
func (el *Arc) SetStart(start float64) {
	tk.Get().Eval("%s itemconfigure %s -start %v", el.GetParent().GetID(), el.GetID(), start)
}

// SetExtent sets the extent argument specifies the size of the angular range
// occupied by the arc. The arc's range extends for degrees degrees
// counter-clockwise from the starting angle. Degrees may be negative. If it is
// greater than 360 or less than -360, then degrees modulo 360 is used as the
// extent.
func (el *Arc) SetExtent(extent float64) {
	tk.Get().Eval("%s itemconfigure %s -extent %v", el.GetParent().GetID(), el.GetID(), extent)
}

// SetStyle sets the style of arc.
// See [widget.canvas.arc.style] for style names.
func (el *Arc) SetStyle(style string) {
	tk.Get().Eval("%s itemconfigure %s -style {%s}", el.GetParent().GetID(), el.GetID(), style)
}

// SetCoords updates the item coordinates.
func (el *Arc) SetCoords(x1, y1, x2, y2 float64) {
	tk.Get().Eval("%s coords %s [list %v %v %v %v]", el.GetParent().GetID(), el.GetID(), x1, y1, x2, y2)
}
