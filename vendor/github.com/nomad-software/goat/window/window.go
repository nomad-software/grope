package window

import (
	"strconv"
	"strings"

	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/log"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

// Window is the struct representing a window.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/toplevel.html
//
//go:generate go run ../internal/tools/generate/main.go -recv=*Window -pkg=common/bind
//go:generate go run ../internal/tools/generate/main.go -recv=*Window -pkg=common/borderwidth
//go:generate go run ../internal/tools/generate/main.go -recv=*Window -pkg=common/color -methods=SetBackgroundColor
//go:generate go run ../internal/tools/generate/main.go -recv=*Window -pkg=common/height
//go:generate go run ../internal/tools/generate/main.go -recv=*Window -pkg=common/relief
//go:generate go run ../internal/tools/generate/main.go -recv=*Window -pkg=common/width
type Window struct {
	ui.Ele
}

// New creates a new window.
// The parent will usually be another window.
func New(parent element.Element) *Window {
	win := &Window{}
	win.SetParent(parent)
	win.SetType("window")

	// Create and show the window.
	tk.Get().Eval("toplevel %s", win.GetID())

	return win
}

// GetStyle gets the ui element style.
// Override and fake this for window because style is not supported.
// See [element.style] for style names.
func (w *Window) GetStyle() string {
	return "Toplevel"
}

// SetSize sets the window size.
func (w *Window) SetSize(width, height int) {
	tk.Get().Eval("wm geometry %s {%dx%d}", w.GetID(), width, height)
}

// SetGeometry sets the window size and position.
func (w *Window) SetGeometry(width, height, x, y int) {
	tk.Get().Eval("wm geometry %s {%dx%d+%d+%d}", w.GetID(), width, height, x, y)
}

// SetTitle sets the window title.
func (w *Window) SetTitle(title string) {
	tk.Get().Eval("wm title %s {%s}", w.GetID(), title)
}

// GetTitle gets the window title.
func (w *Window) GetTitle() string {
	tk.Get().Eval("wm title %s", w.GetID())
	return tk.Get().GetStrResult()
}

// WaitForVisibility waits until this window is visible.
// This is typically used to wait for a newly-created window to appear on
// the screen before taking some action.
func (w *Window) WaitForVisibility() {
	tk.Get().Eval("tkwait visibility %s", w.GetID())
}

// SetOpacity sets the window opacity if it's supported.
func (w *Window) SetOpacity(opacity float64) {
	tk.Get().Eval("wm attributes %s -alpha %v", w.GetID(), opacity)
}

// GetOpacity gets the window opacity if it's supported.
func (w *Window) GetOpacity() float64 {
	tk.Get().Eval("wm attributes %s -alpha", w.GetID())
	return tk.Get().GetFloatResult()
}

// SetFullScreen sets the window to be fullscreen or not.
func (w *Window) SetFullScreen(fullscreen bool) {
	tk.Get().Eval("wm attributes %s -fullscreen %v", w.GetID(), fullscreen)
}

// IsFullScreen gets if the window is fullscreen.
func (w *Window) IsFullScreen() bool {
	tk.Get().Eval("wm attributes %s -fullscreen", w.GetID())
	return tk.Get().GetBoolResult()
}

// SetTopmost sets the window to be the top-most. This makes the window not
// able to be lowered behind any others.
func (w *Window) SetTopmost(top bool) {
	tk.Get().Eval("wm attributes %s -topmost %v", w.GetID(), top)
}

// IsTopmost gets if the window is the top-most.
func (w *Window) IsTopmost() bool {
	tk.Get().Eval("wm attributes %s -topmost", w.GetID())
	return tk.Get().GetBoolResult()
}

// SetIconify sets whether the window is minimised.
func (w *Window) SetIconify(iconify bool) {
	if iconify {
		tk.Get().Eval("wm iconify %s", w.GetID())
	} else {
		tk.Get().Eval("wm deiconify %s", w.GetID())
	}
}

// SetMinSize sets the minimum size of the window.
func (w *Window) SetMinSize(width, height int) {
	tk.Get().Eval("wm minsize %s %d %d", w.GetID(), width, height)
}

// SetMinSize sets the maximum size of the window.
func (w *Window) SetMaxSize(width, height int) {
	tk.Get().Eval("wm maxsize %s %d %d", w.GetID(), width, height)
}

// SetProtocolCommand binds a callback to be called when the specified protocol
// is triggered. A window manager protocol is a class of messages sent from a
// window manager to a Tk application outside of the normal event processing
// system.
// See [window.protocol] for protocol names.
func (w *Window) SetProtocolCommand(protocol string, callback command.Callback) {
	name := command.GenerateName(protocol)
	tk.Get().CreateCommand(w, name, callback)
	tk.Get().Eval("wm protocol %s {%s} {%s}", w.GetID(), protocol, name)
}

// UnBindProtocol unbinds a previously bound callback.
func (w *Window) UnBindProtocol(protocol string) {
	name := command.GenerateName(protocol)
	tk.Get().Eval("wm protocol %s {%s} {}", w.GetID(), protocol)
	tk.Get().DestroyCommand(name)
}

// SetResizeable sets if a window width and height can be resized.
func (w *Window) SetResizeable(width, height bool) {
	tk.Get().Eval("wm resizable %s %v %v", w.GetID(), width, height)
}

// GetResizeable gets if a window width and height can be resized.
func (w *Window) GetResizeable() []bool {
	tk.Get().Eval("wm resizable %s", w.GetID())
	result := tk.Get().GetStrResult()

	strs := strings.Split(result, " ")
	res := make([]bool, 0)

	for _, s := range strs {
		i, err := strconv.ParseBool(s)
		if err != nil {
			log.Error(err)
		}
		res = append(res, i)
	}

	return res
}

// IsAbove returns if this window is above another.
func (w *Window) IsAbove(other *Window) bool {
	tk.Get().Eval("wm stackorder %s isabove %s", w.GetID(), other.GetID())
	return tk.Get().GetBoolResult()
}

// IsAbove returns if this window is above another.
func (w *Window) IsBelow(other *Window) bool {
	tk.Get().Eval("wm stackorder %s isbelow %s", w.GetID(), other.GetID())
	return tk.Get().GetBoolResult()
}

// Wait waits for the window to be destroyed.
// This is typically used to wait for a user to finish interacting with a
// dialog box before using the result of that interaction.
func (w *Window) Wait() {
	tk.Get().Eval("tkwait window %s", w.GetID())
}

// SetIcon sets the default icon for this window. This is applied to all future
// child windows as well.
//
// The data in the images is taken as a snapshot at the time of invocation. If
// the images are later changed, this is not reflected to the titlebar icons.
// Multiple images are accepted to allow different images sizes (e.g., 16x16
// and 32x32) to be provided. The window manager may scale provided icons to an
// appropriate size.
func (w *Window) SetIcon(imgs []*image.Image, applyToChildwindows bool) {
	ids := make([]string, 0)

	for _, img := range imgs {
		ids = append(ids, img.GetID())
	}

	if applyToChildwindows {
		tk.Get().Eval("wm iconphoto %s -default %s", w.GetID(), strings.Join(ids, " "))
	} else {
		tk.Get().Eval("wm iconphoto %s %s", w.GetID(), strings.Join(ids, " "))
	}
}

// SetPadding sets the padding.
func (w *Window) SetPadding(p int) {
	tk.Get().Eval("%s configure -padx %d -pady %d", w.GetID(), p, p)
}
