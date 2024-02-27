package row

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "row"
)

// Row represents a row in the list view.
type Row struct {
	element.Ele
}

// GetValues gets the rows values.
func (el *Row) GetValues() []string {
	tk.Get().Eval("%s item %s -values", el.GetParent().GetID(), el.GetID())
	return tk.Get().GetStrSliceResult()
}

// UpdateValues updates the rows values.
func (el *Row) UpdateValues(values ...string) {
	valStr := strings.Join(values, "} {")
	tk.Get().Eval("%s item %s -values [list {%s}]", el.GetParent().GetID(), el.GetID(), valStr)
}

// SetTags sets tags for this row.
func (el *Row) SetTags(tags ...string) {
	valStr := strings.Join(tags, " ")
	tk.Get().Eval("%s item %s -tags [list %s]", el.GetParent().GetID(), el.GetID(), valStr)
}
