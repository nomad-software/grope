package widget

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvaswidget"
)

// Widget represents a widget in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/anchor
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Widget -pkg=canvas/zorder
type Widget struct {
	element.Ele
}

// Creates a new widget.
func New(parent element.Element) *Widget {
	widget := &Widget{}
	widget.SetParent(parent)
	widget.SetType(Type)

	return widget
}

// SetWidth sets the width
func (el *Widget) SetWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetHeight sets the width
func (el *Widget) SetHeight(height float64) {
	tk.Get().Eval("%s itemconfigure %s -height %v", el.GetParent().GetID(), el.GetID(), height)
}

// SetWidget sets the widget.
func (el *Widget) SetWidget(e element.Element) {
	tk.Get().Eval("%s itemconfigure %s -window %s", el.GetParent().GetID(), el.GetID(), e.GetID())
}

// SetCoords updates the item coordinates.
func (el *Widget) SetCoords(x, y float64) {
	tk.Get().Eval("%s coords %s [list %v %v]", el.GetParent().GetID(), el.GetID(), x, y)
}
