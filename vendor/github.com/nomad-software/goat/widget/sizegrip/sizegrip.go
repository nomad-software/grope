package sizegrip

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "sizegrip"
)

// Used as a bottom-right corner resize widget.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_sizegrip.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*SizeGrip -pkg=common/bind
type SizeGrip struct {
	widget.Widget
}

// New creates a new sizegrip.
func New(parent element.Element) *SizeGrip {
	grip := &SizeGrip{}
	grip.SetParent(parent)
	grip.SetType(Type)

	tk.Get().Eval("ttk::sizegrip %s", grip.GetID())

	return grip
}
