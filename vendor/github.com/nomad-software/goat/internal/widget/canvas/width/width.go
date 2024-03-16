package width

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetOutlineWidth sets the outline width.
func (el stub) SetOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the hover outline width.
func (el stub) SetHoverOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -activewidth %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the disabled outline width.
func (el stub) SetDisabledOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -disabledwidth %v", el.GetParent().GetID(), el.GetID(), width)
}
