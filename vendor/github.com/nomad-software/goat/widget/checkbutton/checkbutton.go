package checkbutton

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "checkbutton"
)

// A checkbutton widget is used to show or change a setting. It has two states,
// selected and deselected. The state of the checkbutton may be linked to a
// value.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_checkbutton.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/command
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/image
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/invoke
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/stringvar -methods=GetValue,SetValue
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/textvar -methods=GetText,SetText
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*CheckButton -pkg=common/width
type CheckButton struct {
	widget.Widget

	textVar  string
	valueVar string
}

// New creates a new checkbutton.
func New(parent element.Element, text string) *CheckButton {
	button := &CheckButton{}
	button.SetParent(parent)
	button.SetType(Type)

	button.textVar = variable.GenerateName("textvar", button.GetID())
	button.valueVar = variable.GenerateName("valuevar", button.GetID())

	tk.Get().Eval("ttk::checkbutton %s -textvariable %s -variable %s", button.GetID(), button.textVar, button.valueVar)

	button.SetText(text)
	button.SetValue("0")

	return button
}

// Check checks the checkbutton.
func (el *CheckButton) Check() {
	tk.Get().SetVarStrValue(el.valueVar, "1")
}

// Check unchecks the checkbutton.
func (el *CheckButton) UnCheck() {
	tk.Get().SetVarStrValue(el.valueVar, "0")
}

// Check half-checks the checkbutton.
func (el *CheckButton) HalfCheck() {
	tk.Get().SetVarStrValue(el.valueVar, "")
}

// IsChecked returns true if the checkbutton is checked.
func (el *CheckButton) IsChecked() bool {
	return tk.Get().GetVarBoolValue(el.valueVar)
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el *CheckButton) Destroy() {
	el.Ele.Destroy()
	tk.Get().DestroyVar(el.textVar)
	tk.Get().DestroyVar(el.valueVar)
}
