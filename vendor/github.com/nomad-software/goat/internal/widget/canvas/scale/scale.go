package scale

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// Scale scales the item.
// originX and originY identify the origin for the scaling operation and scaleX
// and scaleY identify the scale factors for x and y coordinates, respectively
// (a scale factor of 1.0 implies no change to that coordinate).
func (el stub) Scale(originX, originY, scaleX, scaleY float64) {
	tk.Get().Eval("%s scale %s %v %v %v %v", el.GetParent().GetID(), el.GetID(), originX, originX, scaleX, scaleY)
}
