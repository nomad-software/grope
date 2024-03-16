package border

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetRelief sets the relief effect.
// See [option.relief] for relief values.
func (el stub) SetRelief(r string) {
	tk.Get().Eval("%s configure -relief {%s}", el.GetID(), r)
}
