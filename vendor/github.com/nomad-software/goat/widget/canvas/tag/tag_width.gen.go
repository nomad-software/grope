// Code generated by tooling; DO NOT EDIT.
package tag

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetOutlineWidth sets the outline width.
func (el *Tag) SetOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the hover outline width.
func (el *Tag) SetHoverOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -activewidth %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetOutlineWidth sets the disabled outline width.
func (el *Tag) SetDisabledOutlineWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -disabledwidth %v", el.GetParent().GetID(), el.GetID(), width)
}
