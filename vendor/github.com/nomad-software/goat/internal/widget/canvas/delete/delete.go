package delete

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// Delete remove this item from the canvas.
func (el stub) Delete() {
	tk.Get().Eval("%s delete %s", el.GetParent().GetID(), el.GetID())
}
