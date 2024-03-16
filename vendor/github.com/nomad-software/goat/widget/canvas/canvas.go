package canvas

import (
	"fmt"

	img "github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/relief"
	wid "github.com/nomad-software/goat/widget"
	"github.com/nomad-software/goat/widget/canvas/arc"
	"github.com/nomad-software/goat/widget/canvas/arc/style"
	"github.com/nomad-software/goat/widget/canvas/image"
	"github.com/nomad-software/goat/widget/canvas/line"
	"github.com/nomad-software/goat/widget/canvas/oval"
	"github.com/nomad-software/goat/widget/canvas/polygon"
	"github.com/nomad-software/goat/widget/canvas/rectangle"
	"github.com/nomad-software/goat/widget/canvas/tag"
	"github.com/nomad-software/goat/widget/canvas/text"
	"github.com/nomad-software/goat/widget/canvas/widget"
)

const (
	Type = "canvas"
)

// Canvas widgets implement structured graphics. A canvas displays any number
// of items, which may be things like rectangles, circles, lines, and text.
// Items may be manipulated (e.g. moved or re-colored) and commands may be
// associated with items in much the same way that the bind command allows
// commands to be bound to widgets.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/canvas.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/color
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/relief
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/scrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Canvas -pkg=common/width
type Canvas struct {
	wid.Widget

	itemRef map[string]element.Element
	items   []element.Element
}

// New creates a new canvas.
func New(parent element.Element) *Canvas {
	canvas := &Canvas{
		itemRef: make(map[string]element.Element),
		items:   make([]element.Element, 0),
	}
	canvas.SetParent(parent)
	canvas.SetType(Type)

	tk.Get().Eval("canvas %s", canvas.GetID())

	canvas.SetBorderWidth(1)
	canvas.SetRelief(relief.Sunken)

	return canvas
}

