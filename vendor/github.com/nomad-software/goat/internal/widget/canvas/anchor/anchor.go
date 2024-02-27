package anchor

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetAnchor sets the anchor.
// See [option.anchor] for anchor values.
func (el stub) SetAnchor(anchor string) {
	tk.Get().Eval("%s itemconfigure %s -anchor {%s}", el.GetParent().GetID(), el.GetID(), anchor)
}
