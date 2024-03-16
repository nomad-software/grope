// Code generated by tooling; DO NOT EDIT.
package entry

import (
	"github.com/nomad-software/goat/internal/tk"

	"github.com/nomad-software/goat/widget/scrollbar"
)




// AttachHorizontalScrollbar attaches the horizontal scrollbar to this widget.
// Once the scrollbar is attached, the widget will also need attaching to the
// scrollbar to complete the attachment.
func (el *Entry) AttachHorizontalScrollbar(bar *scrollbar.HorizontalScrollbar) {
	tk.Get().Eval("%s configure -xscrollcommand [list %s set]", el.GetID(), bar.GetID())
}

