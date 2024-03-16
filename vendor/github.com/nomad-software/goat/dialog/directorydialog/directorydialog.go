package directorydialog

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "directorydialog"
)

// DirectoryDialog is a dialog box used to choose a directory.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/chooseDirectory.html
type DirectoryDialog struct {
	element.Ele

	title        string
	initialDir   string
	dirMustExist bool

	value string
}

// New creates a new directory dialog.
func New(parent element.Element, title string) *DirectoryDialog {
	dialog := &DirectoryDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)
	dialog.SetDirectoryMustExist(true)

	return dialog
}

// SetTitle sets the dialog title.
func (el *DirectoryDialog) SetTitle(title string) {
	el.title = title
}

// SetInitialDirectory sets the initial directory of the dialog.
func (el *DirectoryDialog) SetInitialDirectory(dir string) {
	el.initialDir = dir
}

// SetDirectoryMustExist sets that the choosen directory must exist.
func (el *DirectoryDialog) SetDirectoryMustExist(exist bool) {
	el.dirMustExist = exist
}

// Show creates and shows the dialog.
// This method call will block until the dialog is closed. Then the value can
// be read.
func (el *DirectoryDialog) Show() {
	tk.Get().Eval(
		"tk_chooseDirectory -parent %s -title {%s} -initialdir {%s} -mustexist %v",
		el.GetParent().GetID(),
		el.title,
		el.initialDir,
		el.dirMustExist,
	)

	el.value = tk.Get().GetStrResult()
}

// GetValue gets the dialog value when closed.
func (el *DirectoryDialog) GetValue() string {
	return el.value
}
