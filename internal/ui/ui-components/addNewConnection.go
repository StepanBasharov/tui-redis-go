package ui_components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// RenderAddNewConnection renders the "add new connection" form screen.
func RenderAddNewConnection(
	width,
	height int,
	inputHost textinput.Model,
	inputPort textinput.Model,
	inputPassword textinput.Model,
	inputDB textinput.Model,
) string {
	content := titleStyle.Render("\uE76D Connect new instance \uE76D") + "\n\n"

	content += fmt.Sprintf("%s\n", borderStyle.Render(inputHost.View()))
	content += fmt.Sprintf("%s\n", borderStyle.Render(inputPort.View()))
	content += fmt.Sprintf("%s\n", borderStyle.Render(inputPassword.View()))
	content += fmt.Sprintf("%s\n", borderStyle.Render(inputDB.View()))

	content += "\n" + separatorStyle.Render("─────────────────────────") + "\n\n"

	content += renderHint("Tab", "next field") + "\n"
	content += renderHint("Enter", "connect") + "\n"
	content += renderHint("Esc", "back")

	box := borderStyle.Render(content)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

// renderHint formats a single key hint with aligned columns.
func renderHint(key, description string) string {
	k := hintKeyStyle.Render(fmt.Sprintf("  %-7s", key))
	d := hintTextStyle.Render(description)
	return k + " " + d
}
