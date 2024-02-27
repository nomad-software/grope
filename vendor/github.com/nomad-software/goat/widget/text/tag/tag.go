package tag

import (
	"fmt"
	"strings"

	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/internal/log"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
)

const (
	Type = "texttag"
)

// Tag represents a tag in a text widget.
type Tag struct {
	element.Ele
}

// Creates a new tag.
func New(parent element.Element) *Tag {
	tag := &Tag{}
	tag.SetParent(parent)
	tag.SetType(Type)

	return tag
}

// SetBorderWidth sets border width.
func (el *Tag) SetBorderWidth(width int) {
	tk.Get().Eval("%s tag configure {%s} -borderwidth %d", el.GetParent().GetID(), el.GetID(), width)
}

// SetRelief sets the relief effect.
// See [option.relief] for relief values.
func (el *Tag) SetRelief(relief string) {
	tk.Get().Eval("%s tag configure {%s} -relief {%s}", el.GetParent().GetID(), el.GetID(), relief)
}

// SetForegroundColor sets the foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetForegroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -foreground {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetBackgroundColor sets the background color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetBackgroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -background {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetSelectForegroundColor sets the selection foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetSelectForegroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -selectforeground {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetSelectBackgroundColor sets the selection background color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetSelectBackgroundColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -selectbackground {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetFont sets the widget's font.
func (el *Tag) SetFont(font string, size string, styles ...string) {
	style := strings.Join(styles, " ")
	tk.Get().Eval("%s tag configure {%s} -font {{%s} %s %s}", el.GetParent().GetID(), el.GetID(), font, size, style)
}

// AlightText aligns the text in different ways.
// See [option.align]
func (el *Tag) AlignText(align string) {
	tk.Get().Eval("%s tag configure {%s} -justify {%s}", el.GetParent().GetID(), el.GetID(), align)
}

// SetLeftMargin sets the left margin.
func (el *Tag) SetLeftMargin(margin int) {
	tk.Get().Eval("%s tag configure {%s} -lmargin1 %d", el.GetParent().GetID(), el.GetID(), margin)
}

// SetLeftWrapMargin sets the left margin of any wrapping text.
func (el *Tag) SetLeftWrapMargin(margin int) {
	tk.Get().Eval("%s tag configure {%s} -lmargin2 %d", el.GetParent().GetID(), el.GetID(), margin)
}

// SetLeftMarginColor sets the background color of the left margin.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetLeftMarginColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -lmargincolor {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetRightMargin sets the left margin.
func (el *Tag) SetRightMargin(margin int) {
	tk.Get().Eval("%s tag configure {%s} -rmargin %d", el.GetParent().GetID(), el.GetID(), margin)
}

// SetRightMarginColor sets the background color of the left margin.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetRightMarginColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -rmargincolor {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetStrikeThrough sets whether the text has a strikethrough.
func (el *Tag) SetStrikeThrough(underline bool) {
	tk.Get().Eval("%s tag configure {%s} -overstrike %v", el.GetParent().GetID(), el.GetID(), underline)
}

// SetStrikeThroughColor sets the strikethrough foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetStrikeThroughColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -overstrikefg {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// SetUnderline sets whether the text us underlined.
func (el *Tag) SetUnderline(underline bool) {
	tk.Get().Eval("%s tag configure {%s} -underline %v", el.GetParent().GetID(), el.GetID(), underline)
}

// SetUnderlineColor sets the underline foreground color.
// See [option.color] for color names.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *Tag) SetUnderlineColor(c string) {
	tk.Get().Eval("%s tag configure {%s} -underlinefg {%s}", el.GetParent().GetID(), el.GetID(), c)
}

