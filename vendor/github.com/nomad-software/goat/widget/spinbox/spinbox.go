package spinbox

import (
	"math"

	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/state"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "spinbox"
)

// A spinbox widget is an entry widget with built-in up and down buttons that
// are used to either modify a numeric value or to select among a set of
// values. The widget implements all the features of the entry widget.
//
// If a list of string values are set to be controlled by this widget it will
// override any numeric range or step set. The widget will instead use the
// values specified beginning with the first value.
//
// Virtual events that can also be bound to.
// <<Increment>>
// <<Decrement>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_spinbox.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/color -methods=SetForegroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/command
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/data
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/delete
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/font
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/justify
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/range
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/scrollbar -methods=AttachHorizontalScrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/selection
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/show
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/stringvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Spinbox -pkg=common/width
type Spinbox struct {
	widget.Widget

	valueVar string
}

// New creates a new spinbox.
func New(parent element.Element) *Spinbox {
	spinbox := &Spinbox{}
	spinbox.SetParent(parent)
	spinbox.SetType(Type)

	spinbox.valueVar = variable.GenerateName(spinbox.GetID())

	tk.Get().Eval("ttk::spinbox %s -textvariable %s", spinbox.GetID(), spinbox.valueVar)

	spinbox.SetState([]string{state.Readonly})
	spinbox.SetFromValue(math.MinInt8)
	spinbox.SetToValue(math.MaxInt8)
	spinbox.SetValue("0")

	spinbox.Bind("<<Increment>>", func(*command.BindData) {
		spinbox.DeselectText() // This doesn't seem to work.
	})

	spinbox.Bind("<<Decrement>>", func(*command.BindData) {
		spinbox.DeselectText() // This doesn't seem to work.
	})

	return spinbox
}

// SetStep sets the increment of each change.
func (el *Spinbox) SetStep(step float64) {
	tk.Get().Eval("%s configure -increment {%v}", el.GetID(), step)
}

// SetWrap sets if the value should wrap if it reaches the end.
func (el *Spinbox) SetWrap(wrap bool) {
	tk.Get().Eval("%s configure -wrap {%v}", el.GetID(), wrap)
}

// SetFormat sets the display format of the number.
// Before is the number of digits before a decimal place.
// After is the number of digits after a decimal place.
func (el *Spinbox) SetFormat(before, after int) {
	tk.Get().Eval("%s configure -format %%%d.%df", el.GetID(), before, after)
}
