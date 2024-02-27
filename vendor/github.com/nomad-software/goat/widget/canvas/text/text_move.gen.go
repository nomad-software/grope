// Code generated by tooling; DO NOT EDIT.
package text

import (
	"github.com/nomad-software/goat/internal/tk"

)



// MoveBy moves the item by a set amount.
func (el *Text) MoveBy(x, y float64) {
	tk.Get().Eval("%s move %s %v %v", el.GetParent().GetID(), el.GetID(), x, y)
}

// MoveTo moves the item to a new coordinate.
func (el *Text) MoveTo(x, y float64) {
	tk.Get().Eval("%s moveto %s %v %v", el.GetParent().GetID(), el.GetID(), x, y)
}
