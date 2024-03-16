package colordialog

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/color"
)

const (
	Type = "colordialog"
)

// ColorDialog is a dialog box used to choose a color.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/chooseColor.html
type ColorDialog struct {
	element.Ele

	title        string
	initialColor string

	value string
}

// New creates a new color dialog.
func New(parent element.Element, title string) *ColorDialog {
	dialog := &ColorDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)
	dialog.SetInitialColor(color.White)

	return dialog
}

// SetTitle sets the dialog title.
func (el *ColorDialog) SetTitle(title string) {
	el.title = title
}

// SetInitialColor sets the initial color of the dialog.
// See [option.color] for color names.
func (el *ColorDialog) SetInitialColor(color string) {
	el.initialColor = color
}

// Show creates and shows the dialog.
// This method call will block until the dialog is closed. Then the value can
// be read.
func (el *ColorDialog) Show() {
	tk.Get().Eval(
		"tk_chooseColor -parent %s -title {%s} -initialcolor {%s}",
		el.GetParent().GetID(),
		el.title,
		el.initialColor,
	)

	el.value = tk.Get().GetStrResult()
}

// GetValue gets the dialog value when closed.
func (el *ColorDialog) GetValue() string {
	return el.value
}
