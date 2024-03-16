package height

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetHeight sets the height.
func (el stub) SetHeight(h int) {
	tk.Get().Eval("%s configure -height %d", el.GetID(), h)
}
