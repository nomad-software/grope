package ui

import (
	"fmt"
	"os"
	"os/exec"
	"unicode/utf8"

	"github.com/nomad-software/goat/app"
	"github.com/nomad-software/goat/command"
	"github.com/nomad-software/goat/dialog/directorydialog"
	"github.com/nomad-software/goat/dialog/messagedialog"
	dtype "github.com/nomad-software/goat/dialog/type"
	"github.com/nomad-software/goat/option/anchor"
	"github.com/nomad-software/goat/option/color"
	"github.com/nomad-software/goat/option/fill"
	"github.com/nomad-software/goat/option/relief"
	"github.com/nomad-software/goat/option/side"
	"github.com/nomad-software/goat/option/underline"
	"github.com/nomad-software/goat/option/wrapmode"
	"github.com/nomad-software/goat/widget/button"
	"github.com/nomad-software/goat/widget/checkbutton"
	"github.com/nomad-software/goat/widget/entry"
	"github.com/nomad-software/goat/widget/frame"
	"github.com/nomad-software/goat/widget/label"
	"github.com/nomad-software/goat/widget/labelframe"
	"github.com/nomad-software/goat/widget/menu"
	"github.com/nomad-software/goat/widget/scrollbar"
	"github.com/nomad-software/goat/widget/sizegrip"
	"github.com/nomad-software/goat/widget/text"
	"github.com/nomad-software/goat/window"
	"github.com/nomad-software/goat/window/protocol"
	"github.com/nomad-software/grope/file/walker"
	"github.com/nomad-software/grope/option"
)

var (
	dir         *entry.Entry
	ignore      *entry.Entry
	glob        *entry.Entry
	pattern     *entry.Entry
	observeCase *checkbutton.CheckButton
	results     *text.Text

	lines = make(map[int]Line)
)

type Line struct {
	File       string
	LineNumber string
}

func New(opts *option.Options) {
	createUI(opts)
}

func createUI(opts *option.Options) {
	app := app.New()

	main := app.GetMainWindow()
	main.SetTitle("Grope GUI")
	main.SetMinSize(800, 500)
	main.SetSize(1200, 800)

	sizegrip := sizegrip.New(main)
	sizegrip.Pack(0, 0, side.Bottom, fill.None, anchor.SouthEast, false)

	optFrame := frame.New(main, 2, relief.Groove)
	optFrame.Pack(10, 3, side.Top, fill.Horizontal, anchor.Center, false)
	optFrame.SetGridColumnWeight(1, 1)

	dirLabel := label.New(optFrame)
	dirLabel.Grid(0, 0, 5, 0, 1, 1, "w")
	dirLabel.SetText("Directory")
	dirLabel.SetWidth(12)

	dir = entry.New(optFrame)
	dir.Grid(1, 0, 5, 0, 4, 1, "ew")
	dir.SetValue(opts.Dir)

	dirButton := button.New(optFrame, "Choose...")
	dirButton.Grid(5, 0, 5, 0, 1, 1, "w")
	dirButton.SetCommand(handleChooseDirectory(main))

	ignoreLabel := label.New(optFrame)
	ignoreLabel.Grid(0, 1, 5, 0, 1, 1, "w")
	ignoreLabel.SetText("Ignore regex")
	ignoreLabel.SetWidth(12)

	ignore = entry.New(optFrame)
	ignore.Grid(1, 1, 5, 0, 1, 1, "ew")
	ignore.SetValue(opts.Ignore)

	globLabel := label.New(optFrame)
	globLabel.Grid(3, 1, 5, 0, 1, 1, "w")
	globLabel.SetText("Glob pattern")
	globLabel.SetWidth(12)

	glob = entry.New(optFrame)
	glob.Grid(4, 1, 5, 0, 1, 1, "ew")
	glob.SetValue(opts.Glob)

	searchLabel := label.New(optFrame)
	searchLabel.Grid(0, 2, 5, 0, 1, 1, "w")
	searchLabel.SetText("Search regex")
	searchLabel.SetWidth(12)

	pattern = entry.New(optFrame)
	pattern.Grid(1, 2, 5, 0, 4, 1, "ew")
	pattern.SetValue(opts.Pattern)

	observeCase = checkbutton.New(optFrame, "Observe case")
	observeCase.Grid(5, 2, 5, 0, 1, 1, "w")
	if opts.Case {
		observeCase.Check()
	}

	resultFrame := labelframe.New(main, "Results", underline.None)
	resultFrame.Pack(10, 0, side.Bottom, fill.Both, anchor.Center, true)
	resultFrame.SetPadding(5)
	resultFrame.SetGridColumnWeight(0, 1)
	resultFrame.SetGridRowWeight(0, 1)

	resultsHScroll := scrollbar.NewHorizontal(resultFrame)
	resultsHScroll.Grid(0, 1, 0, 0, 1, 1, "esw")

	resultsVScroll := scrollbar.NewVertical(resultFrame)
	resultsVScroll.Grid(1, 0, 0, 0, 1, 1, "nes")

	results = text.New(resultFrame)
	results.Grid(0, 0, 0, 0, 1, 1, "nesw")
	results.SetFont("Iosevka Term Curly", "14")
	results.SetWrapMode(wrapmode.None)
	results.SetWidth(0)
	results.SetHeight(0)

	results.AttachHorizontalScrollbar(resultsHScroll)
	results.AttachVerticalScrollbar(resultsVScroll)

	resultsHScroll.AttachWidget(results)
	resultsVScroll.AttachWidget(results)

	buttonFrame := frame.New(main, 0, relief.Groove)
	buttonFrame.Pack(0, 0, side.Bottom, fill.None, anchor.Center, false)

	searchButton := button.New(buttonFrame, "Search")
	searchButton.Pack(0, 0, side.Top, fill.None, anchor.Center, false)
	searchButton.SetCommand(handleSearch)

	createWindowBinds(main, opts)
	createMenu(main)
	configureResultTags(results)

	app.Start()
}

