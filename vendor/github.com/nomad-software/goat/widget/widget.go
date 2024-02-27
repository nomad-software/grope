package widget

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui"
	"github.com/nomad-software/goat/option/state"
)

// Widget defines a widget at the lowest level.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_widget.html
type Widget struct {
	ui.Ele
}

// SetState sets the state of the widget.
// See [option.state] for state values.
func (w *Widget) SetState(state []string) {
	tk.Get().Eval("%s state {%s}", w.GetID(), strings.Join(state, " "))
}

// RemoveState removes states from the widget.
// See [option.state] for state values.
func (w *Widget) RemoveState(state []string) {
	tk.Get().Eval("%s state {!%s}", w.GetID(), strings.Join(state, " !"))
}

// InState returns true if the widget is in the passed state.
// See [option.state] for state values.
func (w *Widget) InState(state string) bool {
	tk.Get().Eval("%s instate {%s}", w.GetID(), state)
	return tk.Get().GetBoolResult()
}

// GetState gets the state of the widget.
// See [option.state] for state values.
func (w *Widget) GetState() []string {
	tk.Get().Eval("%s state", w.GetID())
	result := tk.Get().GetStrResult()
	return strings.Split(result, " ")
}

// ResetState resets all state to default.
func (w *Widget) ResetState() {
	w.RemoveState(w.GetState())
}

// Enable enables the widget.
// See [option.state.Disabled]
func (w *Widget) Enable() {
	w.RemoveState([]string{state.Disabled})
}

// Disable disables the widget.
// See [option.state.Disabled]
func (w *Widget) Disable() {
	w.SetState([]string{state.Disabled})
}

// Focus sets that the widget has keyboard focus.
// See [option.state.Focus]
func (w *Widget) Focus() {
	w.SetState([]string{state.Focus})
}

// Blur sets that the widget does not have keyboard focus.
// See [option.state.Focus]
func (w *Widget) Blur() {
	w.RemoveState([]string{state.Focus})
}

// Press set the widget to be pressed.
// See [option.state.Pressed]
func (w *Widget) Press() {
	w.SetState([]string{state.Pressed})
}

// UnPress set the widget to not be pressed.
// See [option.state.Pressed]
func (w *Widget) UnPress() {
	w.RemoveState([]string{state.Pressed})
}

// Select sets the widget to be selected.
// See [option.state.Selected]
func (w *Widget) Select() {
	w.SetState([]string{state.Selected})
}

// Deselect sets the widget to not be selected.
// See [option.state.Selected]
func (w *Widget) Deselect() {
	w.RemoveState([]string{state.Selected})
}

// ReadOnly sets the widget to be readonly.
// See [option.state.Readonly]
func (w *Widget) ReadOnly() {
	w.SetState([]string{state.Readonly})
}

// ReadAnWrite sets the widget to be read and write.
// See [option.state.Readonly]
func (w *Widget) ReadAndWrite() {
	w.RemoveState([]string{state.Readonly})
}

// Invalid sets the widget to be invalid.
// See [option.state.Invalid]
func (w *Widget) Invalid() {
	w.SetState([]string{state.Invalid})
}

// Valid sets the widget to be valid.
// See [option.state.Invalid]
func (w *Widget) Valid() {
	w.RemoveState([]string{state.Invalid})
}

// Pack uses a geometry method for loosely placing this widget inside its
// parent using a web browser model. Widgets flow around each other in the
// available space.
//
// outerPad = The amound of padding to add around the widget.
// innerPad = The amound of padding to add inside the widget.
// side     = The side to place the widget inside its parent.
// fill     = The space to fill inside its parent.
// anchor   = The anchor position of the widget inside its parent.
// expand   = Whether or not to expand to fill the entire given space.
//
// See [option.side] for side values.
// See [option.fill] for fill values.
// See [option.anchor] for anchor values.
func (w *Widget) Pack(outerPad, innerPad int, side, fill, anchor string, expand bool) {
	tk.Get().Eval("pack %s -padx %d -pady %d -ipadx %d -ipady %d -side {%s} -fill {%s} -anchor {%s} -expand %v", w.GetID(), outerPad, outerPad, innerPad, innerPad, side, fill, anchor, expand)
}

// Grid uses a geometry method for placing this widget inside its parent using
// an imaginary grid. Somewhat more direct and intuitive than pack. Choose grid
// for tabular layouts, and when there's no good reason to choose something
// else.
//
// If a widget's cell is larger than its default dimensions, the sticky
// parameter may be used to position (or stretch) the widget within its
// cell. The sticky argument is a string that contains zero or more of the
// characters n, s, e or w. Each letter refers to a side (north, south,
// east, or west) that the widget will 'stick' to. If both n and s (or e
// and w) are specified, the widget will be stretched to fill the entire
// height (or width) of its cell. The sticky parameter subsumes the
// combination of anchor and fill that is used by pack. The default is an
// empty string, which causes the widget to be centered in its cell, at its
// default size.
//
// column   = The column in which to place this widget.
// row      = The row in which to place this widget.
// outerPad = The amound of padding to add around the widget.
// innerPad = The amound of padding to add inside the widget.
// colSpan  = The amount of column this widget should span across.
// rowSpan  = The amount of rows this widget should span across.
// sticky   = Which edges of the cell the widget should touch. See note above.
func (w *Widget) Grid(column, row, outerPad, innerPad, colspan, rowspan int, sticky string) {
	tk.Get().Eval("grid %s -column %d -row %d -padx %d -pady %d -ipadx %d -ipady %d -columnspan %d -rowspan %d -sticky {%s}", w.GetID(), column, row, outerPad, outerPad, innerPad, innerPad, colspan, rowspan, sticky)
}

// Place uses a geometry method for placing this widget inside its parent using
// absolute positioning.
//
// x      = The horizontal position of the widget inside its parent.
// y      = The vertical position of the widget inside its parent.
// width  = The width of the widget.
// height = The height of the widget.
// anchor = The anchor position of the widget inside its parent.
// border = How the widget interacts with the parent's border.
//
// See [option.anchor] for anchor values.
// See [option.bordermode] for border mode values.
func (w *Widget) Place(x, y, width, height int, anchor string, border string) {
	tk.Get().Eval("place %s -x %d -y %d -width %d -height %d -anchor {%s} -bordermode {%s}", w.GetID(), x, y, width, height, anchor, border)
}

// PlaceRelative uses a geometry method for placing this widget inside its
// parent using relative positioning. In this case the position and size is
// specified as a floating-point number between 0.0 and 1.0 relative to the
// height of the parent. 0.5 means the widget will be half as high as the
// parent and 1.0 means the widget will have the same height as the parent, and
// so on.
//
// x = The relative horizontal position of the widget inside its parent.
// y = The relative vertical position of the widget inside its parent.
// width = The relative width of the widget.
// height = The relative height of the widget.
// anchor = The anchor position of the widget inside its parent.
// border = How the widget interacts with the parent's border.
//
// See [option.anchor] for anchor values.
// See [option.bordermode] for border mode values.
func (w *Widget) PlaceRelative(x, y, width, height float64, anchor string, border string) {
	tk.Get().Eval("place %s -relx %v -rely %v -relwidth %v -relheight %v -anchor {%s} -bordermode {%s}", w.GetID(), x, y, width, height, anchor, border)
}
