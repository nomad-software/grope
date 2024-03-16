package default_

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetDefault sets this wiget as the default.
func (el stub) SetDefault() {
	tk.Get().Eval("%s configure -default active", el.GetID())
}
