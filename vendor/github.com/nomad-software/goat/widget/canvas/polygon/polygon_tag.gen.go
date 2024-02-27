// Code generated by tooling; DO NOT EDIT.
package polygon

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"

)



// SetTags sets the tags.
func (el *Polygon) SetTags(tags ...string) {
	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s itemconfigure %s -tags [list %s]", el.GetParent().GetID(), el.GetID(), tagStr)
}

// DeleteTags delete the tags.
func (el *Polygon) DeleteTag(tag string) {
	tk.Get().Eval("%s dtag %s {%s}", el.GetParent().GetID(), el.GetID(), tag)
}

// DeleteAllTags delete all tags.
func (el *Polygon) DeleteAllTags() {
	tk.Get().Eval("%s itemconfigure %s -tags {}", el.GetParent().GetID(), el.GetID())
}
