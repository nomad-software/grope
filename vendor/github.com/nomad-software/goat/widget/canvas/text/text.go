package text

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "canvastext"
)

// Text represents a text item in the canvas.
//
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/anchor
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/bind
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/delete
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/fill
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/move
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/state
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/tag
//go:generate go run ../../../internal/tools/generate/main.go -recv=*Text -pkg=canvas/zorder
type Text struct {
	element.Ele
}

// Creates a new text item.
func New(parent element.Element) *Text {
	text := &Text{}
	text.SetParent(parent)
	text.SetType(Type)

	return text
}

// SetAngle sets the angle.
func (el *Text) SetAngle(angle float64) {
	tk.Get().Eval("%s itemconfigure %s -angle %v", el.GetParent().GetID(), el.GetID(), angle)
}

// SetFont sets the font.
func (el *Text) SetFont(name, size string, modifiers ...string) {
	modStr := strings.Join(modifiers, " ")
	tk.Get().Eval("%s itemconfigure %s -font {{%s} %s %s}", el.GetParent().GetID(), el.GetID(), name, size, modStr)
}

// AlightText aligns the text in different ways when a maximum width is
// specified.
// See [option.align]
func (el *Text) AlignText(align string) {
	tk.Get().Eval("%s itemconfigure %s -justify {%s}", el.GetParent().GetID(), el.GetID(), align)
}

// SetUnderline sets the character which is underlined.
// See [option.underline] for options.
func (el *Text) SetUnderline(index int) {
	tk.Get().Eval("%s itemconfigure %s -underline %d", el.GetParent().GetID(), el.GetID(), index)
}

// SetMaxWidth sets the maximum text width.
// If this option is zero (the default) the text is broken into lines only at
// newline characters. However, if this option is non-zero then any line that
// would be longer than lineLength is broken just before a space character to
// make the line shorter than lineLength; the space character is treated as if
// it were a newline character.
func (el *Text) SetMaxWidth(width float64) {
	tk.Get().Eval("%s itemconfigure %s -width %v", el.GetParent().GetID(), el.GetID(), width)
}

// SetCoords updates the item coordinates.
func (el *Text) SetCoords(x, y float64) {
	tk.Get().Eval("%s coords %s [list %v %v]", el.GetParent().GetID(), el.GetID(), x, y)
}
