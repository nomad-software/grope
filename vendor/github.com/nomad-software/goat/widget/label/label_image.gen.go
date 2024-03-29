// Code generated by tooling; DO NOT EDIT.
package label

import (
	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"

)



// SetImage sets an image and its position.
// See [option.compound] for compound values.
func (el *Label) SetImage(img *image.Image, compound string) {
	tk.Get().Eval("%s configure -image %s -compound {%s}", el.GetID(), img.GetID(), compound)
}
