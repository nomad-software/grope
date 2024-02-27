package panedwindow

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "panedwindow"
)

// A paned window widget displays a number of subwindows, stacked either
// vertically or horizontally. The user may adjust the relative sizes of the
// subwindows by dragging the sash between panes.
//
// Virtual events that can also be bound to.
// <<EnteredChild>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_panedwindow.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*PanedWindow -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*PanedWindow -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*PanedWindow -pkg=common/width
type PanedWindow struct {
	widget.Widget
}

// New creates a new paned window.
// See [option.orientation] for orientation strings.
func New(parent element.Element, orientation string) *PanedWindow {
	paned := &PanedWindow{}
	paned.SetParent(parent)
	paned.SetType(Type)

	tk.Get().Eval("ttk::panedwindow %s -orient %s", paned.GetID(), orientation)

	return paned
}

// AddPane adds a widget to a pane.
func (el *PanedWindow) AddPane(e element.Element) {
	tk.Get().Eval("%s insert end %s", el.GetID(), e.GetID())
}

// InsertPane inserts a widget to a pane at the specified index.
func (el *PanedWindow) InsertPane(index int, e element.Element) {
	tk.Get().Eval("%s insert %d %s", el.GetID(), index, e.GetID())
}

// RemovePane removes a pane.
func (el *PanedWindow) RemovePane(index int) {
	tk.Get().Eval("%s forget %d", el.GetID(), index)
}

// SetPaneWeight sets the pane weight.
// Weight is an integer specifying the relative stretchability of the pane.
// When the paned window is resized, the extra space is added or subtracted to
// each pane proportionally to its weight.
func (el *PanedWindow) SetPaneWeight(index, weight int) {
	tk.Get().Eval("%s pane %d -weight %d", el.GetID(), index, weight)
}

// SetSashPos sets the sash between pane's position.
// May adjust the positions of adjacent sashes to ensure that positions are
// monotonically increasing.  Sash positions are further constrained to be
// between 0 and the total size of the widget. Must be called after the UI has
// been drawn.
func (el *PanedWindow) SetSashPos(sashIndex, pos int) {
	tk.Get().Eval("%s sashpos %d %d", el.GetID(), sashIndex, pos)
}

// GetSashPos gets the position of the sash.
// Must be called after the UI has been drawn.
func (el *PanedWindow) GetSashPos(sashIndex int) int {
	tk.Get().Eval("%s sashpos %d", el.GetID(), sashIndex)
	return tk.Get().GetIntResult()
}
