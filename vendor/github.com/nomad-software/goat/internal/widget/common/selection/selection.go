package selection

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SelectText selects the text according to the passed limits.
func (el stub) SelectText(start, end int) {
	tk.Get().Eval("%s selection range %d %d", el.GetID(), start, end)
}

// IsTextSelected returns true if text is selected.
func (el stub) IsTextSelected() bool {
	tk.Get().Eval("%s selection present", el.GetID())
	return tk.Get().GetBoolResult()
}

// DeselectText deselected all text.
func (el stub) DeselectText() {
	tk.Get().Eval("%s selection clear", el.GetID())
}
