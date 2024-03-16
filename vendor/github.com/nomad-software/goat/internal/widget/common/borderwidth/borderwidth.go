package borderwidth

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetBorderWidth sets the border width.
func (el stub) SetBorderWidth(b int) {
	tk.Get().Eval("%s configure -borderwidth %d", el.GetID(), b)
}
