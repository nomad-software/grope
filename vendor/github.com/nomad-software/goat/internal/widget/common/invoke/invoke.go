package invoke

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// Invoke invokes the command associated with this widget.
func (el stub) Invoke() {
	tk.Get().Eval("%s invoke", el.GetID())
}
