package column

import (
	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "column"
)

// Column represents a column in the list view.
type Column struct {
	element.Ele
}

// SetHeading sets the heading.
// See [option.anchor] for anchor values.
func (el *Column) SetHeading(text, anchor string) {
	tk.Get().Eval("%s heading %s -text {%s} -anchor {%s}", el.GetParent().GetID(), el.GetID(), text, anchor)
}

// SetHeadingImage sets the heading image to display at the right of the
// heading.
func (el *Column) SetHeadingImage(img *image.Image) {
	tk.Get().Eval("%s heading %s -image %s", el.GetParent().GetID(), el.GetID(), img.GetID())
}

// SetHeadingCommand sets the heading command.
func (el *Column) SetHeadingCommand(callback command.Callback) {
	name := command.GenerateName(el.GetParent().GetID(), el.GetID())
	tk.Get().CreateCommand(el, name, callback)

	tk.Get().Eval("%s heading %s -command %s", el.GetParent().GetID(), el.GetID(), name)
}

// DeleteHeadingCommand deletes the heading command.
func (el *Column) DeleteHeadingCommand() {
	tk.Get().Eval("%s heading %s -command {}", el.GetParent().GetID(), el.GetID())

	name := command.GenerateName(el.GetParent().GetID(), el.GetID())
	tk.Get().DestroyCommand(name)
}

// SetMinWidth sets the width of the column.
func (el *Column) SetWidth(width int) {
	tk.Get().Eval("%s column %s -width %d", el.GetParent().GetID(), el.GetID(), width)
}

// SetMinWidth sets the minimum width of the column.
func (el *Column) SetMinWidth(width int) {
	tk.Get().Eval("%s column %s -minwidth %d", el.GetParent().GetID(), el.GetID(), width)
}

// SetStretch sets if the column stretches or not.
func (el *Column) SetStretch(stretch bool) {
	tk.Get().Eval("%s column %s -stretch %v", el.GetParent().GetID(), el.GetID(), stretch)
}
