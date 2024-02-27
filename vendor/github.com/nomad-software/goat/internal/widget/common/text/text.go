package text

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetText sets the text.
func (el stub) SetText(text string) {
	tk.Get().Eval("%s configure -text {%s}", el.GetID(), text)
}

// GetText gets the text.
func (el stub) GetText() string {
	tk.Get().Eval("%s cget -text", el.GetID())
	return tk.Get().GetStrResult()
}
