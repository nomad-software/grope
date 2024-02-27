package bind

import (
	"testing"

	"github.com/nomad-software/goat/command"
	"github.com/stretchr/testify/assert"
)

func TestUIElementBind(t *testing.T) {
	el := stub{}

	el.Bind("<<Modified>>", func(data *command.BindData) {
		assert.Equal(t, ".", data.Element.GetID())
	})

	el.GenerateEvent("<<Modified>>")
}

func TestUIElementUnBind(t *testing.T) {
	el := stub{}

	el.Bind("<<Modified>>", func(data *command.BindData) {
		assert.Fail(t, "this should have been unbound")
	})

	el.UnBind("<<Modified>>")
	el.GenerateEvent("<<Modified>>")
}
