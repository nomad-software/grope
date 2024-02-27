package underline

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetUnderline sets the character which is underlined.
// See [option.underline] for options.
func (el stub) SetUnderline(index int) {
	tk.Get().Eval("%s configure -underline %d", el.GetID(), index)
}
