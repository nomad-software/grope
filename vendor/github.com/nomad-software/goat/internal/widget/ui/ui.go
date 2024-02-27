package ui

import (
	"strconv"
	"strings"

	"github.com/nomad-software/goat/internal/log"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

// Ele provides a base implementation of an ui element.
type Ele struct {
	element.Ele
}

// Update causes operations that are normally deferred, such as display updates
// and window layout calculations, to be performed immediately.
func (el *Ele) Update() {
	tk.Get().Eval("update idletasks")
}

// GetClass gets the ui element class.
// See [element.class] for class names.
func (el *Ele) GetClass() string {
	tk.Get().Eval("%s cget -class ", el.GetID())

	result := tk.Get().GetStrResult()

	if result == "" {
		tk.Get().Eval("winfo class %s", el.GetID())
		result = tk.Get().GetStrResult()
	}

	return result
}

// GetStyle gets the ui element class.
// See [element.style] for style names.
func (el *Ele) GetStyle() string {
	tk.Get().Eval("%s cget -style ", el.GetID())
	return tk.Get().GetStrResult()
}

// SetCursor sets the cursor of the ui element.
// See [option.cursor] for cursor names.
func (el *Ele) SetCursor(cursor string) {
	tk.Get().Eval("%s configure -cursor {%s}", el.GetID(), cursor)
}

// GetCursor gets the cursor of the ui element.
// See [option.cursor] for cursor names.
func (el *Ele) GetCursor() string {
	tk.Get().Eval("%s cget -cursor", el.GetID())
	return tk.Get().GetStrResult()
}

// SetKeyboadFocus sets that this ui element accepts the focus during keyboard
// traversal.
func (el *Ele) SetKeyboadFocus(focus bool) {
	tk.Get().Eval("%s configure -takefocus %v", el.GetID(), focus)
}

// AcceptsKeyboadFocus returns true if this ui element accepts the focus during
// keyboard traversal.
func (el *Ele) AcceptsKeyboadFocus() bool {
	tk.Get().Eval("%s cget -takefocus", el.GetID())
	return tk.Get().GetBoolResult()
}

// Destroy removes the ui element from the UI and cleans up its resources. Once
// destroyed you cannot refer to this ui element again or you will get a bad
// path name error from the interpreter.
func (el *Ele) Destroy() {
	tk.Get().Eval("destroy %s", el.GetID())
	el.SetType("destroyed")
}

// GetWidth gets the width of the ui element.
//
// Returns an int giving a ui element width in pixels. When a ui element is
// first created its width will be 1 pixel; the width will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (el *Ele) GetWidth() int {
	tk.Get().Eval("winfo width %s", el.GetID())
	return tk.Get().GetIntResult()
}

// GetHeight gets the height of the ui element.
//
// Returns an int giving a ui element height in pixels. When a ui element
// is first created its height will be 1 pixel; the height will eventually be
// changed by a geometry manager to fulfil the window's needs.
func (el *Ele) GetHeight() int {
	tk.Get().Eval("winfo height %s", el.GetID())
	return tk.Get().GetIntResult()
}

// GetOSHandle gets the OS specific window handle.
//
// Returns a low-level platform-specific identifier for a window. On Unix
// platforms, this is the X window identifier. Under Windows, this is the
// Windows HWND. On the Macintosh the value has no meaning outside Tk.
func (el *Ele) GetOSHandle() int64 {
	tk.Get().Eval("winfo id %s", el.GetID())
	result := tk.Get().GetStrResult()

	// Remove the 0x prefix.
	if len(result) > 2 {
		result = result[2:]
	}

	hwnd, err := strconv.ParseInt(result, 16, 0)
	if err != nil {
		log.Error(err)
	}

	return hwnd
}

// GetCursorPos gets the x and y position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns a
// list with two integers, which are the pointer's x and y coordinates measured
// in pixels in the screen's root window. If a virtual root window is in use on
// the screen, the position is computed in the virtual root. If the mouse
// pointer is not on the same screen as ui element then both of the returned
// coordinates are -1.
func (el *Ele) GetCursorPos() []int {
	tk.Get().Eval("winfo pointerxy %s", el.GetID())
	result := tk.Get().GetStrResult()

	strs := strings.Split(result, " ")
	pos := make([]int, 0)

	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Error(err)
		}
		pos = append(pos, i)
	}

	return pos
}

