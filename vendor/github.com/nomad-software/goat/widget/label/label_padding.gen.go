// Code generated by tooling; DO NOT EDIT.
package label

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetPadding sets the padding.
func (el *Label) SetPadding(p int) {
	tk.Get().Eval("%s configure -padding %d", el.GetID(), p)
}
