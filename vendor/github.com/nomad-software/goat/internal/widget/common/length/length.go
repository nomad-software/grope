package length

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetLength sets the length.
func (el stub) SetLength(l int) {
	tk.Get().Eval("%s configure -length %d", el.GetID(), l)
}