// GetCursorXPos gets the x position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns the
// pointer's x coordinate, measured in pixels in the screen's root window. If a
// virtual root window is in use on the screen, the position is measured in the
// virtual root. If the mouse pointer is not on the same screen as ui element
// then -1 is returned.
func (el *Ele) GetCursorXPos() int {
	tk.Get().Eval("winfo pointerx %s", el.GetID())
	return tk.Get().GetIntResult()
}

// GetCursorYPos gets the y position of the cursor on the ui element.
//
// If the mouse pointer is on the same screen as the ui element, it returns the
// pointer's y coordinate, measured in pixels in the screen's root window. If a
// virtual root window is in use on the screen, the position is measured in the
// virtual root. If the mouse pointer is not on the same screen as ui element
// then -1 is returned.
func (el *Ele) GetCursorYPos() int {
	tk.Get().Eval("winfo pointery %s", el.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenWidth gets the width of the screen this ui element is on.
func (el *Ele) GetScreenWidth() int {
	tk.Get().Eval("winfo screenwidth %s", el.GetID())
	return tk.Get().GetIntResult()
}

// GetScreenHeight gets the height of the screen this ui element is on.
func (el *Ele) GetScreenHeight() int {
	tk.Get().Eval("winfo screenheight %s", el.GetID())
	return tk.Get().GetIntResult()
}

// GetXPos gets the x position of the ui element.
// You may need to wait until the ui element has been updated for this to
// return the correct value.
func (el *Ele) GetXPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo x %s", el.GetID())
	} else {
		tk.Get().Eval("winfo rootx %s", el.GetID())
	}
	return tk.Get().GetIntResult()
}

// GetYPos gets the y position of the ui element.
// You may need to wait until the ui element has been updated for this to
// return the correct value.
func (el *Ele) GetYPos(relativeToParent bool) int {
	if relativeToParent {
		tk.Get().Eval("winfo y %s", el.GetID())
	} else {
		tk.Get().Eval("winfo rooty %s", el.GetID())
	}
	return tk.Get().GetIntResult()
}

// Focus gives focus to the ui element.
func (el *Ele) Focus(force bool) {
	if force {
		tk.Get().Eval("focus -force %s", el.GetID())
	} else {
		tk.Get().Eval("focus %s", el.GetID())
	}
}

// Lower lowers a ui element below another if specified or below all of its
// siblings in the stacking order
func (el *Ele) Lower(e element.Element) {
	if e != nil {
		tk.Get().Eval("lower %s %s", el.GetID(), e.GetID())
	} else {
		tk.Get().Eval("lower %s", el.GetID())
	}
}

// Raise raises a ui element above another if specified or above all of its
// siblings in the stacking order.
func (el *Ele) Raise(e element.Element) {
	if e != nil {
		tk.Get().Eval("raise %s %s", el.GetID(), e.GetID())
	} else {
		tk.Get().Eval("raise %s", el.GetID())
	}
}

// EnableGeometryAutoSize sets if the element should change it size when
// requested to do so by a geometry manager.
//
// The geometry manager normally computes how large a master must be to just
// exactly meet the needs of its slaves, and it sets the requested width and
// height of the master to these dimensions. This causes geometry information
// to propagate up through a window hierarchy to a top-level window so that the
// entire sub-tree sizes itself to fit the needs of the leaf windows. However,
// this command may be used to turn off propagation for one or more masters. If
// propagation is disabled then it will not set the requested width and height
// of the master window. This may be useful if, for example, you wish for a
// master window to have a fixed size that you specify.
func (el *Ele) EnableGeometryAutoSize(enable bool) {
	tk.Get().Eval("pack propagate %s %v", el.GetID(), enable)
	tk.Get().Eval("grid propagate %s %v", el.GetID(), enable)
}

// SetGridColumnWeight is used by the grid geometry manager to configure column weights.
func (el *Ele) SetGridColumnWeight(column, weight int) {
	tk.Get().Eval("grid columnconfigure %s %d -weight %d", el.GetID(), column, weight)
}

// SetGridRowWeight is used by the grid geometry manager to configure row weights.
func (el *Ele) SetGridRowWeight(row, weight int) {
	tk.Get().Eval("grid rowconfigure %s %d -weight %d", el.GetID(), row, weight)
}
