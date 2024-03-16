package progressbar

import (
	"time"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "progressbar"
)

// A progress bar widget shows the status of a long-running operation. They can
// operate in two modes: determinate mode shows the amount completed relative
// to the total amount of work to be done, and indeterminate mode provides an
// animated display to let the user know that something is happening.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_progressbar.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*ProgressBar -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*ProgressBar -pkg=common/floatvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*ProgressBar -pkg=common/length
type ProgressBar struct {
	widget.Widget

	valueVar string
}

// New creates a new progress bar.
// See [option.orientation] for orientation strings.
func New(parent element.Element, orientation string) *ProgressBar {
	bar := &ProgressBar{}
	bar.SetParent(parent)
	bar.SetType(Type)

	bar.valueVar = variable.GenerateName(bar.GetID())

	tk.Get().Eval("ttk::progressbar %s -orient {%s} -variable %s", bar.GetID(), orientation, bar.valueVar)

	return bar
}

// SetMode sets the mode of the progress bar.
// See [option.progressmode] for mode strings.
func (el *ProgressBar) SetMode(mode string) {
	tk.Get().Eval("%s configure -mode {%s}", el.GetID(), mode)
}

// SetMaxValue sets the maximum value.
func (el *ProgressBar) SetMaxValue(value float64) {
	tk.Get().Eval("%s configure -maximum %v", el.GetID(), value)
}

// GetPhase gets the current phase.
// The widget periodically increments the value of this option whenever the
// value is greater than 0 and, in determinate mode, less than maximum. This
// option may be used by the current theme to provide additional animation
// effects.
func (el *ProgressBar) GetPhase() float64 {
	tk.Get().Eval("%s cget -phase", el.GetID())
	return tk.Get().GetFloatResult()
}

// Increment increments the value by the specified amount.
func (el *ProgressBar) Increment(increment float64) {
	tk.Get().Eval("%s step %v", el.GetID(), increment)
}

// StartAutoIncrement starts the auto-incrementing of the progress bar.
func (el *ProgressBar) StartAutoIncrement(d time.Duration) {
	tk.Get().Eval("%s start %d", el.GetID(), d/time.Millisecond)
}

// StopAutoIncrement stops the auto-incrementing of the progress bar.
func (el *ProgressBar) StopAutoIncrement(d time.Duration) {
	tk.Get().Eval("%s stop", el.GetID())
}
