package tag

import (
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element" // IGNORE
)

type stub struct{ element.Element } // IGNORE

// SetTags sets the tags.
func (el stub) SetTags(tags ...string) {
	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s itemconfigure %s -tags [list %s]", el.GetParent().GetID(), el.GetID(), tagStr)
}

// DeleteTags delete the tags.
func (el stub) DeleteTag(tag string) {
	tk.Get().Eval("%s dtag %s {%s}", el.GetParent().GetID(), el.GetID(), tag)
}

// DeleteAllTags delete all tags.
func (el stub) DeleteAllTags() {
	tk.Get().Eval("%s itemconfigure %s -tags {}", el.GetParent().GetID(), el.GetID())
}
