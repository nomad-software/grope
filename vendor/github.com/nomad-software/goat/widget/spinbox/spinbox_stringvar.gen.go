// Code generated by tooling; DO NOT EDIT.
package spinbox

import (
	"github.com/nomad-software/goat/internal/tk"

)







// SetValue sets the value.
func (el *Spinbox) SetValue(val string) {
	tk.Get().SetVarStrValue(el.valueVar, val)
}

// GetValue gets the value.
func (el *Spinbox) GetValue() string {
	return tk.Get().GetVarStrValue(el.valueVar)
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el *Spinbox) Destroy() {
	el.Ele.Destroy()
	tk.Get().DestroyVar(el.valueVar)
}