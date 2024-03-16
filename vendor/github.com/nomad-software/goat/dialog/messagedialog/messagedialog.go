package messagedialog

import (
	"github.com/nomad-software/goat/dialog/button"
	"github.com/nomad-software/goat/dialog/icon"
	dtype "github.com/nomad-software/goat/dialog/type"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "messagedialog"
)

// MessageDialog is a dialog box with a user defined message and buttons.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/messageBox.html
type MessageDialog struct {
	element.Ele

	title         string
	message       string
	detail        string
	dialogType    string
	icon          string
	defaultButton string

	value string
}

// New creates a new message dialog.
func New(parent element.Element, title string) *MessageDialog {
	dialog := &MessageDialog{}

	dialog.SetParent(parent)
	dialog.SetType(Type)

	dialog.SetTitle(title)
	dialog.SetDialogType(dtype.Ok)
	dialog.SetIcon(icon.Info)
	dialog.SetDefaultButton(button.Ok)

	return dialog
}

// SetTitle sets the dialog title.
func (el *MessageDialog) SetTitle(title string) {
	el.title = title
}

// SetMessage sets the dialog message.
func (el *MessageDialog) SetMessage(msg string) {
	el.message = msg
}

// SetDetail sets the dialog detail.
func (el *MessageDialog) SetDetail(detail string) {
	el.detail = detail
}

// SetDialogType sets the dialog type.
// See [dialog.dialogtype] for dialog type values.
func (el *MessageDialog) SetDialogType(typ string) {
	switch typ {
	case dtype.AbortRetryIgnore:
		el.defaultButton = button.Abort

	case dtype.RetryCancel:
		el.defaultButton = button.Retry

	case dtype.YesNo:
		el.defaultButton = button.Yes

	case dtype.YesNoCancel:
		el.defaultButton = button.Yes

	default:
		el.defaultButton = button.Ok
	}

	el.dialogType = typ
}

// SetIcon sets the dialog icon.
// See [dialog.dialogicon] for icon values.
func (el *MessageDialog) SetIcon(icon string) {
	el.icon = icon
}

// SetDefaultButton sets the dialog default button.
// See [dialog.dialogbutton] for button values.
func (el *MessageDialog) SetDefaultButton(button string) {
	el.defaultButton = button
}

// Show creates and shows the dialog.
// This method call will block until the dialog is closed. Then the value can
// be read.
func (el *MessageDialog) Show() {
	tk.Get().Eval(
		"tk_messageBox -parent %s -title {%s} -message {%s} -detail {%s} -type {%s} -icon {%s} -default {%s}",
		el.GetParent().GetID(),
		el.title,
		el.message,
		el.detail,
		el.dialogType,
		el.icon,
		el.defaultButton,
	)

	el.value = tk.Get().GetStrResult()
}

// GetValue gets the dialog value when closed.
func (el *MessageDialog) GetValue() string {
	return el.value
}
