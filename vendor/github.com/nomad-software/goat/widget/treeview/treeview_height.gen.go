// Code generated by tooling; DO NOT EDIT.
package treeview

import (
	"github.com/nomad-software/goat/internal/tk"

)



// SetHeight sets the height.
func (el *TreeView) SetHeight(h int) {
	tk.Get().Eval("%s configure -height %d", el.GetID(), h)
}