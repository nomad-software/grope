package show

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// Show shows the text as a sequence of the specified characters.
// This is useful to hide the text value.
func (el stub) ShowCharsAs(r rune) {
	tk.Get().Eval("%s configure -show {%s}", el.GetID(), string(r))
}
