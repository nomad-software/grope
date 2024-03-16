// Code generated by tooling; DO NOT EDIT.
package radiobutton

import (
	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/tk"

)



// SetCommand set the command to execute on interaction with the widget.
func (el *RadioButton) SetCommand(callback command.Callback) {
	name := command.GenerateName(el.GetID())

	tk.Get().CreateCommand(el, name, callback)
	tk.Get().Eval("%s configure -command %s", el.GetID(), name)
}

// DeleteCommand deletes the command.
func (el *RadioButton) DeleteCommand() {
	tk.Get().Eval("%s configure -command {}", el.GetID())

	name := command.GenerateName(el.GetID())
	tk.Get().DestroyCommand(name)
}
