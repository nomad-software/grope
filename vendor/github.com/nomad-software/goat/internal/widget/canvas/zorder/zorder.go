package zorder

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

type stub struct{ element.Element } // IGNORE

// Lower will lower this item below the specified other or if other is nil,
// below all others.
func (el stub) Lower(other element.Element) {
	if other != nil {
		tk.Get().Eval("%s lower %s %s", el.GetParent().GetID(), el.GetID(), other.GetID())
	} else {
		tk.Get().Eval("%s lower %s", el.GetParent().GetID(), el.GetID())
	}
}

// Raise will raise this item above the specified other or if other is nil,
// above all others.
func (el stub) Raise(other element.Element) {
	if other != nil {
		tk.Get().Eval("%s raise %s %s", el.GetParent().GetID(), el.GetID(), other.GetID())
	} else {
		tk.Get().Eval("%s raise %s", el.GetParent().GetID(), el.GetID())
	}
}
