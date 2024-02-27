package label

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "label"
)

// An label widget displays a one-line text string.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_label.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Label -pkg=common/image
//go:generate go run ../../internal/tools/generate/main.go -recv=*Label -pkg=common/padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*Label -pkg=common/textvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Label -pkg=common/underline
//go:generate go run ../../internal/tools/generate/main.go -recv=*Label -pkg=common/width
type Label struct {
	widget.Widget

	textVar string
}

// New creates a new entry.
func New(parent element.Element) *Label {
	label := &Label{}
	label.SetParent(parent)
	label.SetType(Type)

	label.textVar = variable.GenerateName(label.GetID())

	tk.Get().Eval("ttk::label %s -textvariable %s", label.GetID(), label.textVar)

	return label
}
