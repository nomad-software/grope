package padding

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetPadding sets the padding.
func (el stub) SetPadding(p int) {
	tk.Get().Eval("%s configure -padding %d", el.GetID(), p)
}
