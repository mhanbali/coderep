package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Practice struct {
	Ext         string `json:"ext"`
	Command     string `json:"command"`
	AutoCommand bool   `json:"auto_command"`
	Exercises   []struct {
		Title        string `json:"title"`
		Instructions string `json:"instructions"`
	} `json:"exercises"`
}

func main() {
	// TODO: I'm sure there's a better way to handle the arguments
	var languageFile string
	a := os.Args[1:]
	switch {
	case len(a) > 1:
		fmt.Println("Too many arguments. Use -h for help.")
		return
	case len(a) == 0:
		fmt.Println("No arguments. Use -h for help.")
		return
	case a[0] == "-h" || a[0] == "--help" || a[0] == "help":
		fmt.Println("Usage: go run . <path to language file>")
		return
	default:
		languageFile = a[0]
	}

	f, err := os.Open(languageFile)
	if err != nil {
		fmt.Println("Invalid file. Use -h for help.")
		return
	}
	defer f.Close()

	app := tview.NewApplication()

	var practice Practice

	byteValue, _ := io.ReadAll(f)
	json.Unmarshal(byteValue, &practice)
	ext := (practice.Ext)     // get the file extension from the language file
	cmd := (practice.Command) // get the command to run from the language file
	instructionsView := InstructionsTextView()
	outputView := OutputTextView()
	exerciseSelect := tview.NewDropDown().
		SetLabel("Select a topic (click or hit Enter): ").
		SetLabelColor(tview.Styles.PrimaryTextColor)
	exerciseSelect.SetBorder(true).SetTitle("Exercises")
	for _, v := range practice.Exercises {
		exerciseSelect.AddOption(v.Title, func() {
			instructionsView.SetText(v.Instructions)
		})
	}
	codeArea := tview.NewTextArea()
	codeArea.SetBorder(true).SetTitle("Code")

	var cmdField tview.Primitive
	if practice.AutoCommand {
		cmdField = CommandButton(codeArea, outputView, ext, cmd)
	} else {
		cmdField = CommandTextView(codeArea, outputView, ext)
	}

	// These are the components that can be cycled with Shift + Tab
	components := []tview.Primitive{
		codeArea,
		cmdField,
		exerciseSelect,
	}

	// TODO: test if KeyBacktab (Shift+Tab) works on platforms other than MacOS
	app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		if e.Key() == tcell.KeyBacktab {
			cycleFocus(app, components, false)
		}
		return e
	})

	// Take all the components and wrap them in a flex to be rendered out
	flex := tview.NewFlex().
		AddItem(instructionsView, 30, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(exerciseSelect, 3, 1, true).
			AddItem(codeArea, 0, 1, false).
			AddItem(cmdField, 3, 1, false), 0, 2, false).
		AddItem(outputView, 30, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).SetFocus(exerciseSelect).Run(); err != nil {
		panic(err)
	}
}
