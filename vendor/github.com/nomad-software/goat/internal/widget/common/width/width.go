package width

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetWidth sets the width.
func (el stub) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}
