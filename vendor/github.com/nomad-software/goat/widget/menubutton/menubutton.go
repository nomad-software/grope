package menubutton

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
	"github.com/nomad-software/goat/widget/menu"
)

const (
	Type = "menubutton"
)

// A menu button is a widget that displays a menu when clicked.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_menubutton.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuButton -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuButton -pkg=common/image
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuButton -pkg=common/invoke
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuButton -pkg=common/textvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuButton -pkg=common/underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuButton -pkg=common/width
type MenuButton struct {
	widget.Widget

	textVar string
}

// New creates a new button.
func New(parent element.Element, text string, menu *menu.Menu) *MenuButton {
	button := &MenuButton{}
	button.SetParent(parent)
	button.SetType(Type)

	button.textVar = variable.GenerateName(button.GetID())

	tk.Get().Eval("ttk::menubutton %s -textvariable %s -menu %s", button.GetID(), button.textVar, menu.GetID())

	button.SetText(text)

	return button
}

// SetMenu sets the menu to show when clicked.
func (el *MenuButton) SetMenu(menu *menu.Menu) {
	tk.Get().Eval("%s configure -menu %s", el.GetID(), menu.GetID())
}

// SetMenuDirection set the direction of the menu when showing.
func (el *MenuButton) SetMenuDirection(direction string) {
	tk.Get().Eval("%s configure -direction %s", el.GetID(), direction)
}
