package fontdialog

import (
	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "fontdialog"
)

// FontDialog is a dialog box used to choose a font.
//
// Virtual events that can also be bound to.
// <<TkFontchooserVisibility>>
// <<TkFontchooserFontChanged>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/fontchooser.html
type FontDialog struct {
	element.Ele

	title string
}

// New creates a new message dialog.
func New(parent element.Element, title string) *FontDialog {
	dialog := &FontDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)

	tk.Get().Eval(
		"tk fontchooser configure -parent %s -title {%s}",
		dialog.GetParent().GetID(),
		dialog.title,
	)

	return dialog
}

// SetTitle sets the dialog title.
func (el *FontDialog) SetTitle(title string) {
	el.title = title
}

// Show creates and shows the dialog.
// This method call will block until the dialog is closed. Then the value can
// be read.
func (el *FontDialog) Show() {
	tk.Get().Eval("tk fontchooser show")
}

// SetCommand sets the command to execute on interaction with the dialog.
func (el *FontDialog) SetCommand(callback command.FontDialogCallback) {
	name := command.GenerateName(el.GetType())

	tk.Get().CreateFontDialogCommand(el, name, callback)
	tk.Get().Eval("tk fontchooser configure -command %s", name)
}

// DeleteCommand deletes the command.
func (el *FontDialog) DeleteCommand() {
	tk.Get().Eval("tk fontchooser configure -command {}")

	name := command.GenerateName(el.GetType())
	tk.Get().DestroyCommand(name)
}