// Bind binds a callback to a specific binding.
// Once the callback is called, the argument contains information about the
// event and data from the ui element.
//
// # Bindings
//
// The binding argument specifies a sequence of one or more event patterns,
// with optional white space between the patterns. Each event pattern may take
// one of three forms. In the simplest case it is a single printing ASCII
// character, such as 'a' or '['. The character may not be a space character or
// the character '<'. This form of pattern matches a KeyPress event for the
// particular character. The second form of pattern is longer but more general.
// It has the following syntax.
//
//	<modifier-modifier-type-detail>
//
// The entire event pattern is surrounded by angle brackets. Inside the angle
// brackets are zero or more modifiers, an event type, and an extra piece of
// information (detail) identifying a particular button or keysym. Any of the
// fields may be omitted, as long as at least one of type and detail is
// present. The fields must be separated by white space or dashes (dashes are
// prefered). The third form of pattern is used to specify a user-defined,
// named virtual event. It has the following syntax.
//
//	<<name>>
//
// The entire virtual event pattern is surrounded by double angle brackets.
// Inside the angle brackets is the user-defined name of the virtual event.
// Modifiers, such as Shift or Control, may not be combined with a virtual
// event to modify it. Bindings on a virtual event may be created before the
// virtual event is defined, and if the definition of a virtual event changes
// dynamically, all windows bound to that virtual event will respond
// immediately to the new definition. Some widgets (e.g. menu and text) issue
// virtual events when their internal state is updated in some ways. Please see
// the documentation for each widget for details.
//
// # Modifiers
//
//	Control       Button1, B1      Mod1, M1, Command      Meta, M
//	Alt           Button2, B2      Mod2, M2, Option       Double
//	Shift         Button3, B3      Mod3, M3               Triple
//	Lock          Button4, B4      Mod4, M4               Quadruple
//	Extended      Button5, B5      Mod5, M5
//
// Where more than one value is listed, separated by commas, the values are
// equivalent. Most of the modifiers have the obvious X meanings. For example,
// Button1 requires that button 1 be depressed when the event occurs. For a
// binding to match a given event, the modifiers in the event must include all
// of those specified in the event pattern. An event may also contain
// additional modifiers not specified in the binding. For example, if button 1
// is pressed while the shift and control keys are down, the pattern
// <Control-Button-1> will match the event, but <Mod1-Button-1> will not. If no
// modifiers are specified, then any combination of modifiers may be present in
// the event.
//
// Meta and M refer to whichever of the M1 through M5 modifiers is associated
// with the Meta key(s) on the keyboard (keysyms Meta_R and Meta_L). If there
// are no Meta keys, or if they are not associated with any modifiers, then
// Meta and M will not match any events. Similarly, the Alt modifier refers to
// whichever modifier is associated with the alt key(s) on the keyboard
// (keysyms Alt_L and Alt_R).
//
// The Double, Triple and Quadruple modifiers are a convenience for specifying
// double mouse clicks and other repeated events. They cause a particular event
// pattern to be repeated 2, 3 or 4 times, and also place a time and space
// requirement on the sequence: for a sequence of events to match a Double,
// Triple or Quadruple pattern, all of the events must occur close together in
// time and without substantial mouse motion in between. For example,
// <Double-Button-1> is equivalent to <Button-1><Button-1> with the extra time
// and space requirement.
//
// The Command and Option modifiers are equivalents of Mod1 resp. Mod2, they
// correspond to Macintosh-specific modifier keys.
//
// The Extended modifier is, at present, specific to Windows. It appears on
// events that are associated with the keys on the “extended keyboard”. On a US
// keyboard, the extended keys include the Alt and Control keys at the right of
// the keyboard, the cursor keys in the cluster to the left of the numeric pad,
// the NumLock key, the Break key, the PrintScreen key, and the / and Enter
// keys in the numeric keypad.
//
// # Types
//
// The type field may be any of the standard X event types, with a few extra
// abbreviations. The type field will also accept a couple non-standard X event
// types that were added to better support the Macintosh and Windows platforms.
// Below is a list of all the valid types; where two names appear together,
// they are synonyms.
//
//	Activate                 Destroy            Map
//	ButtonPress, Button      Enter              MapRequest
//	ButtonRelease            Expose             Motion
//	Circulate                FocusIn            MouseWheel
//	CirculateRequest         FocusOut           Property
//	Colormap                 Gravity            Reparent
//	Configure                KeyPress, Key      ResizeRequest
//	ConfigureRequest         KeyRelease         Unmap
//	Create                   Leave              Visibility
//	Deactivate
//
// Most of the above events have the same fields and behaviors as events in the
// X Windowing system. You can find more detailed descriptions of these events
// in any X window programming book. A couple of the events are extensions to
// the X event system to support features unique to the Macintosh and Windows
// platforms.
//
// # Details
//
// The last part of a long event specification is detail. In the case of a
// ButtonPress or ButtonRelease event, it is the number of a button (1-5). If a
// button number is given, then only an event on that particular button will
// match; if no button number is given, then an event on any button will match.
// Note: giving a specific button number is different than specifying a button
// modifier; in the first case, it refers to a button being pressed or
// released, while in the second it refers to some other button that is already
// depressed when the matching event occurs. If a button number is given then
// type may be omitted: if will default to ButtonPress. For example, the
// specifier <1> is equivalent to <ButtonPress-1>.
//
// If the event type is KeyPress or KeyRelease, then detail may be specified in
// the form of an X keysym. Keysyms are textual specifications for particular
// keys on the keyboard; they include all the alphanumeric ASCII characters
// (e.g. “a” is the keysym for the ASCII character “a”), plus descriptions for
// non-alphanumeric characters (“comma”is the keysym for the comma character),
// plus descriptions for all the non-ASCII keys on the keyboard (e.g. “Shift_L”
// is the keysym for the left shift key, and “F1” is the keysym for the F1
// function key, if it exists). The complete list of keysyms is not presented
// here; it is available in other X documentation and may vary from system to
// system. If necessary, you can use the %K notation described below to print
// out the keysym name for a particular key. If a keysym detail is given, then
// the type field may be omitted; it will default to KeyPress. For example,
// <Control-comma> is equivalent to <Control-KeyPress-comma>.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/bind.html
func (el *Tag) Bind(binding string, callback command.BindCallback) {
	if ok := tk.Binding.MatchString(binding); !ok {
		log.Error(fmt.Errorf("invalid binding: %s", binding))
		return
	}

	name := command.GenerateName(binding, el.GetID())

	tk.Get().CreateBindCommand(el, name, callback)
	tk.Get().Eval("%s tag bind %s {%s} {%s %%b %%k %%x %%y %%D %%K %%X %%Y}", el.GetParent().GetID(), el.GetID(), binding, name)
}

// UnBind unbinds a command from the passed binding.
func (el *Tag) UnBind(binding string) {
	if ok := tk.Binding.MatchString(binding); !ok {
		log.Error(fmt.Errorf("invalid binding: %s", binding))
		return
	}

	name := command.GenerateName(binding, el.GetID())

	tk.Get().Eval("%s tag bind %s {%s} {}", el.GetParent().GetID(), el.GetID(), binding)
	tk.Get().DestroyCommand(name)
}
