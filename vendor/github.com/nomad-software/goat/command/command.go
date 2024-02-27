package command

import (
	"fmt"

	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/internal/widget/ui/element/hash"
)

// BindCallback is the callback that is specified for a binding.
type Callback = func(*CommandData)
type BindCallback = func(*BindData)
type FontDialogCallback = func(*FontData)

// CommandData is data which is passed to a command callback when invoked.
type CommandData struct {
	CommandName string
	Element     element.Element
	Callback    Callback
}

// BindData is data which is passed to a bind callback when invoked.
type BindData struct {
	CommandName string
	Element     element.Element
	Callback    BindCallback
	Event       Event
}

// Event is the part of the callback data that contains information about any
// events that have taken place.
type Event struct {
	MouseButton int
	KeyCode     int
	ElementX    int
	ElementY    int
	Wheel       int
	Key         string
	ScreenX     int
	ScreenY     int
}

// FontData is data which is passed to a font dialog callback when invoked.
type FontData struct {
	CommandName string
	Element     element.Element
	Callback    FontDialogCallback
	Font        Font
}

// Font is the part of the callback data that contain information about
// dialog interaction.
type Font struct {
	Name      string
	Size      string
	Modifiers []string
}

// GenerateName generates a custom command name.
func GenerateName(args ...string) string {
	args = append(args, "command")
	hash := hash.Generate(args...)

	return fmt.Sprintf("cmd-%s", hash)
}
