package combobox

import (
	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/state"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "combobox"
)

// A combobox combines a text field with a pop-down list of values; the user
// may select the value of the text field from among the values in the list.
//
// This widget has two types of values that can be set. First, a list of values
// can be set to populate the drop-down list which can then be selected via a
// mouse. Second, the value can be set independently and in addition to the
// value list.
//
// Virtual events that can also be bound to.
// <<ComboboxSelected>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_combobox.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/color -methods=SetForegroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/data
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/font
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/justify
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/scrollbar -methods=AttachHorizontalScrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/selection
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/stringvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Combobox -pkg=common/width
type Combobox struct {
	widget.Widget

	valueVar string
}

// New creates a new combobox.
func New(parent element.Element) *Combobox {
	combo := &Combobox{}
	combo.SetParent(parent)
	combo.SetType(Type)

	combo.valueVar = variable.GenerateName(combo.GetID())

	tk.Get().Eval("ttk::combobox %s -textvariable %s", combo.GetID(), combo.valueVar)

	combo.SetState([]string{state.Readonly})
	combo.Bind("<<ComboboxSelected>>", func(*command.BindData) {
		combo.DeselectText()
	})

	return combo
}

// GetSelectedIndex gets the selected index of the combobox.
func (el *Combobox) GetSelectedIndex() int {
	tk.Get().Eval("%s current", el.GetID())
	return tk.Get().GetIntResult()
}

// SetSelectedIndex sets the selected index of the combobox.
func (el *Combobox) SetSelectedIndex(index int) {
	tk.Get().Eval("%s current %d", el.GetID(), index)
}
