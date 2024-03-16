package element

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElementOverrideID(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	assert.Equal(t, "window", el.GetType())
	assert.Equal(t, ".", el.GetID())
}

func TestElementGenerateID(t *testing.T) {
	el := &Ele{}
	el.SetType("window")

	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}$`, el.GetID())
}

func TestElementMainWindowParent(t *testing.T) {
	el := &Ele{}
	el.SetID(".")
	el.SetType("window")

	child := &Ele{}
	child.SetType("window")
	child.SetParent(el)

	assert.Implements(t, (*Element)(nil), child.GetParent())
	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}$`, child.GetID())
}

func TestElementParent(t *testing.T) {
	el := &Ele{}
	el.SetType("window")

	child := &Ele{}
	child.SetType("window")
	child.SetParent(el)

	assert.Implements(t, (*Element)(nil), child.GetParent())
	assert.Regexp(t, `^\.window-[A-Z0-9]{1,8}\.window-[A-Z0-9]{1,8}$`, child.GetID())
}
