package scale

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "scale"
)

// A scale widget is typically used to control the numeric value that varies
// uniformly over some range. A scale displays a slider that can be moved along
// over a trough, with the relative position of the slider over the trough
// indicating the value.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_scale.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Scale -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Scale -pkg=common/command
//go:generate go run ../../internal/tools/generate/main.go -recv=*Scale -pkg=common/floatvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Scale -pkg=common/length
//go:generate go run ../../internal/tools/generate/main.go -recv=*Scale -pkg=common/range
type Scale struct {
	widget.Widget

	valueVar string
}

// New creates a new progress bar.
// See [option.orientation] for orientation strings.
func New(parent element.Element, orientation string) *Scale {
	scale := &Scale{}
	scale.SetParent(parent)
	scale.SetType(Type)

	scale.valueVar = variable.GenerateName(scale.GetID())

	tk.Get().Eval("ttk::scale %s -orient {%s} -variable %s", scale.GetID(), orientation, scale.valueVar)

	return scale
}
