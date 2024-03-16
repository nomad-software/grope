package app

import "github.com/nomad-software/goat/internal/tk"

// init configures the environment.
func init() {
	// Fix to remove hard-coded background colors from widgets.
	tk.Get().Eval("ttk::style configure TEntry -fieldbackground {SystemWindow}")
	tk.Get().Eval("ttk::style configure TSpinbox -fieldbackground {SystemWindow}")
	tk.Get().Eval("ttk::style configure Treeview -fieldbackground {SystemWindow}")
}
