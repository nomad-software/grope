package radiobutton

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "radiobutton"
)

// Radio button widgets are used in groups to show or change a set of
// mutually-exclusive options. Radio buttons have an associated selected value;
// when a radio button is selected, it sets the associated value.
//
// To create a group of radio button that work properly in unison, all radio
// button widgets within the group must share the same immediate parent
// (usually a frame) and all must have individual selected values set.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_radiobutton.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/command
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/image
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/invoke
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/stringvar -methods=GetValue,SetValue
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/textvar -methods=GetText,SetText
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*RadioButton -pkg=common/width
type RadioButton struct {
	widget.Widget

	textVar  string
	valueVar string

	// This is the value of the above value variable when this radio button is
	// selected. This should be different for all radio buttons in the same
	// group.
	selectedValue string
}

// New creates a new radio button.
func New(parent element.Element, text string) *RadioButton {
	button := &RadioButton{}
	button.SetParent(parent)
	button.SetType(Type)

	button.textVar = variable.GenerateName(button.GetID())

	if parent != nil {
		button.valueVar = variable.GenerateName(button.GetType(), button.GetParent().GetID())
	} else {
		button.valueVar = variable.GenerateName(button.GetType())
	}

	tk.Get().Eval("ttk::radiobutton %s -textvariable %s -variable %s", button.GetID(), button.textVar, button.valueVar)

	button.SetText(text)
	button.SetSelectedValue(text)

	return button
}

// GetSelectedValue gets this radio button's selected value.
func (el *RadioButton) GetSelectedValue() string {
	return el.selectedValue
}

// SetSelectedValue sets this radio button's selected value.
func (el *RadioButton) SetSelectedValue(value string) {
	el.selectedValue = value
	tk.Get().Eval("%s configure -value {%s}", el.GetID(), el.selectedValue)
}

// Select selects the radio button.
func (el *RadioButton) Select() {
	tk.Get().SetVarStrValue(el.valueVar, el.selectedValue)
}

// IsSelected return true if the radio button is selected.
func (el *RadioButton) IsSelected() bool {
	return tk.Get().GetVarStrValue(el.valueVar) == el.selectedValue
}
