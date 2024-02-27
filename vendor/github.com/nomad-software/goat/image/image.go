package image

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

// Image provides a base implementation of an image.
type Image struct {
	element.Ele
}

// New creates a image.
func New(data string, typ string) *Image {
	img := &Image{}
	img.SetType(typ)

	tk.Get().Eval("image create photo %s -format {%s}", img.GetID(), typ)
	tk.Get().Eval("%s configure -data {%s}", img.GetID(), data)

	return img
}

// Blank clears the image of all data.
func (i *Image) Blank() {
	tk.Get().Eval("%s blank", i.GetID())
}

// SetGamma sets the gamma.
func (i *Image) SetGamma(gamma float64) {
	tk.Get().Eval("%s configure -gamma %v", i.GetID(), gamma)
}

// GetGamma gets the gamma.
func (i *Image) GetGamma() float64 {
	tk.Get().Eval("%s cget -gamma", i.GetID())
	return tk.Get().GetFloatResult()
}

// SetWidth sets the width.
func (i *Image) SetWidth(width int) {
	tk.Get().Eval("%s configure -width %d", i.GetID(), width)
}

// GetWidth gets the width.
func (i *Image) GetWidth() int {
	tk.Get().Eval("%s cget -width", i.GetID())
	return tk.Get().GetIntResult()
}

// SetHeight sets the height.
func (i *Image) SetHeight(height int) {
	tk.Get().Eval("%s configure -height %d", i.GetID(), height)
}

// GetHeight gets the height.
func (i *Image) GetHeight() int {
	tk.Get().Eval("%s cget -height", i.GetID())
	return tk.Get().GetIntResult()
}

// Destroy deletes the image and cleans up its resources. Once destroyed you
// cannot refer to this image again or you will get a bad path name error from
// the interpreter.
func (i *Image) Destroy() {
	tk.Get().Eval("image delete %s", i.GetID())
	i.SetType("destroyed")
}
