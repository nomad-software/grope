// Code generated by tooling; DO NOT EDIT.
package label

import (
	"github.com/nomad-software/goat/internal/tk"

)







// SetText sets the text.
func (el *Label) SetText(text string) {
	tk.Get().SetVarStrValue(el.textVar, text)
}

// GetText gets the text.
func (el *Label) GetText() string {
	return tk.Get().GetVarStrValue(el.textVar)
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el *Label) Destroy() {
	el.Ele.Destroy()
	tk.Get().DestroyVar(el.textVar)
}