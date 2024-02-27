package scrollbar

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
	"github.com/nomad-software/goat/widget/scrollbar"
)

type stub struct{ element.Element } // IGNORE
func (el stub) GetID() string       { return "" } // IGNORE

// AttachHorizontalScrollbar attaches the horizontal scrollbar to this widget.
// Once the scrollbar is attached, the widget will also need attaching to the
// scrollbar to complete the attachment.
func (el stub) AttachHorizontalScrollbar(bar *scrollbar.HorizontalScrollbar) {
	tk.Get().Eval("%s configure -xscrollcommand [list %s set]", el.GetID(), bar.GetID())
}

// AttachVerticalScrollbar attaches the vertical scrollbar to this widget.
// Once the scrollbar is attached, the widget will also need attaching to the
// scrollbar to complete the attachment.
func (el stub) AttachVerticalScrollbar(bar *scrollbar.VerticalScrollbar) {
	tk.Get().Eval("%s configure -yscrollcommand [list %s set]", el.GetID(), bar.GetID())
}
