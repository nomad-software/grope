// Code generated by tooling; DO NOT EDIT.
package menu

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetBorderWidth sets the border width.
func (el *Menu) SetBorderWidth(b int) {
	tk.Get().Eval("%s configure -borderwidth %d", el.GetID(), b)
}
