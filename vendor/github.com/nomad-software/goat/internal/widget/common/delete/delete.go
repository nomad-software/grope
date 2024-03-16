package delete_

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// DeleteText deletes all text.
func (el stub) DeleteText(start, end int) {
	tk.Get().Eval("%s delete %d %d", el.GetID(), start, end)
}
