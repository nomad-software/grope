package openfiledialog

import (
	"fmt"
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "openfiledialog"
)

// OpenFileDialog is a dialog box used to open a file.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/getOpenFile.html
type OpenFileDialog struct {
	element.Ele

	title       string
	defaultExt  string
	fileTypes   []string
	initialDir  string
	initialFile string
	multiple    bool
	typeVar     string

	values []string
}

// New creates a new open file dialog.
func New(parent element.Element, title string) *OpenFileDialog {
	dialog := &OpenFileDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)
	dialog.EnableMultipleSelection(false)
	dialog.AddFileType("All files", "*")

	dialog.typeVar = variable.GenerateName(dialog.GetType(), dialog.GetID())

	return dialog
}

// SetTitle sets the dialog title.
func (el *OpenFileDialog) SetTitle(title string) {
	el.title = title
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
func (el *OpenFileDialog) AddFileType(name, ext string) {
	el.fileTypes = append(el.fileTypes, fmt.Sprintf("{{%s} {%s}}", name, ext))
}

// SetInitialDirectory sets the initial directory of the dialog.
func (el *OpenFileDialog) SetInitialDirectory(dir string) {
	el.initialDir = dir
}

// SetInitialFile sets the initial file of the dialog.
func (el *OpenFileDialog) SetInitialFile(file string) {
	el.initialFile = file
}

// EnableMultipleSelection enables multiple file selection.
func (el *OpenFileDialog) EnableMultipleSelection(multiple bool) {
	el.multiple = multiple
}

// Show creates and shows the dialog.
// This method call will block until the dialog is closed. Then the value can
// be read.
func (el *OpenFileDialog) Show() {
	tk.Get().Eval(
		"tk_getOpenFile -parent %s -title {%s} -defaultextension {%s} -filetypes {%s} -initialdir {%s} -initialfile {%s} -multiple %v -typevariable %s",
		el.GetParent().GetID(),
		el.title,
		el.defaultExt,
		strings.Join(el.fileTypes, " "),
		el.initialDir,
		el.initialFile,
		el.multiple,
		el.typeVar,
	)

	el.values = tk.Get().GetStrSliceResult()
}

// GetValue gets the dialog value when closed.
func (el *OpenFileDialog) GetValue() string {
	if len(el.values) > 0 {
		return el.values[0]
	}

	return ""
}

// GetValues gets the dialog values when closed.
func (el *OpenFileDialog) GetValues() []string {
	return el.values
}