// AddArc adds an arc to the canvas.
// The first four coordinates specify the oval that this arc is drawn on.
func (el *Canvas) AddArc(x1, y1, x2, y2 float64) *arc.Arc {
	tk.Get().Eval("%s create arc [list %v %v %v %v]", el.GetID(), x1, y1, x2, y2)
	id := tk.Get().GetStrResult()

	item := arc.New(el)
	item.SetID(id)
	item.SetStyle(style.Pie)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddImage adds an image to the canvas.
func (el *Canvas) AddImage(img *img.Image, x, y float64) *image.Image {
	tk.Get().Eval("%s create image [list %v %v] -image %s", el.GetID(), x, y, img.GetID())
	id := tk.Get().GetStrResult()

	item := image.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddLine adds a line to the canvas.
// The arguments give the coordinates for a series of two or more points that
// describe a series of connected line segments.
func (el *Canvas) AddLine(x1, y1, x2, y2 float64, others ...float64) *line.Line {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s create line [list %v %v %v %v %s]", el.GetID(), x1, y1, x2, y2, otherStr)
	id := tk.Get().GetStrResult()

	item := line.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddArc adds an oval to the canvas.
// The four coordinates specify the oval.
func (el *Canvas) AddOval(x1, y1, x2, y2 float64) *oval.Oval {
	tk.Get().Eval("%s create oval [list %v %v %v %v]", el.GetID(), x1, y1, x2, y2)
	id := tk.Get().GetStrResult()

	item := oval.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddPolygon adds a polygon to the canvas.
// The arguments give the coordinates for a series of three or more points that
// describe a series of connected polygon vertices.
func (el *Canvas) AddPolygon(x1, y1, x2, y2, x3, y3 float64, others ...float64) *polygon.Polygon {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s create polygon [list %v %v %v %v %v %v %s]", el.GetID(), x1, y1, x2, y2, x3, y3, otherStr)
	id := tk.Get().GetStrResult()

	item := polygon.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddRectangle adds a rectangle to the canvas.
// The four coordinates specify the rectangle.
func (el *Canvas) AddRectangle(x1, y1, x2, y2 float64) *rectangle.Rectangle {
	tk.Get().Eval("%s create rectangle [list %v %v %v %v]", el.GetID(), x1, y1, x2, y2)
	id := tk.Get().GetStrResult()

	item := rectangle.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddText adds text to the canvas.
func (el *Canvas) AddText(txt string, x, y float64) *text.Text {
	tk.Get().Eval("%s create text [list %v %v] -text {%s}", el.GetID(), x, y, txt)
	id := tk.Get().GetStrResult()

	item := text.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// AddWidget adds a widget to the canvas.
func (el *Canvas) AddWidget(e element.Element, x, y float64) *widget.Widget {
	tk.Get().Eval("%s create window [list %v %v] -window %s", el.GetID(), x, y, e.GetID())
	id := tk.Get().GetStrResult()

	item := widget.New(el)
	item.SetID(id)

	el.itemRef[id] = item
	el.items = append(el.items, item)

	return item
}

// GetTag gets a tag from the canvas in order to modify its properties.
// Tags exist once they've been added to a canvas item.
func (el *Canvas) GetTag(name string) *tag.Tag {
	t := tag.New(el)
	t.SetID(name)

	return t
}

// SetSelectionTolerance sets the selection tolerance.
// Specifies a floating-point value indicating how close the mouse cursor must
// be to an item before it is considered to be “inside” the item. Defaults to
// 1.0.
func (el *Canvas) SetSelectionTolerance(tolerance float64) {
	tk.Get().Eval("%s configure -closeenough %v", el.GetID(), tolerance)
}

// SetConfineScrollRegion sets if the scroll region should be confined.
func (el *Canvas) SetConfineScrollRegion(confine bool) {
	tk.Get().Eval("%s configure -confine %v", el.GetID(), confine)
}

// SetScrollRegion sets the scroll region.
// Specifies a list with four coordinates describing the left, top, right, and
// bottom coordinates of a rectangular region. This region is used for
// scrolling purposes and is considered to be the boundary of the information
// in the canvas.
func (el *Canvas) SetScrollRegion(left, top, right, bottom float64) {
	tk.Get().Eval("%s configure -scrollregion [list %v %v %v %v]", el.GetID(), left, top, right, bottom)
}

// SetScrollStep sets the scroll step which specifies an increment for scrolling.
func (el *Canvas) SetScrollStep(step float64) {
	tk.Get().Eval("%s configure -xscrollincrement %v -yscrollincrement %v", el.GetID(), step, step)
}

// GetXPosFromElementXPos gets the x canvas position from the screen position.
func (el *Canvas) GetXPosFromElementXPos(x, grid int) int {
	tk.Get().Eval("%s canvasx %d %d", el.GetID(), x, grid)
	xPos := tk.Get().GetFloatResult()

	return int(xPos)
}

// GetYPosFromElementYPos gets the y canvas position from the screen position.
func (el *Canvas) GetYPosFromElementYPos(y, grid int) int {
	tk.Get().Eval("%s canvasy %d %d", el.GetID(), y, grid)
	yPos := tk.Get().GetFloatResult()

	return int(yPos)
}

// SetScanMark sets the scan mark.
// Records x and y and the canvas's current view; used in conjunction with
// later scan dragto commands. Typically this command is associated with a
// mouse button press in the widget and x and y are the coordinates of the
// mouse.
func (el *Canvas) SetScanMark(x, y int) {
	tk.Get().Eval("%s scan mark %d %d", el.GetID(), x, y)
}

// ScanDragTo computes the difference between its x and y arguments (which are
// typically mouse coordinates) and the x and y arguments to the last scan mark
// command for the widget. It then adjusts the view by gain times the
// difference in coordinates, where gain defaults to 10. This command is
// typically associated with mouse motion events in the widget, to produce the
// effect of dragging the canvas at high speed through its window. The return
// value is an empty string.
func (el *Canvas) ScanDragTo(x, y, gain int) {
	tk.Get().Eval("%s scan dragto %d %d %d", el.GetID(), x, y, gain)
}

// GetItems returns all items in the canvas.
func (el *Canvas) GetItems() []element.Element {
	return el.items
}

// GetItemNear gets the nearest item to the coordinates supplied and returns it.
func (el *Canvas) GetItemNear(x, y, radius int) element.Element {
	tk.Get().Eval("%s find closest %d %d %d", el.GetID(), x, y, radius)
	id := tk.Get().GetStrResult()

	if e, ok := el.itemRef[id]; ok {
		return e
	}

	return nil
}

// GetItemsEnclosed gets all the items fully enclosed by the passed rectangular
// region.
func (el *Canvas) GetItemsEnclosed(x1, y1, x2, y2 int) []element.Element {
	tk.Get().Eval("%s find enclosed %d %d %d %d", el.GetID(), x1, y1, x2, y2)
	ids := tk.Get().GetStrSliceResult()

	items := make([]element.Element, 0)

	for _, id := range ids {
		if e, ok := el.itemRef[id]; ok {
			items = append(items, e)
		}
	}

	return items
}

// GetItemsEnclosed gets all the items overlapping the passed rectangular
// region.
func (el *Canvas) GetItemsOverlapping(x1, y1, x2, y2 int) []element.Element {
	tk.Get().Eval("%s find overlapping %d %d %d %d", el.GetID(), x1, y1, x2, y2)
	ids := tk.Get().GetStrSliceResult()

	items := make([]element.Element, 0)

	for _, id := range ids {
		if e, ok := el.itemRef[id]; ok {
			items = append(items, e)
		}
	}

	return items
}

// Clear deletes all items from the canvas.
func (el *Canvas) Clear() {
	for _, item := range el.items {
		tk.Get().Eval("%s delete %s", el.GetID(), item.GetID())
	}
	clear(el.itemRef)
	clear(el.items)
}
