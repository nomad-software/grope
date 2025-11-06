package text

import (
	"strconv"
	"strings"

	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/option/state"
	"github.com/nomad-software/goat/option/wrapmode"
	"github.com/nomad-software/goat/widget"
	"github.com/nomad-software/goat/widget/text/tag"
)

const (
	Type = "text"
)

// A text widget displays one or more lines of text and allows that text to be
// edited.
//
// Virtual events that can also be bound to.
// <<Modified>>
// <<Selection>>
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/color
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/font
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/relief
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/scrollbar
//go:generate go run ../../internal/tools/generate/main.go -recv=*Text -pkg=common/width
type Text struct {
	widget.Widget
}

// New creates a new text widget.
func New(parent element.Element) *Text {
	text := &Text{}
	text.SetParent(parent)
	text.SetType(Type)

	tk.Get().Eval("text %s -highlightthickness 0", text.GetID())

	text.EnableUndo(true)
	text.SetUndoLevels(100)
	text.SetWrapMode(wrapmode.Word)

	return text
}

// Enable enables the widget.
// See [option.state.Disabled]
func (el *Text) Enable() {
	tk.Get().Eval("%s configure -state {%s}", el.GetID(), state.Normal)
}

// Disable disables the widget.
// See [option.state.Disabled]
func (el *Text) Disable() {
	tk.Get().Eval("%s configure -state {%s}", el.GetID(), state.Disabled)
}

// EnableUndo enables undo functionality.
func (el *Text) EnableUndo(enable bool) {
	tk.Get().Eval("%s configure -undo %v", el.GetID(), enable)
}

// SetUndoLevels sets the maximum amount of undo levels.
func (el *Text) SetUndoLevels(levels int) {
	tk.Get().Eval("%s configure -maxundo %d", el.GetID(), levels)
}

// SetWrapMode sets the text wrap mode.
func (el *Text) SetWrapMode(mode string) {
	tk.Get().Eval("%s configure -wrap {%s}", el.GetID(), mode)
}

// AppendText appends text to the end.
func (el *Text) AppendText(text string) {
	text = strings.ReplaceAll(text, "$", `\$`)
	text = strings.ReplaceAll(text, "\\", `\\`)

	text = strings.ReplaceAll(text, "{", `\{`)
	text = strings.ReplaceAll(text, "}", `\}`)

	text = strings.ReplaceAll(text, "[", `\[`)
	text = strings.ReplaceAll(text, "]", `\]`)

	tk.Get().Eval("%s insert end [subst {%s}]", el.GetID(), text)
}

// AppendLine appends text to the end and adds a newline.
func (el *Text) AppendLine(text string) {
	text = strings.ReplaceAll(text, "$", `\$`)
	text = strings.ReplaceAll(text, "\\", `\\`)

	text = strings.ReplaceAll(text, "{", `\{`)
	text = strings.ReplaceAll(text, "}", `\}`)

	text = strings.ReplaceAll(text, "[", `\[`)
	text = strings.ReplaceAll(text, "]", `\]`)

	tk.Get().Eval("%s insert end [subst {%s\n}]", el.GetID(), text)
}

// InsertText inserts text at the specified line and character.
func (el *Text) InsertText(line, char int, text string) {
	text = strings.ReplaceAll(text, "$", `\$`)
	text = strings.ReplaceAll(text, "\\", `\\`)

	text = strings.ReplaceAll(text, "{", `\{`)
	text = strings.ReplaceAll(text, "}", `\}`)

	text = strings.ReplaceAll(text, "[", `\[`)
	text = strings.ReplaceAll(text, "]", `\]`)

	tk.Get().Eval("%s insert %d.%d [subst {%s}]", el.GetID(), line, char, text)
}

// GetText gets the current text.
func (el *Text) GetText() string {
	tk.Get().Eval("%s get 0.0 end", el.GetID())
	return tk.Get().GetStrResult()
}

// GetLineText gets the passed line's text.
// Lines start at 1.
func (el *Text) GetLineText(line int) string {
	tk.Get().Eval("%s get %d.0 %d.end", el.GetID(), line, line)
	return tk.Get().GetStrResult()
}

// Clear clears all the text.
func (el *Text) Clear() {
	tk.Get().Eval("%s delete 0.0 end", el.GetID())
}

// SetText sets the text.
func (el *Text) SetText(text string) {
	el.Clear()
	el.AppendText(text)
}

// Undo undo's the last edit.
func (el *Text) Undo() {
	tk.Get().Eval("%s edit undo", el.GetID())
}

// Redo redo's the last edit.
func (el *Text) Redo() {
	tk.Get().Eval("%s edit redo", el.GetID())
}

// ResetUndo clears all undo's.
func (el *Text) ResetUndo() {
	tk.Get().Eval("%s edit reset", el.GetID())
}

// Cut cuts text to the clipboard.
func (el *Text) Cut() {
	tk.Get().Eval("tk_textCut %s", el.GetID())
}

// Copy copies text to the clipboard.
func (el *Text) Copy() {
	tk.Get().Eval("tk_textCopy %s", el.GetID())
}

// Paste pastes text from the clipboard.
func (el *Text) Paste() {
	tk.Get().Eval("tk_textPaste %s", el.GetID())
}

// See scroll the context to show the specified line and character.
func (el *Text) See(line, char int) {
	tk.Get().Eval("%s see %d.%d", el.GetID(), line, char)
}

// SetPadding sets the padding.
func (el *Text) SetPadding(p int) {
	tk.Get().Eval("%s configure -padx %d -pady %d", el.GetID(), p, p)
}

// SetSelectForegroundColor sets the selection color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Text) SetSelectForegroundColor(c string) {
	tk.Get().Eval("%s configure -selectforeground {%s}", el.GetID(), c)
}

// SetSelectBackgroundColor sets the selection color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Text) SetSelectBackgroundColor(c string) {
	tk.Get().Eval("%s configure -selectbackground {%s}", el.GetID(), c)
}

// SetLineTag tags the passed line.
// The line numbering starts from 1.
func (el *Text) TagLine(line int, tag string) {
	tk.Get().Eval("%s tag add {%s} %d.0 %d.end", el.GetID(), tag, line, line)
}

// TagText tags a range of text in the widget.
// The line numbering starts from 1, character starts from 0.
func (el *Text) TagText(line int, char int, length int, tag string) {
	tk.Get().Eval("%s tag add {%s} %d.%d %d.%d", el.GetID(), tag, line, char, line, char+length)
}

// GetTag gets a tag from the canvas in order to modify its properties.
// Tags exist once they've been added to a canvas item.
func (el *Text) GetTag(name string) *tag.Tag {
	t := tag.New(el)
	t.SetID(name)

	return t
}

// GetInsertPos returns the insert cursor's position.
// The returned slice contains the line number and character index.
// Lines begin at 1 and character indexes begin at 0.
func (el *Text) GetInsertPos() []int {
	tk.Get().Eval("%s index insert", el.GetID())

	index := tk.Get().GetStrResult()
	result := make([]int, 0)

	for _, val := range strings.Split(index, ".") {
		i, _ := strconv.Atoi(val)
		result = append(result, i)
	}

	return result
}

// GetCurrentPos returns the current cursor's position.
// The returned slice contains the line number and character index.
// Lines begin at 1 and character indexes begin at 0.
func (el *Text) GetCurrentPos() []int {
	tk.Get().Eval("%s index current", el.GetID())

	index := tk.Get().GetStrResult()
	result := make([]int, 0)

	for _, val := range strings.Split(index, ".") {
		i, _ := strconv.Atoi(val)
		result = append(result, i)
	}

	return result
}
