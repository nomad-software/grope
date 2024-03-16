package button

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "button"
)

// A button is a widget that issues a command when pressed
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_button.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/command
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/default
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/image
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/invoke
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/textvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*Button -pkg=common/width
type Button struct {
	widget.Widget

	textVar string
}

// New creates a new button.
func New(parent element.Element, text string) *Button {
	button := &Button{}
	button.SetParent(parent)
	button.SetType(Type)

	button.textVar = variable.GenerateName(button.GetID())

	tk.Get().Eval("ttk::button %s -textvariable %s", button.GetID(), button.textVar)

	button.SetText(text)

	return button
}
