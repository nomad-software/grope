package menu

import (
	"fmt"

	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui"
	"github.com/nomad-software/goat/internal/widget/ui/element/hash"
)

const (
	Type = "menu"
)

// Menubar is the cascading menu that items are selected from.
//
// Virtual events that can also be bound to.
// <<MenuSelect>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Menu -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Menu -pkg=common/borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*Menu -pkg=common/color -methods=SetForegroundColor,SetBackgroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*Menu -pkg=common/relief
type Menu struct {
	ui.Ele

	checkButtonVars []string
	radioButtonVar  string
}

// New creates a new menu.
// See [option.underline] for underline options.
func New(bar *MenuBar, label string, underline int) *Menu {
	menu := &Menu{}
	menu.SetParent(bar)
	menu.SetType(Type)

	menu.radioButtonVar = variable.GenerateName(label, menu.GetID())

	tk.Get().Eval("menu %s -type normal -tearoff 0", menu.GetID())
	tk.Get().Eval("%s add cascade -menu %s -label {%s} -underline %d", menu.GetParent().GetID(), menu.GetID(), label, underline)

	return menu
}

// NewPopUp creates a new popup menu that doesn't have a bar as a parent.
func NewPopUp() *Menu {
	menu := &Menu{}
	menu.SetType("popup-menu")

	menu.radioButtonVar = fmt.Sprintf("var-%s", hash.Generate())

	tk.Get().Eval("menu %s -type normal -tearoff 0", menu.GetID())

	return menu
}

// AddMenuEntry adds a cascading menu entry.
// See [option.underline] for underline options.
func (m *Menu) AddMenuEntry(menu *Menu, label string, underline int) {
	origId := menu.GetID()
	menu.SetParent(m)

	// Update the menu id.
	tk.Get().Eval("%s clone %s", origId, menu.GetID())
	tk.Get().Eval("%s add cascade -label {%s} -underline %d -menu %s", m.GetID(), label, underline, menu.GetID())
}

// AddEntry adds a menu entry with an optional cosmetic shortcut and a callback
// to execute when selected.
// The shortcut will need to be bound using the Bind method.
func (m *Menu) AddEntry(label string, shortcut string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(m, name, callback)

	tk.Get().Eval("%s add command -label {%s} -accelerator {%s} -command {%s}", m.GetID(), label, shortcut, name)
}

// AddImageEntry is the same as AddEntry but also displays an image.
// The shortcut will need to be bound using the Bind method.
// See [option.compound] for image positions.
func (m *Menu) AddImageEntry(img *image.Image, compound string, label string, shortcut string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(m, name, callback)

	tk.Get().Eval("%s add command -label {%s} -accelerator {%s} -image %s -compound {%s} -command {%s}", m.GetID(), label, shortcut, img.GetID(), compound, name)
}

// AddCheckButtonEntry adds an item to the menu that acts as a check button.
// The shortcut will need to be bound using the Bind method.
func (m *Menu) AddCheckButtonEntry(label string, shortcut string, callback command.Callback) {
	varName := variable.GenerateName(label, m.GetID())
	m.checkButtonVars = append(m.checkButtonVars, varName)

	cmdName := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(m, cmdName, callback)

	tk.Get().Eval("%s add checkbutton -variable %s -label {%s} -accelerator {%s} -command {%s}", m.GetID(), varName, label, shortcut, cmdName)
}

// AddImageCheckButtonEntry is the same as AddCheckButtonEntry but also
// displays an image.
// The shortcut will need to be bound using the Bind method.
// See [option.compound] for image positions.
func (m *Menu) AddImageCheckButtonEntry(img *image.Image, compound string, label string, shortcut string, callback command.Callback) {
	varName := variable.GenerateName(label, m.GetID())
	m.checkButtonVars = append(m.checkButtonVars, varName)

	cmdName := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(m, cmdName, callback)

	tk.Get().Eval("%s add checkbutton -variable %s -label {%s} -accelerator {%s} -image %s -compound {%s} -command {%s}", m.GetID(), varName, label, shortcut, img.GetID(), compound, cmdName)
}

// AddRadioButtonEntry adds an item to the menu that acts as a radio button.
// The shortcut will need to be bound using the Bind method.
func (m *Menu) AddRadioButtonEntry(label string, shortcut string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(m, name, callback)

	tk.Get().Eval("%s add radiobutton -variable %s -label {%s} -accelerator {%s} -command {%s}", m.GetID(), m.radioButtonVar, label, shortcut, name)
}

// AddImageRadioButtonEntry is the same as AddRadioButtonEntry but also
// displays an image.
// The shortcut will need to be bound using the Bind method.
// See [option.compound] for image positions.
func (m *Menu) AddImageRadioButtonEntry(img *image.Image, compound string, label string, shortcut string, callback command.Callback) {
	name := command.GenerateName(label, m.GetID())
	tk.Get().CreateCommand(m, name, callback)

	tk.Get().Eval("%s add radiobutton -variable %s -label {%s} -accelerator {%s} -image %s -compound {%s} -command {%s}", m.GetID(), m.radioButtonVar, label, shortcut, img.GetID(), compound, name)
}

// AddSeparator adds a separator.
func (m *Menu) AddSeparator() {
	tk.Get().Eval("%s add separator", m.GetID())
}

// DisableEntry disables an entry.
func (m *Menu) DisableEntry(index int) {
	tk.Get().Eval("%s entryconfigure %d -state disable", m.GetID(), index)
}

// EnableEntry enables an entry.
func (m *Menu) EnableEntry(index int) {
	tk.Get().Eval("%s entryconfigure %d -state normal", m.GetID(), index)
}

// Invoke invokes the entries callback.
func (m *Menu) Invoke(index int) {
	tk.Get().Eval("%s invoke %d", m.GetID(), index)
}

// SetCheckButtonEntry selects or deselects a check button entry at the
// specified index. This will also execute the callback.
func (m *Menu) SetCheckButtonEntry(index int, selected bool) {
	if index >= 0 && index < len(m.checkButtonVars) {
		name := m.checkButtonVars[index]
		if selected {
			tk.Get().SetVarStrValue(name, "1")
		} else {
			tk.Get().SetVarStrValue(name, "0")
		}
	}
}

// GetCheckButtonEntrySelected gets if the check button entry at the passed
// index is checked or not. The index only applies to check button entries in
// the menu not any other type of entry. If there are no check button entries
// in the menu this method returns false.
func (m *Menu) IsCheckButtonEntrySelected(index int) bool {
	if index >= 0 && index < len(m.checkButtonVars) {
		name := m.checkButtonVars[index]
		return tk.Get().GetVarBoolValue(name)
	}

	return false
}

// SelectRadioButtonEntry selects or deselects the radio button entry at the
// specified index. This will also execute the callback.
func (m *Menu) SelectRadioButtonEntry(index int) {
	tk.Get().Eval("%s entrycget %d -label", m.GetID(), index)
	label := tk.Get().GetStrResult()

	tk.Get().SetVarStrValue(m.radioButtonVar, label)
}

// GetSelectedRadioButtonEntry gets the value of the selected radio button
// entry. This value will be the same as the entry's label. This method will
// return an empty string if no radio button entry exists or none are selected.
func (m *Menu) GetSelectedRadioButtonEntry() string {
	return tk.Get().GetVarStrValue(m.radioButtonVar)
}

// PopUp shows a popup menu at the specified coords.
func (m *Menu) PopUp(x int, y int) {
	tk.Get().Eval("tk_popup %s %d %d", m.GetID(), x, y)
}
