// Code generated by tooling; DO NOT EDIT.
package combobox

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetWidth sets the width.
func (el *Combobox) SetWidth(w int) {
	tk.Get().Eval("%s configure -width %d", el.GetID(), w)
}