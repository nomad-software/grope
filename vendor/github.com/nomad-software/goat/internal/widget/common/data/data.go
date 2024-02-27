package data

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetData sets the data of a widget.
func (el stub) SetData(data ...string) {
	values := strings.Join(data, "} {")
	tk.Get().Eval("%s configure -values [list {%s}]", el.GetID(), values)
}

// GetData gets the set data of the widget.
// If no data has been set, this will return empty.
func (el stub) GetData() []string {
	tk.Get().Eval("%s cget -values", el.GetID())
	return tk.Get().GetStrSliceResult()
}
