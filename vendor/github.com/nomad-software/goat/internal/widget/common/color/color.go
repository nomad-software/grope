package color

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetForegroundColor sets the foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetForegroundColor(c string) {
	tk.Get().Eval("%s configure -foreground {%s}", el.GetID(), c)
}

// SetBackgroundColor sets the background color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetBackgroundColor(c string) {
	tk.Get().Eval("%s configure -background {%s}", el.GetID(), c)
}

// SetInsertColor sets the insert color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el stub) SetInsertColor(c string) {
	tk.Get().Eval("%s configure -insertbackground {%s}", el.GetID(), c)
}
