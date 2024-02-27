package state

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetState sets the state.
// See [option.state] for state values.
func (el stub) SetState(state string) {
	tk.Get().Eval("%s itemconfigure %s -state {%s}", el.GetParent().GetID(), el.GetID(), state)
}
