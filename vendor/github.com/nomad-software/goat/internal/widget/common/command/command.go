package command

import (
	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" //IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetCommand set the command to execute on interaction with the widget.
func (el stub) SetCommand(callback command.Callback) {
	name := command.GenerateName(el.GetID())

	tk.Get().CreateCommand(el, name, callback)
	tk.Get().Eval("%s configure -command %s", el.GetID(), name)
}

// DeleteCommand deletes the command.
func (el stub) DeleteCommand() {
	tk.Get().Eval("%s configure -command {}", el.GetID())

	name := command.GenerateName(el.GetID())
	tk.Get().DestroyCommand(name)
}
