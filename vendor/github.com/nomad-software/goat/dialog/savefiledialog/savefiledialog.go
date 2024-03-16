package savefiledialog

import (
	"fmt"
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "savefiledialog"
)

// SaveFileDialog is a dialog box used to open a file.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/getOpenFile.html
type SaveFileDialog struct {
	element.Ele

	title            string
	confirmOverwrite bool
	defaultExt       string
	fileTypes        []string
	initialDir       string
	initialFile      string
	typeVar          string

	value string
}

// New creates a new save file dialog.
func New(parent element.Element, title string) *SaveFileDialog {
	dialog := &SaveFileDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)
	dialog.SetConfirmOverwrite(true)
	dialog.AddFileType("All files", "*")

	dialog.typeVar = variable.GenerateName(dialog.GetType(), dialog.GetID())

	return dialog
}

// SetTitle sets the dialog title.
func (el *SaveFileDialog) SetTitle(title string) {
	el.title = title
}

// SetConfirmOverwrite sets file overwrite confirmation.
func (el *SaveFileDialog) SetConfirmOverwrite(overwrite bool) {
	el.confirmOverwrite = overwrite
}

// AddFileType adds a file type to the dialog.
//
// On the Unix and Macintosh platforms, extensions are matched using glob-style
// pattern matching. On the Windows platform, extensions are matched by the
// underlying operating system.
//
// The types of possible extensions are.
// The special extension "*" matches any file.
// The special extension "" matches any files that do not have an extension
// (i.e., the filename contains no full stop character).
// Any character string that does not contain any wild card characters (* and
// ?).
//
// Due to the different pattern matching rules on the various platforms, to
// ensure portability, wild card characters are not allowed in the extensions,
// except as in the special extension "*". Extensions without a full stop
// character (e.g. "~") are allowed but may not work on all platforms.
func (el *SaveFileDialog) AddFileType(name, ext string) {
	el.fileTypes = append(el.fileTypes, fmt.Sprintf("{{%s} {%s}}", name, ext))
}

// SetInitialDirectory sets the initial directory of the dialog.
func (el *SaveFileDialog) SetInitialDirectory(dir string) {
	el.initialDir = dir
}

// SetInitialFile sets the initial file of the dialog.
func (el *SaveFileDialog) SetInitialFile(file string) {
	el.initialFile = file
}

// Show creates and shows the dialog.
// This method call will block until the dialog is closed. Then the value can
// be read.
func (el *SaveFileDialog) Show() {
	tk.Get().Eval(
		"tk_getSaveFile -parent %s -title {%s} -confirmoverwrite %v -defaultextension {%s} -filetypes {%s} -initialdir {%s} -initialfile {%s} -typevariable %s",
		el.GetParent().GetID(),
		el.title,
		el.confirmOverwrite,
		el.defaultExt,
		strings.Join(el.fileTypes, " "),
		el.initialDir,
		el.initialFile,
		el.typeVar,
	)

	el.value = tk.Get().GetStrResult()
}

// GetValues gets the dialog value when closed.
func (el *SaveFileDialog) GetValue() string {
	return el.value
}
