package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
}

// This is the left column that displays the instructions
// Depending on the topic selected, the instructions will change
func InstructionsTextView() *tview.TextView {
	i := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft)
	i.SetBorder(true).SetTitle("Instructions")

	i.SetText("To navigate you can either use your mouse, or SHIFT + TABn`\n\nIf you are running the commands yourself the file name will be code.EXT and the languages default extension.\n\nExample for Python:\n\ncode.py\n\nSelect an exercise to view instructions.")

	return i
}

// This is the right column that displays the resulting output
// It will either be the code or the error message
func OutputTextView() *tview.TextView {
	o := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetWrap(true).
		SetTextAlign(tview.AlignLeft)
	o.SetBorder(true).SetTitle("Output")

	return o
}

func CommandTextView(codearea *tview.TextArea, output *tview.TextView, ext string) *tview.InputField {
	var cmdField *tview.InputField
	cmdField = tview.NewInputField().
		SetLabel("Enter command: ").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetFieldTextColor(tcell.ColorDefault).
		SetLabelColor(tcell.ColorYellow).
		SetLabelWidth(15).
		SetDoneFunc(func(key tcell.Key) {
			// When the user presses enter, retrieve the input text and clear the input field
			cmd := cmdField.GetText()
			code := codearea.GetText()
			err := os.WriteFile("code/code"+ext, []byte(code), 0644)
			if err != nil {
				panic(err)
			}

			output.SetText("")

			// TODO: tab + shift from the code text area causes it to run the command
			output.SetText(RunCmd(cmd))
		})
	cmdField.SetBorder(true).SetTitle("Commands")
	return cmdField
}

func CommandButton(codearea *tview.TextArea, output *tview.TextView, ext string, cmd string) *tview.Button {
	cmdButton := tview.NewButton("Run Command").SetSelectedFunc(func() {
		// When the user presses enter, retrieve the input text and clear the input field
		code := codearea.GetText()
		err := os.WriteFile("code/code"+ext, []byte(code), 0644)
		if err != nil {
			panic(err)
		}

		output.SetText("")

		// TODO: tab + shift from the code text area causes it to run the command
		output.SetText(RunCmd(cmd))
	})
	cmdButton.SetBorder(true).SetRect(0, 0, 22, 3)

	return cmdButton
}
