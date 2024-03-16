package labelframe

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "labelframe"
)

// LabelFrame is a container used to group other widgets together. It has an
// optional label, which may be a plain text string or another widget.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_labelframe.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=common/padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=common/text
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=common/underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*LabelFrame -pkg=common/width
type LabelFrame struct {
	widget.Widget
}

// New creates a new label frame.
func New(parent element.Element, text string, underline int) *LabelFrame {
	frame := &LabelFrame{}
	frame.SetParent(parent)
	frame.SetType(Type)

	tk.Get().Eval("ttk::labelframe %s -text {%s} -underline %d", frame.GetID(), text, underline)

	return frame
}

// SetLabelAnchor sets the anchor for the label.
// See [option.anchor] for anchor values.
func (l *LabelFrame) SetLabelAnchor(anchor string) {
	tk.Get().Eval("%s configure -labelanchor {%s}", l.GetID(), anchor)
}
