package main

import (
	"os/exec"

	"github.com/rivo/tview"
)

// https://github.com/rivo/tview/issues/100#issuecomment-763131391
// cycles the focus between the elements passed in
func cycleFocus(app *tview.Application, elements []tview.Primitive, reverse bool) {
	for i, el := range elements {
		if !el.HasFocus() {
			continue
		}

		if reverse {
			i = i - 1
			if i < 0 {
				i = len(elements) - 1
			}
		} else {
			i = i + 1
			i = i % len(elements)
		}

		app.SetFocus(elements[i])
		return
	}
}

// Execute the command in the terminal
func RunCmd(cmd string) string {
	c := exec.Command("bash", "-c", cmd)
	o, err := c.CombinedOutput()
	if err != nil {
		return "[red]Error executing command: " + err.Error() + "[-]" + string(o)
	}
	return string(o)
}
