package floatvar

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/widget" // IGNORE
)

type stub struct { // IGNORE
	widget.Widget        // IGNORE
	valueVar      string // IGNORE
}                             // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetValue sets the value.
func (el stub) SetValue(val float64) {
	tk.Get().SetVarFloatValue(el.valueVar, val)
}

// GetValue gets the value.
func (el stub) GetValue() float64 {
	return tk.Get().GetVarFloatValue(el.valueVar)
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el stub) Destroy() {
	el.Ele.Destroy()
	tk.Get().DestroyVar(el.valueVar)
}
