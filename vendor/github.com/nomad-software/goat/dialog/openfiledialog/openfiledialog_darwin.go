package openfiledialog

import "github.com/nomad-software/goat/internal/tk"

func init() {
	tk.Get().Eval("catch {tk_getOpenFile foo bar}")
	tk.Get().Eval("set ::tk::dialog::file::showHiddenVar false")
	tk.Get().Eval("set ::tk::dialog::file::showHiddenBtn true")
}
