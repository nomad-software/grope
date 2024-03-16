package fill

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetFillColor sets the fill color.
// See [option.color] for color names.
func (el stub) SetFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -fill {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetHoverFillColor sets the hover fill color.
// See [option.color] for color names.
func (el stub) SetHoverFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -activefill {%s}", el.GetParent().GetID(), el.GetID(), color)
}

// SetDisabledFillColor sets the active fill color.
// See [option.color] for color names.
func (el stub) SetDisabledFillColor(color string) {
	tk.Get().Eval("%s itemconfigure %s -disabledfill {%s}", el.GetParent().GetID(), el.GetID(), color)
}
