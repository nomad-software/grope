package menu

import (
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui"
	"github.com/nomad-software/goat/window"
)

// Menubar is the bar across the top of a window holding the menu items.
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/menu.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuBar -out=menubar -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuBar -out=menubar -pkg=common/borderwidth
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuBar -out=menubar -pkg=common/color -methods=SetForegroundColor,SetBackgroundColor
//go:generate go run ../../internal/tools/generate/main.go -recv=*MenuBar -out=menubar -pkg=common/relief
type MenuBar struct {
	ui.Ele
}

// NewBar creates a new menu bar to hold the menu.
func NewBar(win *window.Window) *MenuBar {
	bar := &MenuBar{}
	bar.SetParent(win)
	bar.SetType("menubar")

	tk.Get().Eval("menu %s -tearoff 0", bar.GetID())
	tk.Get().Eval("%s configure -menu %s", bar.GetParent().GetID(), bar.GetID())

	return bar
}

// DisableMenu disables the menu at the specified index.
func (m *MenuBar) DisableMenu(index int) {
	tk.Get().Eval("%s entryconfigure %d -state disable", m.GetID(), index)
}

// EnableMenu enables the menu at the specified index.
func (m *MenuBar) EnableMenu(index int) {
	tk.Get().Eval("%s entryconfigure %d -state normal", m.GetID(), index)
}
