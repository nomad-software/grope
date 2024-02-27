package line

import (
	"fmt"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvasline"
)

// Line represents a line in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/dash
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/scale
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/width
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Line -pkg=canvas/zorder
type Line struct {
	element.Ele
}

// Creates a new line.
func New(parent element.Element) *Line {
	line := &Line{}
	line.SetParent(parent)
	line.SetType(Type)

	return line
}

// SetArrow sets the arrows.
// See [widget.canvas.line.arrow] for arrow options.
func (el *Line) SetArrow(arrow string) {
	tk.Get().Eval("%s itemconfigure %s -arrow {%s}", el.GetParent().GetID(), el.GetID(), arrow)
}

// SetArrowShape sets the arrow shape.
// The first element of the list gives the distance along the line from the
// neck of the arrowhead to its tip. The second element gives the distance
// along the line from the trailing points of the arrowhead to the tip, and the
// third element gives the distance from the outside edge of the line to the
// trailing points. If this option is not specified then Tk picks a
// “reasonable” shape.
func (el *Line) SetArrowShape(base, delta, width float64) {
	tk.Get().Eval("%s itemconfigure %s -arrowshape [list %v %v %v]", el.GetParent().GetID(), el.GetID(), base, delta, width)
}

// SetCapStyle specifies the ways in which caps are to be drawn at the
// endpoints of the line. If this option is not specified then it defaults to
// butt. Where arrowheads are drawn the cap style is ignored.
// See [widget.canvas.line.capstyle] for cap style options.
func (el *Line) SetCapStyle(style string) {
	tk.Get().Eval("%s itemconfigure %s -capstyle {%s}", el.GetParent().GetID(), el.GetID(), style)
}

// SetJoinStyle specifies the ways in which joints are to be drawn at the
// vertices of the line. If this option is not specified then it defaults to
// round. If the line only contains two points then this option is irrelevant.
// See [widget.canvas.line.joinstyle] for join style options.
func (el *Line) SetJoinStyle(style string) {
	tk.Get().Eval("%s itemconfigure %s -joinstyle {%s}", el.GetParent().GetID(), el.GetID(), style)
}

// SetSmoothMethod sets the smooth method.
// If the smoothing method is bezier, this indicates that the line should be
// drawn as a curve, rendered as a set of quadratic splines: one spline is
// drawn for the first and second line segments, one for the second and third,
// and so on. Straight-line segments can be generated within a curve by
// duplicating the end-points of the desired line segment. If the smoothing
// method is raw, this indicates that the line should also be drawn as a curve
// but where the list of coordinates is such that the first coordinate pair
// (and every third coordinate pair thereafter) is a knot point on a cubic
// Bezier curve, and the other coordinates are control points on the cubic
// Bezier curve. Straight line segments can be generated within a curve by
// making control points equal to their neighbouring knot points. If the last
// point is a control point and not a knot point, the point is repeated (one or
// two times) so that it also becomes a knot point.
// See [widget.canvas.line.smoothmethod] for join style options.
func (el *Line) SetSmoothMethod(method string) {
	tk.Get().Eval("%s itemconfigure %s -smooth {%s}", el.GetParent().GetID(), el.GetID(), method)
}

// SetSplineSteps specifies the degree of smoothness desired for curves: each
// spline will be approximated with number line segments. This option is
// ignored unless the smooth method is set.
func (el *Line) SetSplineSteps(n int) {
	tk.Get().Eval("%s itemconfigure %s -splinesteps %d", el.GetParent().GetID(), el.GetID(), n)
}

// SetCoords updates the item coordinates.
func (el *Line) SetCoords(x1, y1, x2, y2 float64, others ...float64) {
	otherStr := ""
	for _, i := range others {
		otherStr += fmt.Sprintf(" %v", i)
	}
	tk.Get().Eval("%s coords %s [list %v %v %v %v %s]", el.GetParent().GetID(), el.GetID(), x1, y1, x2, y2, otherStr)
}
