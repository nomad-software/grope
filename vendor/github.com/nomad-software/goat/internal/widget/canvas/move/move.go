package move

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// MoveBy moves the item by a set amount.
func (el stub) MoveBy(x, y float64) {
	tk.Get().Eval("%s move %s %v %v", el.GetParent().GetID(), el.GetID(), x, y)
}

// MoveTo moves the item to a new coordinate.
func (el stub) MoveTo(x, y float64) {
	tk.Get().Eval("%s moveto %s %v %v", el.GetParent().GetID(), el.GetID(), x, y)
}
