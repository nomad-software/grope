package textvar

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/widget" // IGNORE
)

type stub struct { // IGNORE
	widget.Widget        // IGNORE
	textVar       string // IGNORE
}                             // IGNORE
func (el stub) GetID() string { return "" } // IGNORE

// SetText sets the text.
func (el stub) SetText(text string) {
	tk.Get().SetVarStrValue(el.textVar, text)
}

// GetText gets the text.
func (el stub) GetText() string {
	return tk.Get().GetVarStrValue(el.textVar)
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el stub) Destroy() {
	el.Ele.Destroy()
	tk.Get().DestroyVar(el.textVar)
}