func showAboutDialog(win *window.Window) {
	dialog := messagedialog.New(win, "About")
	dialog.SetMessage("Grope GUI")
	dialog.SetDetail("A graphical user interface for the grope command line tool.")
	dialog.SetDialogType(dtype.Ok)
	dialog.Show()
}

func updateOptions(opts *option.Options) {
	opts.Dir = dir.GetValue()
	opts.Ignore = ignore.GetValue()
	opts.Glob = glob.GetValue()
	opts.Pattern = pattern.GetValue()
	opts.Case = observeCase.IsChecked()
}

func createWindowBinds(win *window.Window, opts *option.Options) {
	win.Bind("<Control-KeyPress-q>", func(*command.BindData) {
		win.Destroy()
	})
	win.Bind("<KeyPress-Escape>", func(*command.BindData) {
		win.Destroy()
	})

	win.Bind("<KeyPress-Return>", func(*command.BindData) {
		handleSearch(nil)
	})

	win.Bind("<F1>", func(*command.BindData) {
		showAboutDialog(win)
	})

	win.SetProtocolCommand(protocol.DeleteWindow, func(*command.CommandData) {
		updateOptions(opts)
		win.Destroy()
	})
}

func createMenu(win *window.Window) {
	bar := menu.NewBar(win)

	file := menu.New(bar, "File", underline.None)
	file.AddEntry("Search", "Return", func(*command.CommandData) {
		handleSearch(nil)
	})

	file.AddSeparator()

	file.AddEntry("Quit", "Ctrl-Q, Esc", func(*command.CommandData) {
		win.Destroy()
	})

	help := menu.New(bar, "Help", underline.None)
	help.AddEntry("About...", "F1", func(*command.CommandData) {
		showAboutDialog(win)
	})
}

func configureResultTags(results *text.Text) {
	results.GetTag("file").SetForegroundColor(color.Blue)
	results.GetTag("line").SetForegroundColor(color.DimGray)
	results.GetTag("line").SetLeftMargin(20)
	results.GetTag("line").SetLeftWrapMargin(20)
	results.GetTag("number").SetForegroundColor(color.ForestGreen)
	results.GetTag("match").SetFont("Iosevka Term Curly", "14", "bold")
	results.GetTag("match").SetForegroundColor(color.Red)
	results.GetTag("line").Bind("<Double-Button-1>", handleResultsDoubleClick)
}

func handleChooseDirectory(win *window.Window) command.Callback {
	return func(data *command.CommandData) {
		dialog := directorydialog.New(win, "Choose directory")
		dialog.SetDirectoryMustExist(true)
		dialog.Show()
		if dialog.GetValue() != "" {
			dir.SetValue(dialog.GetValue())
		}
	}
}

func handleSearch(data *command.CommandData) {
	opts := &option.Options{
		Case:    observeCase.IsChecked(),
		Dir:     dir.GetValue(),
		Glob:    glob.GetValue(),
		Ignore:  ignore.GetValue(),
		Pattern: pattern.GetValue(),
	}

	results.Enable()
	results.Clear()
	clear(lines)

	walker := walker.New(opts)
	matches, errors := walker.Walk()

	ln := 1
	for matches != nil && errors != nil {
		select {
		case match, ok := <-matches:
			if !ok {
				matches = nil
				continue
			}

			results.AppendLine(match.File)
			results.TagLine(ln, "file")

			ln += 1
			for _, line := range match.Lines {
				lines[ln] = Line{
					File:       match.File,
					LineNumber: line.Number,
				}

				results.AppendLine(fmt.Sprintf("%s: %s", line.Number, line.Text))
				results.TagLine(ln, "line")
				results.TagText(ln, 0, len(line.Number), "number")

				indexes := walker.Regex.FindAllStringSubmatchIndex(line.Text, -1)
				offset := len(line.Number) + 2

				for _, index := range indexes {
					start := utf8.RuneCountInString(line.Text[:index[0]])
					length := utf8.RuneCountInString(line.Text[index[0]:index[1]])
					results.TagText(ln, offset+start, length, "match")
				}

				ln += 1
			}
			results.AppendLine("")
			ln += 1
			results.Update()

		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
	results.Disable()
}

func handleResultsDoubleClick(data *command.BindData) {
	el := data.Element.GetParent().(*text.Text)
	pos := el.GetInsertPos()

	if val, ok := lines[pos[0]]; ok {
		cmd := exec.Command("/opt/neovide/neovide", val.File, "+"+val.LineNumber)
		if err := cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}
