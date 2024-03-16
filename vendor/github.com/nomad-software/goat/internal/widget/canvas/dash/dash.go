package dash

import (
	"fmt"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetOutlineDash sets the outline dash.
// Each element represents the number of pixels of a line segment. Only the odd
// segments are drawn using the “outline” color. The other segments are drawn
// transparent.
func (el stub) SetOutlineDash(first, second float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s itemconfigure %s -dash [list %v %v %s]", el.GetParent().GetID(), el.GetID(), first, second, otherStr)
}

// SetHoverOutlineDash sets the hover outline dash.
func (el stub) SetHoverOutlineDash(first, second float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s itemconfigure %s -activedash [list %v %v %s]", el.GetParent().GetID(), el.GetID(), first, second, otherStr)
}

// SetDisabledOutlineDash sets the disabled outline dash.
func (el stub) SetDisabledOutlineDash(first, second float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s itemconfigure %s -disableddash [list %v %v %s]", el.GetParent().GetID(), el.GetID(), first, second, otherStr)
}

// SetOutlineDashOffset sets the starting offset in pixels.
func (el stub) SetOutlineDashOffset(offset float64) {
	tk.Get().Eval("%s itemconfigure %s -dashoffset %v", el.GetParent().GetID(), el.GetID(), offset)
}
