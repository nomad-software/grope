package entry

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/tk/variable"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

func init() {
	// Open up the entry a bit, the standard one seems a bit narrow.
	tk.Get().Eval("ttk::style configure TEntry -padding 3")
}

const (
	Type = "entry"
)

// An entry widget displays a one-line text string and allows that string to be
// edited by the user. Entry widgets support horizontal scrolling.
//
// Virtual events that can also be bound to.
// <<Clear>>
// <<Copy>>
// <<Cut>>
// <<Paste>>
// <<PasteSelection>>
// <<PrevWindow>>
// <<TraverseIn>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_entry.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/color -methods=SetForegroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/delete
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/font
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/justify
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/scrollbar -methods=AttachHorizontalScrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/selection
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/show
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/stringvar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Entry -pkg=common/width
type Entry struct {
	widget.Widget

	valueVar string
}

// New creates a new entry.
func New(parent element.Element) *Entry {
	entry := &Entry{}
	entry.SetParent(parent)
	entry.SetType(Type)

	entry.valueVar = variable.GenerateName(entry.GetID())

	tk.Get().Eval("ttk::entry %s -textvariable %s", entry.GetID(), entry.valueVar)

	return entry
}
